package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	// "io"

	// "github.com/golang/glog"
)

var (
	ErrReadIncomplete = errors.New("read incomplete.")
)

//ICMPExtension
type ICMPExtension struct {
	Len   uint16
	Value []*ICMPExtensionRecord
	O    *Object
}

//ICMPExtensionRecord
type ICMPExtensionRecord struct {
	Len      uint16
	ClassNum uint8
	TypeNum  uint8
	Value    []byte
}

//NewICMPExtension
func NewICMPExtension(obj *Object) *ICMPExtension {
	return &ICMPExtension{
		Value: make([]*ICMPExtensionRecord, 0),
		O:obj,
	}
}

func (i *ICMPExtension) Parsing() {
	_len := i.O.ReadUint16()
	_bytes := i.O.ReadBytes(int(_len))
	
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
