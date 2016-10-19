package warts

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrHeaderLength = errors.New("header length error.")
	ErrHeaderFormat = errors.New("header format error")
)
var (
	IsOldAddress = false
)

type ObjType uint16

const (
	ListType               ObjType = 0x0001
	CycleStartType         ObjType = 0x0002
	CycleDefinitionType    ObjType = 0x0003
	CycleStopType          ObjType = 0x0004
	AddrType               ObjType = 0x0005
	TracerouteType         ObjType = 0x0006
	PingType               ObjType = 0x0007
	MDAType                ObjType = 0x0008
	AliasType              ObjType = 0x0009
	NeighbourDiscoveryType ObjType = 0x000a
	TCPBitType             ObjType = 0x000b
	StingType              ObjType = 0x000c
	SniffType              ObjType = 0x000d
)

type Header struct {
	MagicNumber uint16
	TypeValue   ObjType
	Length      uint32
}

func NewHeader(fp io.Reader) *Header {
	buf := make([]byte, 8)
	n, err := fp.Read(buf)
	if err != nil {
		if err == io.EOF {
			os.Exit(0)
		}
		log.Panicln(err)
	}
	if n < 8 {
		log.Panicln(ErrHeaderLength)
	}

	//
	_buf := bytes.NewBuffer(buf)
	var m uint16
	var t ObjType
	var l uint32
	binary.Read(_buf, binary.BigEndian, &m)
	binary.Read(_buf, binary.BigEndian, &t)
	binary.Read(_buf, binary.BigEndian, &l)
	if m != 0x1205 {
		log.Panicln(ErrHeaderFormat)
	}
	if t == 0x05 {
		IsOldAddress = true
	}
	_h := &Header{
		MagicNumber: m,
		TypeValue:   t,
		Length:      l,
	}
	return _h
}

func (h *Header) String() string {
	return fmt.Sprintf("%d, %x, %d\n", h.MagicNumber, h.TypeValue, h.Length)
}
