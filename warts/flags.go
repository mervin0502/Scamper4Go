package warts

import (
	"errors"
	"log"
	"reflect"

	"github.com/golang/glog"
	// "unsafe"
)

//
var (
	ErrUnknownFlag = errors.New("unknown flag field.")
	ErrUnknownType = errors.New("unknown  type ")
)

type Flags struct {
	Flag []uint8
	O    *Object
}

//NewFlags
func NewFlags(obj *Object) *Flags {
	_flags := make([]uint8, 0)
	for {
		_u := obj.ReadUint8()
		var i uint8
		// glog.Infof("%d::", _u)
		for i = 1; i < 8; i++ {
			_bit := (_u >> (i - 1)) & 0x01
			// glog.Infof("%d |", _bit)
			_flags = append(_flags, _bit)
		}
		// glog.Infof(" \n")
		if uint(_u>>7) == 0 {
			break
		}
	}
	// glog.Infof(" \n\n")
	_obj := &Flags{
		Flag: _flags,
		O:    obj,
	}
	// glog.Infof("%d", _obj.Flag)
	return _obj
}

func (f *Flags) Parsing(opts interface{}) {
	var s uint8
	_flags := f.Flag
	for _, v := range _flags {
		s += v
	}
	if s > 0 || len(_flags) >= 8 {
		f.O.ReadUint16()
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
			v := vs.Field(i)
			// log.Println(v.Kind().String())
			switch v.Kind() {
			case reflect.Uint8:
				_b := f.O.ReadUint8()
				vs.Field(i).SetUint(uint64(_b))
				// glog.V(2).Infof("%s:%x", vs.Type().Field(i).Name, _b)
				break
			case reflect.Uint16:
				_u16 := f.O.ReadUint16()
				vs.Field(i).SetUint(uint64(_u16))
				glog.V(2).Infof("%s:%#0x", vs.Type().Field(i).Name, _u16)
				break
			case reflect.Uint32:
				_u32 := f.O.ReadUint32()
				vs.Field(i).SetUint(uint64(_u32))
				glog.V(2).Infof("%s:%#0x", vs.Type().Field(i).Name, _u32)
				break
			case reflect.String:
				_str := f.O.ReadString()
				vs.Field(i).SetString(_str)
				glog.V(2).Infof("%s:%#0x", vs.Type().Field(i).Name, _str)
				break
			case reflect.Ptr:
				glog.V(2).Infof("###%s", vs.Type().Field(i).Name)
				v.Set(reflect.New(v.Type().Elem()))
				// glog.Info(v.Elem().FieldByName("O"))
				v.Elem().FieldByName("O").Set(reflect.ValueOf(f.O))
				v.MethodByName("Parsing").Call(nil)
			default:
				log.Panicln(ErrUnknownType)
			}

		} //for
	}
}
