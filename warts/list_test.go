package warts

import (
	"io"
	"testing"
)

func Test_ReadListRecord(t *testing.T) {
	srcFile := "../data/test.warts.creating"
	_, err := os.Stat(srcFile)
	if err == nil || os.IsExist(err) {
		t.Log("exist")
	}
	fp, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
	defer fp.Close()
	if err != nil {
		t.Fatal(err)
	}
	// addr := NewAddress()
	// h := NewHeader(fp)
	// t.Log(h)
	// l := NewListRecord()
	// l.Parsing(fp)
	// t.Logf("%d\t%d\t%s\t%s\t%s\n", l.WListId, l.PListId, l.ListName, l.Options.Description, l.Options.MonitorName)
	// h = NewHeader(fp)
	// t.Log(h)

	// o := NewCycleDefinitionRecord()

	// o.Parsing(fp)
	// t.Logf("%d==%d==%d==%d==%d==%s\n", o.CCycleId, o.ListId, o.HCycleId, o.StartTime, o.Options.StopTime, o.Options.Hostname)
	// h = NewHeader(fp)
	// t.Log(h)
	// tr := NewTracerouteRecord()
	// tr.Parsing(fp)
	// t.Logf("%v", tr.Options.Time)
	for {

		h := NewHeader(fp)
		switch h.TypeValue {
		case ListType:
			l := NewListRecord()
			l.Parsing(fp)
		case CycleStartType:
		case CycleDefinitionType:
			o := NewCycleDefinitionRecord()
			o.Parsing(fp)
		case TracerouteType:
			tr := NewTracerouteRecord()
			tr.Parsing(fp)
		default:

		}
	}
}
