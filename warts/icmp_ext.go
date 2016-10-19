package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
)

var (
	ErrReadIncomplete = errors.New("read incomplete.")
)

//ICMPExtension
type ICMPExtension struct {
	Len   uint16
	Value []*ICMPExtensionRecord
}

//ICMPExtensionRecord
type ICMPExtensionRecord struct {
	Len      uint16
	ClassNum uint8
	TypeNum  uint8
	Value    []byte
}

//NewICMPExtension
func NewICMPExtension() *ICMPExtension {
	return &ICMPExtension{
		Value: make([]*ICMPExtensionRecord, 0),
	}
}

func (i *ICMPExtension) Parsing(fp io.Reader) {
	_len := ReadUint16(fp)
	_bytes := make([]byte, _len)
	n, err := fp.Read(_bytes)
	if err != nil {
		log.Panicln(err)
	}
	if n != len(_bytes) {
		log.Panicln(ErrReadIncomplete)
	}
	buf := bytes.NewBuffer(_bytes)
	for {
		var _len uint16
		var _classNum uint8
		var _typeNum uint8

		binary.Read(buf, binary.BigEndian, &_len)
		binary.Read(buf, binary.BigEndian, &_classNum)
		binary.Read(buf, binary.BigEndian, &_typeNum)
		_value := make([]byte, _len)
		binary.Read(buf, binary.BigEndian, &_value)
		i.Value = append(i.Value, &ICMPExtensionRecord{
			Len:      _len,
			ClassNum: _classNum,
			TypeNum:  _typeNum,
			Value:    _value,
		})
		// log.Println(_value)
		if buf.Len() == 0 {
			break
		}
	}
}
