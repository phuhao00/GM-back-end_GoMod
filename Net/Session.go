package net

import (
	. "HA-back-end/public"
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"github.com/golang-collections/go-datastructures/queue"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	SessionState_Logout = 1
	SessionState_Login = 1 << 1
	SessionState_InGame = 1 << 2
)

type Session struct {
	Conn            net.Conn
	Userid          int64
	User 			interface{}
	State 			int32
	Account 		string

	packet          *Packet
	writeQueue		*queue.Queue
	HandleConnected func()
	ProcessPacket   func(session *Session, packet *Packet)
	HandleClose     func(session *Session)
	chan_packet 	*chan *ClientPacket
	IsClient 		bool
	connected bool
	closed bool
	syncRequests 		sync.Map

	cipher 			*rc4.Cipher
	reqIdNext uint32

	protocol 		ProtocolType
}

func NewSession(chanpacket *chan *ClientPacket, conn net.Conn, protocol ProtocolType) *Session {
	self := &Session{
		Userid:0,
		Conn:conn,
		chan_packet:chanpacket,
		packet:NewPacket(0, 0),
		writeQueue: queue.New(256),
		State: SessionState_Logout,
		protocol: protocol,
	}
	return self
}

func (self *Session) Send(packet *Packet) (int, error) {
	packet.Pack()
	err := self.writeQueue.Put(packet)
	return 0, err
}

func (self *Session) SendExistPacket(packet *Packet) (int, error) {
	err := self.writeQueue.Put(packet)
	return 0, err
}

func (self *Session) InitRC4(key string) {
	self.cipher, _ = rc4.NewCipher([]byte(key))
}

func (self *Session) RC4EnCode(data []byte) []byte {
	m_data := []byte{}
	m_data = data[0:]
	if self.cipher != nil {
		m_data = make([]byte, len(data))
		self.cipher.XORKeyStream(m_data, data)
	}
	return m_data
}

func (self *Session) SendPbmsg(messageid uint32, errcode uint32, pbmsg proto.Message) (int, error)  {
	pkt,err := NewPbPacket(messageid, errcode, pbmsg)
	if err != nil {
		fmt.Println("Marshal SendPbmsg err :", messageid)
		return 0, err
	}
	return self.Send(pkt)
}

func (self *Session) SendGMJsonMsg(messageid uint32, errcode uint32, jsonMsg interface{}) (int, error) {
	var mData []byte
	var err error
	if jsonMsg != nil {
		mData, err = json.Marshal(jsonMsg)
		if err != nil {
			fmt.Println("Marshal SendPbmsg err :", messageid)
			return 0, err
		}
	}

	mData = append(mData, 0)
	crcStr := GetCRC32Code(mData)
	mData = append(mData, crcStr...)

	mData = self.RC4EnCode(mData)
	pkt := NewPacket(uint32(messageid), errcode).Append(mData)
	return self.Send(pkt)
}

func (self *Session) SendPbmsgWithRequestId(messageid uint32, errcode uint32, pbmsg proto.Message, reqId uint32) (int, error)  {
	pkt,err := NewPbPacket(messageid, errcode, pbmsg)
	if err != nil {
		fmt.Println("Marshal SendPbmsg err :", messageid)
		return 0, err
	}
	if reqId > 0 {
		pkt.hasReqId = true
		pkt.ReqId = reqId
	}
	return self.Send(pkt)
}

func (self *Session) SendAsyncRequest(request *AsyncPBRequest) (int, error) {
	return self.SendPbmsgWithRequestId(request.MsgId, request.Errcode, request.Request, request.Id)
}

func (self *Session) Cor_SyncSendPbmsg(request *SyncPBRequest) {
	reqid := atomic.AddUint32(&self.reqIdNext, 1)
	request.Id = reqid
	if request.Ch_response == nil {
		request.Ch_response = make(chan *Packet)
	}
	self.syncRequests.Store(request.Id,request)
	self.SendPbmsgWithRequestId(request.MsgId, request.Errcode, request.Request, request.Id)
	packet := <- request.Ch_response
	if packet == nil {
		request.ResponseErrCode = -1
	} else {
		request.ResponseErrCode = int32(packet.ErrCode)
		if request.Response != nil {
			err := proto.Unmarshal(packet.MessageData(), request.Response)
			if err != nil {
				request.ResponseErrCode = -2
			}
		}
	}

	return
}

func (self *Session) GetUserID() int64 {
	return self.Userid
}

func (self *Session) GetPacket() (*Packet, error) {
	self.packet = NewPacket(0, 0)
	// 读header
	_, err := self.readData(self.packet.data, PACK_HEAD_SIZE)
	if err != nil {
		return nil, err
	}

	//logs.Print("Session get header")
	// 读够了数据，解析头（同时会扩大data到必须的大小）
	self.packet.parseHeader()
	//logs.Print("Session parseHeader finish")
	if self.packet.msgLength > 0 {
		// 读取包体
		_, err = self.readData(self.packet.data[PACK_HEAD_SIZE:], self.packet.msgLength)
		if err != nil {
			return nil, err
		}
	}
	//logs.Print("Session get body")
	if self.packet.hasReqId {
		self.packet.ReqId = binary.LittleEndian.Uint32(self.packet.data[len(self.packet.data)-4: ])
		self.packet.data = self.packet.data[: len(self.packet.data)-4]
	}
	return self.packet, nil
}

var clienttimeout int64 = 20
//go
func (self *Session) Run() {
	go self.ReadPacketCoroutine()
	go self.WritePacketCoroutine()
}

func (self *Session) ReadPacketCoroutine() {
	for {
		if self.closed {break }

		packet, _ := self.GetPacket()
		if packet == nil {
			break
		}
		self.connected = true
		if packet.MessageId == 0 {
			if !self.IsClient {
				self.SendPbmsg(0, 0, nil)
			}
		} else {
			if packet.ReqId > 0 {
				v,_ := self.syncRequests.Load(packet.ReqId)
				if(v == nil) {
					if self.chan_packet != nil {
						*self.chan_packet <- &ClientPacket{Session:self, Packet:packet}
						//self.ProcessPacket(self, packet)
					}
				} else {
					v.(*SyncPBRequest).Ch_response <- packet
					self.syncRequests.Delete(packet.ReqId)
				}
			} else {
				if self.chan_packet != nil {
					*self.chan_packet <- &ClientPacket{Session:self, Packet:packet}
					//self.ProcessPacket(self, packet)
				}
			}
		}
	}
	self.Close()
}

func (self *Session) WritePacketCoroutine(){
	i := 0
	for {
		//有可能read协程关闭了，但是写的协程并没有包需要发送，导致write协程不关闭？？
		if self.closed {break }

		queueLen := self.writeQueue.Len()
		if queueLen == 0 {
			queueLen = 1
		}
		packets, err := self.writeQueue.Get(queueLen)
		if err != nil {
			break
		}

		count := len(packets)
		//如果kcp没有超过winsize，则会丢入缓存中，超过了，会等待winsize腾出空间
		for i=0; i < count; i++ {
			self.Conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
			_, err = self.Conn.Write(packets[i].(*Packet).data)
			if err != nil {
				break
			}
		}
		if err != nil {
			break
		}
		time.Sleep(time.Millisecond*50)
	}
	self.Close()
}

func (self *Session) readData(to []byte, need int) (int, error) {
	if !self.IsClient {
		self.Conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(clienttimeout)))
	}
	n, err := io.ReadAtLeast(self.Conn, to, need)
	if err != nil {
		log.Printf("Session %p get data error:%s %d %s", self, err, self.GetUserID(), self.protocol)
	}
	return n, err

	//geted, n := 0, 0
	//var err error
	//for {
	//	n, err = self.conn.Read(self.packet.data)
	//
	//	logs.Printf("Session %p get data,len=%d,err=%s", self, len(self.packet.data), err)
	//	if err != nil {
	//		return n, err
	//	}
	//	geted += n
	//	if(geted >= need) {
	//		break
	//	}
	//}
	//return geted,err
}

func (self *Session) Close() {
	if self.closed {return }
	self.closed = true

	self.connected = false
	log.Printf("Session %p Close!", self)
	self.Conn.Close()
	if(self.HandleClose != nil) {
		self.HandleClose(self)
	}
	self.syncRequests.Range(func(k, v interface{}) bool {
		v.(*SyncPBRequest).Ch_response <- nil
		self.syncRequests.Delete(k)
		return true
	})
	self.writeQueue.Dispose()
}

func (self *Session) GetRemoteAddr() string {
	return self.Conn.RemoteAddr().String()
}

type ClientPacket struct {
	Session *Session
	Packet  *Packet
}

type SyncPBRequest struct {
	Id 				uint32
	MsgId 			uint32
	Errcode 		uint32
	Request 			proto.Message
	Response			proto.Message
	ResponseErrCode  int32
	Ch_response 		chan *Packet
	Time 			int64
}

type AsyncPBRequest struct {
	Id uint32
	MsgId 			uint32
	Errcode 		uint32
	Request 			proto.Message
}

type ClientPacketHandler func(clientpacket *ClientPacket)
