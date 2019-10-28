package warts

import (
	"bytes"
	"encoding/binary"

	"github.com/golang/glog"
)

type Object struct {
	data []byte
	p    int
}

func NewObject(bs []byte) *Object {
	o := &Object{
		data: bs,
		p:    0,
	}
	return o
}

//ReadUint8
func (o *Object) ReadUint8() uint8 {
	if o.p == len(o.data) {
		return 0
	}
	var i uint8
	binary.Read(bytes.NewBuffer(o.data[o.p:o.p+1]), binary.BigEndian, &i)
	o.p += 1
	return i
}

//ReadUint16
func (o *Object) ReadUint16() uint16 {
	if o.p == len(o.data) {
		return 0
	}
	var i uint16
	binary.Read(bytes.NewBuffer(o.data[o.p:o.p+2]), binary.BigEndian, &i)
	o.p += 2
	return i
}

//ReadUint32
func (o *Object) ReadUint32() uint32 {
	if o.p == len(o.data) {
		return 0
	}
	var i uint32
	binary.Read(bytes.NewBuffer(o.data[o.p:o.p+4]), binary.BigEndian, &i)
	o.p += 4
	return i
}

//ReadString
func (o *Object) ReadString() string {
	buf := bytes.NewBuffer(nil)
	for k, b := range o.data[o.p:] {
		if b == 0x00 {
			o.p += k + 1
			glog.V(2).Infof("ReadString:%v", o.data[o.p:])
			break
		}
		buf.WriteByte(b)
	}
	return buf.String()
}
func (o *Object) ReadBytes(l int) []byte {
	if o.p == len(o.data) {
		return nil
	}
	_bytes := make([]byte, l)
	copy(_bytes, o.data[o.p:o.p+l])
	o.p += l
	return _bytes
}
