package warts

import (
	"os"
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
	h := NewHeader(fp)
	t.Log(h)
	l := NewListRecord()
	l.Parsing(fp)
	t.Logf("%d\t%d\t%s\t%s\t%s\n", l.WListId, l.PListId, l.ListName, l.Options.Description, l.Options.MonitorName)
	h = NewHeader(fp)
	t.Log(h)

	o := NewCycleDefinitionRecord()

	o.Parsing(fp)
	t.Logf("%d==%d==%d==%d==%d==%s\n", o.CCycleId, o.ListId, o.HCycleId, o.StartTime, o.Options.StopTime, o.Options.Hostname)
}
