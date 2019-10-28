package warts

// "log"

type CycleStartRecord struct {
	CCycleID  uint32 `cCycleId`
	ListID    uint32 `listId`
	HCycleID  uint32 `hCycleId`
	StartTime uint32 `startTime`
	Options   *CycleStartRecordOptions

	O *Object
}

type CycleStartRecordOptions struct {
	StopTime uint32 `stopTime`
	Hostname string `hostname`
}

//CycleDefinitionRecord
type CycleDefinitionRecord struct {
	CCycleID  uint32 `cCycleId`
	ListID    uint32 `listId`
	HCycleID  uint32 `hCycleId`
	StartTime uint32 `startTime`
	Options   *CycleDefinitionRecordOptions

	O *Object
}

//CycleDefinitionRecordOptions
type CycleDefinitionRecordOptions struct {
	StopTime uint32 `stopTime`
	Hostname string `hostname`
}

//CycleStopRecord
type CycleStopRecord struct {
	CCycleID uint32 `cCycleId`
	StopTime uint32 `stopTime`

	O *Object
}

//NewCycleStartRecord
func NewCycleStartRecord(data []byte) *CycleStartRecord {
	return &CycleStartRecord{
		Options: &CycleStartRecordOptions{},
		O:       NewObject(data),
	}
}

//NewCycleStartRecord
func NewCycleDefinitionRecord(data []byte) *CycleDefinitionRecord {
	return &CycleDefinitionRecord{
		Options: &CycleDefinitionRecordOptions{},
		O:       NewObject(data),
	}
}

//NewCycleStartRecord
func NewCycleStopRecord(data []byte) *CycleStopRecord {
	return &CycleStopRecord{
		O: NewObject(data),
	}
}

//Parsing
func (c *CycleStartRecord) Parsing() {
	//cCycleId
	c.CCycleID = c.O.ReadUint32()
	//listId
	c.ListID = c.O.ReadUint32()
	//hCycleId
	c.HCycleID = c.O.ReadUint32()
	//startTime
	c.StartTime = c.O.ReadUint32()

	//flags
	flags := NewFlags(c.O)
	flags.Parsing(c.Options)
}

//Parsing
func (c *CycleDefinitionRecord) Parsing() {
	//cCycleId
	c.CCycleID = c.O.ReadUint32()
	//listId
	c.ListID = c.O.ReadUint32()
	//hCycleId
	c.HCycleID = c.O.ReadUint32()
	//startTime
	c.StartTime = c.O.ReadUint32()

	//flags
	flags := NewFlags(c.O)
	flags.Parsing(c.Options)
}

//Parsing
func (c *CycleStopRecord) Parsing() {
	//hCycleId
	c.CCycleID = c.O.ReadUint32()
	//startTime
	c.StopTime = c.O.ReadUint32()

	//flags
	NewFlags(c.O)
	// flags := NewFlags(c.O)
	// flags.Parsing(c.Options)
}
