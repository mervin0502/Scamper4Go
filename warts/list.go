package warts

// "log"
// "reflect"
// "os"
import (
	"fmt"
)

type ListRecord struct {
	WListID  uint32             `wListId`
	PListID  uint32             `pListId`
	ListName string             `listName`
	Options  *ListRecordOptions `options`

	O *Object
}
type ListRecordOptions struct {
	Description string `description`
	MonitorName string `monitorName`
}

func NewListRecord(data []byte) *ListRecord {
	return &ListRecord{
		Options: &ListRecordOptions{},
		O:       NewObject(data),
	}
}

func (l *ListRecord) Parsing() {
	//wlistid
	l.WListID = l.O.ReadUint32()
	//plistid
	l.PListID = l.O.ReadUint32()
	// log.Println(l.PListId)
	//listname
	l.ListName = l.O.ReadString()
	// log.Println(l.ListName)
	flags := NewFlags(l.O)
	flags.Parsing(l.Options)

}

func (l *ListRecord) String() string {
	return fmt.Sprintf("%d | %d | %s", l.WListID, l.PListID, l.ListName)
}
