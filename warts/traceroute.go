package warts

import (
	"os"
)

//TracerouteRecord
type TracerouteRecord struct {
	Options  *TracerouteRecordOptions
	HopCount uint16
	HopRcord *TracerouteHopRecord
}

//
type TracerouteRecordOptions struct {
	WListId         uint32 `wlistid`
	WCycleId        uint32
	SrcIP           uint32
	DstIP           uint32
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
	SrcAdress
	DstAdress
	UserID uint32
}
type TracerouteHopRecord struct {
}
