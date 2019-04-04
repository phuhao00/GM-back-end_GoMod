package ServerMgr

import (
	. "HA-back-end/Net"
	"fmt"
	"sync"
	"time"
)
//
type SManager struct {
	ClientList map[string]*Client
	MsgList    map[uint32]uint32
	UserList   map[string]map[int64]string
	msgIdNext  uint32
	isFinish   bool
	conf 		map[string]string
}

var mutex sync.Mutex
var ServerManager = &SManager{
	ClientList: make(map[string]*Client),
	MsgList:    make(map[uint32]uint32),
	UserList:   make(map[string]map[int64]string),
	msgIdNext:  uint32(time.Now().Unix()),
	isFinish:   true,
}
//
func (self *SManager)RunClient (addr string) (success bool)  {
	client, ok := self.ClientList[addr]
	if !ok {
		go self.Run(addr)
	}
	return ok && client.IsConnected()
}
//
func(self *SManager) Run(addr string)  {
	client := NewClient(fmt.Sprintf("client_http2gm%s", addr), addr,ProtocolType_TCP)
	client.Chan_Packet=make(chan *ClientPacket, 256)
	client.Chan_Connection = make(chan *Session)
	go client.Run()

	for {
		select {
		case session := <-client.Chan_Connection:
			self.OnConnected(addr, session)
		case packet := <-client.Chan_Packet:
			self.OnPacket(packet)
		}
		time.Sleep(time.Millisecond * 200)
	}
}
//
func (self *SManager) OnConnected(addr string, session *Session) {
	fmt.Printf("连接服务器%s成功", addr)
}
//
func (self *SManager) OnPacket(packet *ClientPacket) {


}
