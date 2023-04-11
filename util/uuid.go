package util

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"time"
)

var rander = rand.Reader
var b58 = NewBitcoinBase58()

func NewUUID() string {
	b := int64ToBytes(time.Now().UnixMicro())
	ss, _ := b58.EncodeToString(b)
	var uuid [8]byte
	io.ReadFull(rander, uuid[:])
	ss1, _ := b58.EncodeToString(uuid[:])
	return ss + ss1
}

func reverseBytes(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}
func int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
func bytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func GetTime(str string) time.Time {
	t := str[:10]
	buf, _ := b58.DecodeString(t)
	i := bytesToInt64(buf)
	return time.UnixMicro(i)
}
