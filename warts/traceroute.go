package warts

// "io"
// "log"
import (
	"github.com/golang/glog"
)

//TracerouteRecord Traceroute structures consist of traceroute parameters, hop records, and an optional series of additional data types for special types of traceroute invokation.
type TracerouteRecord struct {
	Options      *TracerouteRecordOptions
	HopCount     uint16
	HopRcords    []*TracerouteHopRecord
	OptionalData interface{}
	End          uint16

	O *Object
}

//TracerouteRecordOptions  The flags and data types that describe traceroute are as follows
type TracerouteRecordOptions struct {
	WListID         uint32
	WCycleID        uint32
	SrcIP           *OldAddress
	DstIP           *OldAddress
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
	IPOffset        uint16
}

//TracerouteHopRecord Hop records are written in series. Each hop record takes the following form
type TracerouteHopRecord struct {
	HopRefAddr  *OldAddress
	TTLOfProbe  uint8
	TTLOfReply  uint8
	HopFlags    uint8
	HopID       uint8
	RTT         uint32
	ICMPType    uint16
	ProbeSize   uint16
	ReplySize   uint16
	IPID        uint16
	ToS         uint8
	NextMTU     uint16
	QuotedIPLen uint16
	QuotedTTL   uint8
	TCPFlags    uint8
	QuotedTOS   uint8
	ICMPExt     *ICMPExtension
	HopAddress  *Address
	HopTX       *TimeVal
}

//TracerouteRecordLastDitch The format of the last-ditch data
//TODH=1
type TracerouteRecordLastDitch struct {
	TODH      uint16
	RspNum    uint16
	HopRcords []*TracerouteHopRecord
}

//TracerouteRecordPMTUD  The format of PMTUD data
//TODH=2
type TracerouteRecordPMTUD struct {
	TODH         uint16 //2
	Options      *TracerouteRecordPMTUDOptions
	HopRecordNum uint16
	HopRcords    []*TracerouteHopRecord
	Notes        *TracerouteRecordPMTUDNotes
}

//TracerouteRecordPMTUDOptions The format of the PMTUD flags and attributes
type TracerouteRecordPMTUDOptions struct {
	MTU           uint16
	PathMTU       uint16
	MTUToFirstHop uint16
	PMTUDVersion  uint8
	NoteCount     uint8
}

//TracerouteRecordPMTUDNotes The format of the PMTUD notes
type TracerouteRecordPMTUDNotes struct {
	Type           uint8
	NextMTU        uint16
	PMTUDHopsIndex uint16
}

//DoubleTree The format of doubletree data
//TODH=3
type DoubleTree struct {
	TODH    uint16 //3
	Options *DoubleTreeOptions
}

//DoubleTreeOptions  The format of the doubletree flags and attributes
type DoubleTreeOptions struct {
	LSSIPAddrID     uint32
	GSSIPAddrID     uint32
	FirstHop        uint8
	LSSIPAddr       *Address
	GSSIPAddr       *Address
	LSSName         string
	DoubleTreeFlags uint8
}

//NewTracerouteRecord new a traceroute structures
func NewTracerouteRecord(data []byte) *TracerouteRecord {

	return &TracerouteRecord{
		Options: &TracerouteRecordOptions{
			// Time: &TimeVal{
			// 	Seconds:     0,
			// 	MicroSconds: 0,
			// },
		},
		HopRcords: make([]*TracerouteHopRecord, 0),
		O:         NewObject(data),
	}
}

//Parsing parsing the traceroute record
func (t *TracerouteRecord) Parsing() {
	//flags
	flags := NewFlags(t.O)
	//TracerouteRecordOptions
	flags.Parsing(t.Options)
	//hop count
	t.HopCount = t.O.ReadUint16()
	glog.V(2).Infof("hop:%d", t.HopCount)
	//TracerouteHopRecord
	if t.HopCount > 0 {
		//hop records
		var i uint16
		for i = 0; i < t.HopCount; i++ {
			hr := &TracerouteHopRecord{}
			flags = NewFlags(t.O)
			flags.Parsing(hr)
			t.HopRcords = append(t.HopRcords, hr)
		}
	}
	//OptionalData
	//PMTUD data, Last-ditch probing results , and doubletree
	//OptionalData > header
	v := t.O.ReadUint16()
	todhType := v >> 12
	todhLen := v & 0x0FFF
	todBytes := t.O.ReadBytes(int(todhLen))
	todObj := NewObject(todBytes)
	switch todhType {
	case 1:
		tod := TracerouteRecordLastDitch{}
		tod.TODH = v
		todObj.ReadUint8()
		tod.RspNum = todObj.ReadUint16()
		var i uint16
		for i = 0; i < t.HopCount; i++ {
			hr := &TracerouteHopRecord{}
			flags = NewFlags(todObj)
			flags.Parsing(hr)
			tod.HopRcords = append(tod.HopRcords, hr)
		}
		t.OptionalData = tod
		break
	case 2:
		tod := TracerouteRecordPMTUD{}
		tod.TODH = v
		todOpt := &TracerouteRecordPMTUDOptions{}
		flags = NewFlags(todObj)
		flags.Parsing(todOpt)
		tod.Options = todOpt
		tod.HopRecordNum = todObj.ReadUint16()
		var i uint16
		for i = 0; i < t.HopCount; i++ {
			hr := &TracerouteHopRecord{}
			flags = NewFlags(todObj)
			flags.Parsing(hr)
			tod.HopRcords = append(tod.HopRcords, hr)
		}
		todNotes := &TracerouteRecordPMTUDNotes{}
		flags = NewFlags(todObj)
		flags.Parsing(todNotes)
		tod.Notes = todNotes
		t.OptionalData = tod
		break
	case 3:
		tod := DoubleTree{}
		tod.TODH = v
		dtOpt := &DoubleTreeOptions{}
		flags = NewFlags(todObj)
		flags.Parsing(dtOpt)
		tod.Options = dtOpt
		t.OptionalData = tod
		break
	default:
		// time.Sleep(5 * time.Second)
		// glog.Errorf("unsupport header type (%x) of optional traceroute data: %d", todhType, t.O.data)
	}

	t.End = t.O.ReadUint16()
}
