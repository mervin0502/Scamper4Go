package warts

import (
	"errors"
	"fmt"

	// "io"
	"log"

	"github.com/golang/glog"

	// "strconv"
	"bytes"
	"encoding/binary"
)

type AddressType uint8

const (
	//IPv4Type
	IPv4Type AddressType = 0x01
	//IPv6Type
	IPv6Type AddressType = 0x02
	//MACType
	MACType AddressType = 0x03
	//FirewireType
	FirewireType AddressType = 0x04
)

var (
	ErrAddressType      = errors.New("unknown address type ")
	ErrReadAddress      = errors.New("read address error")
	ErrUndefinedAddress = errors.New("undefined address")
)

type Address struct {
	Value []byte
	O     *Object
}

func NewOldAddress(data []byte) {
	glog.V(2).Infof("new address")
	// time.Sleep(3 * time.Second)
	O := NewObject(data)
	// log.Println("NewAdress")
	//id
	O.ReadUint8()
	//type
	_type := AddressType(O.ReadUint8())

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
		O: O,
	}
	addr.Value = O.ReadBytes(typeLen)
	addrCounter++
	addressContainer[addrCounter] = addr
}

//Parsing
func (a *Address) Parsing() {
	// glog.V(2).Infof("%x", a.O.data[a.O.p:])
	v := a.O.ReadUint8()
	if v != 0 {
		//defined address
		_type := AddressType(a.O.ReadUint8())
		glog.V(2).Infof("address %d= %v", v, _type)
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
		a.Value = a.O.ReadBytes(typeLen)
		glog.V(2).Infof("#################\n#%d###:%d:%s\n###########", len(addressContainer), addrCounter, a.String())
		addressContainer[addrCounter] = a
		addrCounter++
	} else {
		//ref address
		ref := a.O.ReadUint32()
		if v, ok := addressContainer[ref]; ok {
			a.Value = v.Value
		} else {
			glog.V(2).Infof("%d|%d: %v", len(addressContainer), ref, addressContainer)
			glog.Fatalln(ErrUndefinedAddress)
		}
	}
}

//Hex
func (a *Address) Hex() string {
	// log.Printf("%d", a.Value)
	// glog.Infof("%T", a)
	var str string
	for _, v := range a.Value {
		str += fmt.Sprintf("%02X", v)
	}
	return str
}

//String
func (a *Address) String() string {
	// log.Println(a.Value)
	var str string
	switch len(a.Value) {
	case 4:
		for _, v := range a.Value {
			str += fmt.Sprintf("%v", v) + "."
		}
		break
	case 16:
		for _, v := range a.Value {
			str += fmt.Sprintf("%X", v) + ":"
		}
		break
	default:
		glog.Fatalln("undefined address length")
	}

	str = str[0 : len(str)-1]
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

type OldAddress struct {
	Address
}

func (oa *OldAddress) Parsing() {
	// glog.V(2).Infof("%x", a.O.data[a.O.p:])
	ref := oa.O.ReadUint32()
	if v, ok := addressContainer[ref]; ok {
		// bs := make([]byte, len(v.Value))
		// copy(bs, v.Value)
		oa.Value = v.Value
		// glog.Infof("IsOldAddress：%d", oa.Value)
		// log.Printf("search old address: %v", a.Value)
	} else {
		log.Panicln(ErrUndefinedAddress)
	}

}
