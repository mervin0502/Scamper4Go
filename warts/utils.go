package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrReadUint8  = errors.New("read uint8 error.")
	ErrReadUint16 = errors.New("read uint16 error.")
	ErrReadUint32 = errors.New("read uint32 error.")
)



//ReadUint8
func ReadUint8(data []byte) uint8 {
	var i uint8
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, &i)
	return i
}

//ReadUint16
func ReadUint16(data []byte) uint16 {
	var i uint16
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, &i)
	return i
}

//ReadUint32
func ReadUint32(data []byte) uint32 {
	var i uint32
	binary.Read(bytes.NewBuffer(data), binary.BigEndian, &i)
	return i
}

//ReadString
func ReadString(data []byte) string {
	buf := bytes.NewBuffer(nil)
	for _, b := range data {
		if b == 0x00 {
			break
		}
		buf.WriteByte(b)
	}

	// log.Println("Read string")
	return buf.String()
}
//TimeVal
type TimeVal struct {
	Seconds     uint32
	MicroSconds uint32

	O *Object
}
func NewTimeVal(obj *Object) *TimeVal {
	return &TimeVal{
		O:obj,
	}
}
//Parsing
func (t *TimeVal) Parsing() {
	t.Seconds = t.O.ReadUint32()
	t.MicroSconds = t.O.ReadUint32()
}
func (t *TimeVal) String() string {
	return fmt.Sprintf("Seconds:%d\tMicrosconds:%d\n", t.Seconds, t.MicroSconds)
}
