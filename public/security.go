package public

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
)
//
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
//
func GetCRC32Code(data []byte) []byte {
	ieee := crc32.NewIEEE()
	io.WriteString(ieee, string(data))
	s := ieee.Sum32()
	str := fmt.Sprintf("%08x", s)
	rdata := []byte(str)
	rdata = append(rdata, 0)
	return rdata
}
