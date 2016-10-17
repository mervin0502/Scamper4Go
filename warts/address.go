package warts

import (
	"errors"
	"log"
	"os"
)

type AddressType uint8

const (
	IPv4Type     AddressType = 0x01
	IPv6Type     AddressType = 0x02
	MACType      AddressType = 0x03
	FirewireType AddressType = 0x04
)

var (
	ErrAddressType      = errors.New("unknown address type ")
	ErrReadAddress      = errors.New("read address error")
	ErrUndefinedAddress = errors.New("undefined address")
)

//DefinedAddress
type DefinedAddress struct {
	Length uint8
	Type   AddressType
	Value  []byte
}

//RefAddress
type RefAddress struct {
	Magic uint8
	ID    uint32
}

type Address struct {
	Value map[uint32]*DefinedAddress
}

//NewAddress
func NewAddress() *Address {
	return &Address{
		Value: make(map[uint32]*DefinedAddress, 0),
	}
}

//Parsing
func (a *Address) Parsing(fp *os.File) *DefinedAddress {
	v := ReadUint8(fp)
	if v != 0 {
		//defined address
		_type := ReadUint8(fp)
		addr := &DefinedAddress{
			Length: v,
			Type:   _type,
		}
		var typeLen int
		switch _type {
		case 1:
			//ipv4
			typeLen = 4
			break
		case 2:
			//ipv6
			typeLen = 16
			break
		case 3:
			//mac
			typeLen = 6
			break
		case 4:
			//firewire link address
			typeLen = 8
			break
		default:
			log.Panicln(ErrAddressType)
		}
		addr.Value = make([]byte, typeLen)
		n, err := fp.Read(addr.Value)
		if err != nil {
			log.Panicln(err)
		}
		if n != typeLen {
			log.Panicln(ErrReadAddress)
		}
		a.Value[len(a.Value)] = addr
		return addr
	} else {
		//ref address
		ref := ReadUint32(fp)
		if v, ok := a.Value[ref]; ok {
			return v
		} else {
			log.Panicln(ErrUndefinedAddress)

		}
	}
	return nil
}
