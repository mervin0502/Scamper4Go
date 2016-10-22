package warts

import (
	"errors"
	"fmt"
	"io"
	"log"
	// "strconv"
	"bytes"
	"encoding/binary"
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

type Address struct {
	Value []byte
}
type OldAddress struct {
	IDMod uint8
	Type  uint8
	Value []byte
}

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

// type Address struct {
// 	Value map[uint32]*DefinedAddress
// }

var (
	AddressArr = make(map[uint32]*Address, 0)
)

func NewAddress(fp io.Reader) {
	// log.Println("NewAdress")
	ReadUint8(fp)
	//defined address
	_type := AddressType(ReadUint8(fp))

	var typeLen int
	// log.Printf("address type %x", _type)
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

	addr := &Address{
		Value: make([]byte, typeLen),
	}
	n, err := fp.Read(addr.Value)
	if err != nil {
		log.Panicln(err)
	}
	if n != typeLen {
		log.Panicln(ErrReadAddress)
	}
	// log.Printf("old address: %v", addr)
	AddressArr[uint32(len(AddressArr)+1)] = addr
}
func (a *Address) Parsing(fp io.Reader) {
	// log.Println("address parsing.")
	// a = new(Address)
	if IsOldAddress {
		ref := ReadUint32(fp)
		if v, ok := AddressArr[ref]; ok {
			// copy(a, v)
			for _, v1 := range v.Value {
				a.Value = append(a.Value, v1)
			}
			// log.Printf("search old address: %v", a.Value)
		} else {
			log.Panicln(ErrUndefinedAddress)
		}
	} else {
		v := ReadUint8(fp)
		// log.Printf("address %x", v)
		if v != 0 {
			//defined address
			_type := AddressType(ReadUint8(fp))

			var typeLen int
			// log.Printf("address type %x", _type)
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
			// a = &Address{
			// 	Value: make([]byte, typeLen),
			// }
			a.Value = make([]byte, typeLen)
			n, err := fp.Read(a.Value)
			if err != nil {
				log.Panicln(err)
			}
			if n != typeLen {
				log.Panicln(ErrReadAddress)
			}
			// log.Printf("defined address: %v", a.String())
			AddressArr[uint32(len(AddressArr))] = a
		} else {
			//ref address
			ref := ReadUint32(fp)
			// log.Printf("ref:  %x", ref)
			if v, ok := AddressArr[ref]; ok {
				// copy(a, v)
				// log.Printf("###############%v", v.Value)
				for _, v1 := range v.Value {
					a.Value = append(a.Value, v1)
				}
				// log.Printf("search defined address: %v", a.Value)
			} else {
				log.Panicln(ErrUndefinedAddress)
			}
		}
	}
}

//String
func (a *Address) String() string {
	// log.Println(a.Value)
	var str string
	for _, v := range a.Value {
		str += fmt.Sprintf("%v", v) + "."
	}

	str = str[0 : len(str)-1]
	// log.Println(str)
	return str
}

//Uint
func (a *Address) Uint() uint64 {
	// log.Println(a.Value)
	buf := bytes.NewBuffer(a.Value)
	var i uint32
	binary.Read(buf, binary.BigEndian, &i)
	return uint64(i)
}
