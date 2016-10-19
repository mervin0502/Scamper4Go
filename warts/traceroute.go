package warts

import (
	"io"
	// "log"
)

//TracerouteRecord
type TracerouteRecord struct {
	Options      *TracerouteRecordOptions
	HopCount     uint16
	HopRcords    []*TracerouteHopRecord
	OptionalData interface{}
	End          uint16
}

//
type TracerouteRecordOptions struct {
	WListId         uint32 `wlistid`
	WCycleId        uint32
	SrcIP           *Address
	DstIP           *Address
	Time            *TimeVal
	StopReason      uint8
	StopData        uint8
	TraceFlags      uint8
	Attempts        uint8
	HopLimit        uint8
	TraceType       uint8
	ProbeSize       uint16
	SrcPort         uint16
	DstPort         uint16
	TTL             uint8
	ToS             uint8
	Timeout         uint8
	AllowHop        uint8
	ProbeHop        uint16
	GapLimit        uint8
	TodoOfGl        uint8
	TodoOfLF        uint8
	ProbeNum        uint16
	MiniTimeOfWait  uint8
	ConfidenceLevel uint8
	SrcAdress       *Address
	DstAdress       *Address
	UserID          uint32
}
type TracerouteHopRecord struct {
	HopRefAddress *Address
	TTLOfProbe    uint8
	TTLOfReply    uint8
	HopFlags      uint8
	HopID         uint8
	RTT           uint32
	ICMPType      uint16
	ProbeSize     uint16
	ReplySize     uint16
	IPID          uint16
	ToS           uint8
	NextMTU       uint16
	QuotedIPLen   uint16
	QuotedTTL     uint8
	TCPFlags      uint8
	QuotedTOS     uint8
	ICMPExt       *ICMPExtension
	HopDefAddress *Address
}

//TracerouteRecordLastDitch
type TracerouteRecordLastDitch struct {
	RespNum  uint16
	HopRcord *TracerouteHopRecord
}

//TracerouteRecordPMTUD
type TracerouteRecordPMTUD struct {
	TraceOptHeader uint16 //2
	Options        *TracerouteRecordPMTUDOptions
	HopRecordNum   uint16
	HopRcord       *TracerouteHopRecord
	Notes          *TracerouteRecordPMTUDNotes
}

type TracerouteRecordPMTUDOptions struct {
	MTU           uint16
	PathMTU       uint16
	MTUToFirstHop uint16
	PMTUDVersion  uint8
	NoteCount     uint8
}

type TracerouteRecordPMTUDNotes struct {
	Type           uint8
	NextMTU        uint16
	PMTUDHopsIndex uint16
}

//NewTracerouteRecord
func NewTracerouteRecord() *TracerouteRecord {
	if !IsOldAddress {
		AddressArr = make(map[uint32]*Address, 0)
	}
	return &TracerouteRecord{
		Options: &TracerouteRecordOptions{
		// Time: &TimeVal{
		// 	Seconds:     0,
		// 	MicroSconds: 0,
		// },
		},
		HopRcords: make([]*TracerouteHopRecord, 0),
	}
}

func (t *TracerouteRecord) Parsing(fp io.Reader) {
	// log.Println(len(AddressArr))

	// log.Println(len(AddressArr))
	//flags
	flags := NewFlags(fp)
	flags.Parsing(t.Options)
	//hop count
	t.HopCount = ReadUint16(fp)
	if t.HopCount > 0 {
		// log.Println("hop record")
		//hop records
		var i uint16
		for i = 0; i < t.HopCount; i++ {
			hr := &TracerouteHopRecord{}
			flags = NewFlags(fp)
			flags.Parsing(hr)
			t.HopRcords = append(t.HopRcords, hr)
		}

	}

	//optional traceroute data
	// otd := ReadUint16(fp)
	// log.Printf("%b", otd)
	// otdType := otd >> 12
	// log.Println(otdType)
	// otdLen := otd & 0x0FFF
	// log.Println(otdLen)
	// ReadBytes(fp, int(otdLen))
	// switch otdType {
	// case 1:
	// 	//last-ditch data
	// 	t.OptionalData = &TracerouteRecordLastDitch{}
	// 	flags := NewFlags(fp)
	// 	flags.Parsing(t.OptionalData)
	// 	break
	// case 2:
	// 	//PMTUD
	// 	t.OptionalData = &TracerouteRecordPMTUD{}
	// 	flags := NewFlags(fp)
	// 	flags.Parsing(t.OptionalData)
	// 	break
	// case 3:
	// default:

	// }
	v := ReadUint16(fp)
	if v>>12 > 0 {
	}
	// log.Printf("%x", v)
}

func (t *TracerouteRecordPMTUDNotes) Parsing(fp io.Reader) {

}
