package warts

import (
	// "log"
	"os"
)

type CycleStartRecord struct {
	CCycleId  uint32 `cCycleId`
	ListId    uint32 `listId`
	HCycleId  uint32 `hCycleId`
	StartTime uint32 `startTime`
	Options   *CycleStartRecordOptions
}

type CycleStartRecordOptions struct {
	StopTime uint32 `stopTime`
	Hostname string `hostname`
}

//CycleDefinitionRecord
type CycleDefinitionRecord struct {
	CCycleId  uint32 `cCycleId`
	ListId    uint32 `listId`
	HCycleId  uint32 `hCycleId`
	StartTime uint32 `startTime`
	Options   *CycleDefinitionRecordOptions
}

//CycleDefinitionRecordOptions
type CycleDefinitionRecordOptions struct {
	StopTime uint32 `stopTime`
	Hostname string `hostname`
}

//CycleStopRecord
type CycleStopRecord struct {
	CCycleId uint32 `cCycleId`
	StopTime uint32 `stopTime`
}

//NewCycleStartRecord
func NewCycleStartRecord() *CycleStartRecord {
	return &CycleStartRecord{
		Options: &CycleStartRecordOptions{},
	}
}

//NewCycleStartRecord
func NewCycleDefinitionRecord() *CycleDefinitionRecord {
	return &CycleDefinitionRecord{
		Options: &CycleDefinitionRecordOptions{},
	}
}

//Parsing
func (c *CycleStartRecord) Parsing(fp *os.File) {
	//cCycleId
	c.CCycleId = ReadUint32(fp)
	//listId
	c.ListId = ReadUint32(fp)
	//hCycleId
	c.HCycleId = ReadUint32(fp)
	//startTime
	c.StartTime = ReadUint32(fp)

	//flags
	flags := NewFlags(fp)
	flags.Parsing(c.Options)
}

//Parsing
func (c *CycleDefinitionRecord) Parsing(fp *os.File) {
	//cCycleId
	c.CCycleId = ReadUint32(fp)
	//listId
	c.ListId = ReadUint32(fp)
	//hCycleId
	c.HCycleId = ReadUint32(fp)
	//startTime
	c.StartTime = ReadUint32(fp)

	//flags
	flags := NewFlags(fp)
	flags.Parsing(c.Options)
}

//Parsing
func (c *CycleStopRecord) Parsing() {

}
