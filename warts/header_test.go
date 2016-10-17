package warts

import (
	"os"
	"testing"
)

func Test_TracerouteHeader(t *testing.T) {
	srcFile := "../data/test.warts.creating"
	_, err := os.Stat(srcFile)
	if err == nil || os.IsExist(err) {
		t.Log("exist")
	}
	f, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	buf := make([]byte, 2)
	n, err := f.Read(buf)
	if err != nil {
		t.Logf("%d=>%v", n, err)
	} else {
		t.Logf("%d", n)
	}
	t.Logf("%x", buf)
}

func Test_NewHeader(t *testing.T) {
	srcFile := "../data/test.warts.creating"
	_, err := os.Stat(srcFile)
	if err == nil || os.IsExist(err) {
		t.Log("exist")
	}
	f, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}
	h := NewHeader(f)
	t.Log(h)
}
