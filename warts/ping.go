package warts

//PingRecord Ping structures consist of ping parameters and responses. The general form of a ping recorded in warts is as follows:
type PingRecord struct {
	Options    *PingOptions
	ReplyCount uint16
}

//PingOptions The flags and data types that describe ping are as follows:
type PingOptions struct {
	ListID       uint32
	CycleID      uint32
	SrcIP        uint32
	DstIP        uint32
	Time         *TimeVal
	StopReason   uint8
	StopData     uint8
	DataLen      uint16
	Data         []byte
	ProbeCount   uint16
	ProbeSize    uint16
	ProbeWaitS   uint8
	TTL          uint8
	RelyCount    uint16
	PingSent     uint16
	PingMethod   uint8
	SrcPort      uint16
	DstPort      uint16
	UserID       uint32
	SrcAddress   *Address
	DstAddress   *Address
	PingFlag     uint8
	TOS          uint8
	CheckSum     uint16
	PathMTU      uint16
	ProbeTimeout uint8
	ProbeWaitMS  uint32
}
