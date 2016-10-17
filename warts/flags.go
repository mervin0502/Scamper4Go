package warts

import (
	"errors"
	"log"
	"os"
	"reflect"
)

//
var (
	ErrUnknownFlag = errors.New("unknown flag field.")
	ErrUnknownType = errors.New("unknown  type ")
)

type Flags struct {
	Flag []uint8
	fp   *os.File
}

//NewFlags
func NewFlags(fp *os.File) *Flags {
	_flags := make([]uint8, 0)
	for {
		_u := ReadUint8(fp)
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

func (f *Flags) Parsing(opts interface{}) {
	var s uint8
	_flags := f.Flag
	for _, v := range _flags {
		s += v
	}
	if s > 0 || len(_flags) >= 8 {
		paramLen := ReadUint16(f.fp)
		log.Println(paramLen)

		vs := reflect.ValueOf(opts).Elem()
		for i := 0; i < len(_flags); i++ {
			//out of the struct field
			log.Println(vs.NumField())
			if i > vs.NumField() {
				log.Println(ErrUnknownFlag)
			}
			if _flags[i] == 0 {
				continue
			}
			log.Println(vs.Field(i).Kind().String())
			//
			switch vs.Field(i).Kind().String() {
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
			default:
				log.Panicln(ErrUnknownType)
			}

		} //for
	}
}
