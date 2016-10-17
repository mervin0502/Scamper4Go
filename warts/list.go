package warts

import (
	"log"
	"os"
	// "reflect"
)

type ListRecord struct {
	WListId  uint32             `wListId`
	PListId  uint32             `pListId`
	ListName string             `listName`
	Options  *ListRecordOptions `options`
}
type ListRecordOptions struct {
	Description string `description`
	MonitorName string `monitorName`
}

func NewListRecord() *ListRecord {
	return &ListRecord{
		Options: &ListRecordOptions{},
	}
}

func (l *ListRecord) Parsing(fp *os.File) {
	//wlistid
	l.WListId = ReadUint32(fp)
	//plistid
	l.PListId = ReadUint32(fp)
	// log.Println(l.PListId)
	//listname
	l.ListName = ReadString(fp)
	log.Println(l.ListName)
	flags := NewFlags(fp)
	flags.Parsing(l.Options)

}
