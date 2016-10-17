package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"
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
func ReadUint8(fp *os.File) uint8 {
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
func ReadUint16(fp *os.File) uint16 {
	_byte := make([]byte, 2)
	n, err := fp.Read(_byte)
	if err != nil {
		log.Panicln(err)
	}
	if n != 2 {
		log.Panicln(ErrReadUint16)
	}
	var i uint16
	binary.Read(bytes.NewBuffer(_byte), binary.BigEndian, &i)
	return i
}

//ReadUint32
func ReadUint32(fp *os.File) uint32 {
	_bytes := make([]byte, 4)
	n, err := fp.Read(_bytes)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(_bytes)
	if n != 4 {
		log.Panicln(ErrReadUint32)
	}
	var i uint32
	binary.Read(bytes.NewBuffer(_bytes), binary.BigEndian, &i)
	return i
}

//ReadString
func ReadString(fp *os.File) string {
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

//ReadTime
func ReadTime(fp *os.File) {

}
