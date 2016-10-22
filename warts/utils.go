package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
)

var (
	ErrReadUint8  = errors.New("read uint8 error.")
	ErrReadUint16 = errors.New("read uint16 error.")
	ErrReadUint32 = errors.New("read uint32 error.")
)

//TimeVal
type TimeVal struct {
	Seconds     uint32
	MicroSconds uint32
}

//ReadUint8
func ReadUint8(fp io.Reader) uint8 {
	_byte := make([]byte, 1)
	n, err := fp.Read(_byte)
	if err != nil {
		log.Panicln(err)
	}
	if n != 1 {
		log.Panicln(ErrReadUint8)
	}
	var i uint8
	binary.Read(bytes.NewBuffer(_byte), binary.BigEndian, &i)
	return i
}

//ReadUint16
func ReadUint16(fp io.Reader) uint16 {
	_byte := make([]byte, 2)
	n, err := fp.Read(_byte)
	if err != nil {
		log.Panicln(err)
	}

	if n != 2 {
		// log.Panicln(ErrReadUint16)
		log.Panicf("%d=%s", n, ErrReadUint16)
	}
	var i uint16
	binary.Read(bytes.NewBuffer(_byte), binary.BigEndian, &i)
	return i
}

//ReadUint32
func ReadUint32(fp io.Reader) uint32 {
	_bytes := make([]byte, 4)
	n, err := fp.Read(_bytes)
	if err != nil {
		log.Panicln(err)
	}
	// log.Println(_bytes)
	if n != 4 {
		log.Panicln(ErrReadUint32)
	}
	var i uint32
	binary.Read(bytes.NewBuffer(_bytes), binary.BigEndian, &i)
	return i
}

//ReadBytes
func ReadBytes(fp io.Reader, _len int) []byte {
	_bytes := make([]byte, _len)
	n, err := fp.Read(_bytes)
	if err != nil {
		log.Panicln(err)
	}
	if n != _len {
		log.Panicln("read bytes error.")
	}
	return _bytes
}

//ReadString
func ReadString(fp io.Reader) string {
	buf := bytes.NewBuffer(nil)
	_byte := make([]byte, 1)
	for {
		n, err := fp.Read(_byte)
		if err == io.EOF {
			log.Panicln(err)
		}
		if n != 1 {
			log.Panic("...")
		}

		if _byte[0] == 0x00 {
			// log.Println(_byte[0])
			break
		}
		buf.WriteByte(_byte[0])
	}
	// log.Println("Read string")
	return buf.String()
}

func (t *TimeVal) Parsing(fp io.Reader) {
	t.Seconds = ReadUint32(fp)
	t.MicroSconds = ReadUint32(fp)
}
func (t *TimeVal) String() string {
	return fmt.Sprintf("Seconds:%d\tMicrosconds:%d\n", t.Seconds, t.MicroSconds)
}

// //ReadTime
// func ReadTime(fp io.Reader) TimeVal {
// 	return TimeVal{
// 		Seconds:     ReadUint32(fp),
// 		MicroSconds: ReadUint32(fp),
// 	}
// }
