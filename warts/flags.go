package warts

import (
	"errors"
	"io"
	"log"
	"reflect"
	// "unsafe"
)

//
var (
	ErrUnknownFlag = errors.New("unknown flag field.")
	ErrUnknownType = errors.New("unknown  type ")
)

type Flags struct {
	Flag []uint8
	fp   io.Reader
	// addr *Address
}

//NewFlags
func NewFlags(fp io.Reader) *Flags {
	_flags := make([]uint8, 0)
	for {
		_u := ReadUint8(fp)
		// log.Printf("%x", _u)
		var i uint8
		for i = 1; i < 8; i++ {
			_bit := (_u >> (i - 1)) & 0x01
			_flags = append(_flags, _bit)
		}
		if uint(_u>>7) == 0 {
			break
		}
	}
	_obj := &Flags{
		Flag: _flags,
		fp:   fp,
	}
	return _obj
}

//NewFlags
// func NewFlagsWithAddress(fp io.Reader, addr *Address) *Flags {
// 	flags := NewFlags(fp)
// 	// flags.addr = addr
// 	return flags
// }
func (f *Flags) Parsing(opts interface{}) {
	var s uint8
	_flags := f.Flag
	for _, v := range _flags {
		s += v
	}
	if s > 0 || len(_flags) >= 8 {
		ReadUint16(f.fp)
		// log.Println(paramLen)
		vs := reflect.ValueOf(opts).Elem()
		// log.Println(_flags)
		for i := 0; i < len(_flags); i++ {
			//out of the struct field

			if i >= vs.NumField() {
				// log.Println(ErrUnknownFlag)
				break
			}

			if _flags[i] == 0 {
				continue
			}
			// log.Println(ts.Field(i).Name)
			// log.Println(vs.Field(i).Type().String())
			//
			v := vs.Field(i)
			// log.Println(v.Kind().String())
			switch v.Kind().String() {
			case "uint8":
				vs.Field(i).SetUint(uint64(ReadUint8(f.fp)))
				break
			case "uint16":
				vs.Field(i).SetUint(uint64(ReadUint16(f.fp)))
				break
			case "uint32":
				vs.Field(i).SetUint(uint64(ReadUint32(f.fp)))
				break
			case "string":
				vs.Field(i).SetString(ReadString(f.fp))
				break
			case "ptr":

				v.Set(reflect.New(v.Type().Elem()))
				in := make([]reflect.Value, 1)
				in[0] = reflect.ValueOf(f.fp)
				v.MethodByName("Parsing").Call(in)
			default:
				log.Panicln(ErrUnknownType)
			}

		} //for
	}
}
