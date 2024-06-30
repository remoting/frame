package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"
)

var rander = rand.Reader

func NewUUID() string {
	b := uint64(time.Now().UnixMicro())
	ss := strconv.FormatUint(b, 36)
	var uuid [8]byte
	io.ReadFull(rander, uuid[:])
	ss1 := strconv.FormatUint(binary.BigEndian.Uint64(uuid[:]), 36)
	return ss + "-" + ss1
}
func NewCode() string {
	var uuid [8]byte
	io.ReadFull(rander, uuid[:])
	return strconv.FormatUint(binary.BigEndian.Uint64(uuid[:]), 36)
}
func GetUUIDTime(str string) (time.Time, error) {
	idx := strings.Index(str, "-")
	if idx > 7 {
		t := str[:idx]
		buf, err := strconv.ParseUint(t, 36, 64)
		if err != nil {
			return time.Now(), err
		}
		return time.UnixMicro(int64(buf)), nil
	}
	return time.Now(), errors.New("error")
}
