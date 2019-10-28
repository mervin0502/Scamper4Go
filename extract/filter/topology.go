package filter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"

	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	// "mervin.me/Scamper4Go/extract"
	"mervin.me/Scamper4Go/extract"
	"mervin.me/Scamper4Go/utils"
	"mervin.me/Scamper4Go/warts"
)

type path struct {
	src, dst string
}

func TopologyDump(dstFile string, filteredIPMap map[*net.IP]int, srcFiles ...string) {
	// glog.Infof("%v", srcFiles)
	if len(srcFiles) == 0 {
		glog.Errorln(extract.ErrNotFiles)
	}
	//write
	fw, err := os.Create(dstFile)
	if err != nil {
		glog.Errorln(err)
	}
	defer fw.Close()
	pathMap := make(map[path]struct{}, 0)

	for _, srcFile := range srcFiles {
		fInfo, err := os.Stat(srcFile)
		glog.Infof("%s", srcFile)
		//not exist
		if err != nil && os.IsNotExist(err) {
			glog.Errorf("%s:%s", extract.ErrNotFindFile, srcFile)
			continue
		}
		if fInfo.IsDir() {
			fs, err := ioutil.ReadDir(srcFile)
			if err != nil {
				glog.Errorln(err)
			}
			for _, f := range fs {
				extractTopoFromFile(filepath.Join(srcFile, f.Name()), filteredIPMap, pathMap)
			}
		} else {
			extractTopoFromFile(srcFile, filteredIPMap, pathMap)
		}
	}
	var buf bytes.Buffer
	for ip := range pathMap {
		buf.WriteString(fmt.Sprintf("%s\t%s\n", ip.src, ip.dst))
	}
	pathMap = nil
	buf.WriteTo(fw)
	buf.Reset()
}
func extractTopoFromFile(srcFile string, filteredIPMap map[*net.IP]int, pathMap map[path]struct{}) {
	//file type
	// glog.Infof("analysisFromFile: %s", srcFile)
	fInfo, _ := os.Stat(srcFile)
	fName := fInfo.Name()
	ext := fName[strings.LastIndex(fName, ".")+1 : len(fName)]
	// var fp io.Reader
	switch ext {
	case "creating":
		fallthrough
	case "warts":
		fp, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
		if err != nil {
			glog.Errorln(err)
		}
		analysisPath(fp, filteredIPMap, pathMap)
		fp.Close()
	case "gz":
		fr, err := os.Open(srcFile)
		if err != nil {
			glog.Errorln(err)
		}
		fp, err := gzip.NewReader(fr)
		if err != nil {
			glog.Errorln(err)
		}
		analysisPath(fp, filteredIPMap, pathMap)
		fp.Close()
		fr.Close()
	default:
		glog.Errorln(extract.ErrUnsupportFileType)
	}
}
func analysisPath(fp io.Reader, filteredIPMap map[*net.IP]int, pathMap map[path]struct{}) {
	// var buf bytes.Buffer
	warts.Init()
Loop1:
	for {

		h, err := warts.NewHeader(fp)
		if err != nil {
			if err == io.EOF {
				break Loop1
			}
			glog.Errorln(err)
			continue Loop1
		}

		data := make([]byte, h.Length)

		n, err := io.ReadFull(fp, data)
		if err != nil {
			if err == io.EOF {
				break Loop1
			}
			glog.Fatal(err)
		}
		if uint32(n) != h.Length {
			glog.Fatalf("not read full:%d/%d", n, h.Length)
		}
		glog.V(2).Infof("%s\n%#x", h, data)
		switch h.TypeValue {
		case warts.ListType:
			l := warts.NewListRecord(data)
			l.Parsing()
			// glog.Info(l)
			break
		case warts.CycleStartType:
			fallthrough
		case warts.CycleDefinitionType:
			o := warts.NewCycleDefinitionRecord(data)
			o.Parsing()
			break
		case warts.CycleStopType:
			o := warts.NewCycleStopRecord(data)
			o.Parsing()
			break
		case warts.AddrType:
			//old address
			warts.NewOldAddress(data)
			break
		case warts.TracerouteType:

			tr := warts.NewTracerouteRecord(data)
			tr.Parsing()
			var pre, cur string
			var p path
			flag1 := false
			flag2 := false
			if warts.IsOldAddress {
				pre = tr.Options.SrcIP.String()
				if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(pre)) {
					flag1 = true
				}
				for _, v := range tr.HopRcords {

					cur = v.HopRefAddr.String()
					if strings.EqualFold(pre, cur) {
						continue
					}
					if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(cur)) {
						flag2 = true
					}
					glog.V(2).Infof("flag:%v/%v->%v", flag1, flag2, v)
					if flag1 || flag2 {
						flag1 = flag2
						continue
					}
					p = path{pre, cur}
					if _, ok := pathMap[p]; !ok {
						pathMap[p] = struct{}{}
					}
					// buf.WriteString(fmt.Sprintf("%s\t%s\n", pre, cur))
					flag2 = flag1
					pre = cur
				}
				cur = tr.Options.DstIP.String()
				if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(cur)) {
					flag2 = true
				}
				p = path{pre, cur}
				if _, ok := pathMap[p]; !ok {
					pathMap[p] = struct{}{}
				}
			} else {
				pre = tr.Options.SrcAdress.String()
				if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(pre)) {
					flag1 = true
				}
				for _, v := range tr.HopRcords {
					cur = v.HopAddress.String()
					if strings.EqualFold(pre, cur) {
						continue
					}
					if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(cur)) {
						flag2 = true
					}
					if flag1 || flag2 {
						flag1 = flag2
						continue
					}
					p = path{pre, cur}
					if _, ok := pathMap[p]; !ok {
						pathMap[p] = struct{}{}
					}
					// buf.WriteString(fmt.Sprintf("%s\t%s\n", pre, cur))
					flag2 = flag1
					pre = cur
				}
				cur = tr.Options.DstAdress.String()
				if filteredIPMap != nil && filterFunc(filteredIPMap, net.ParseIP(cur)) {
					flag2 = true
				}
				if !flag1 && !flag2 {
					flag2 = flag1
					continue
				}
				p = path{pre, cur}
				if _, ok := pathMap[p]; !ok {
					pathMap[p] = struct{}{}
				}
			}
			break
		default:
		} //switch
	} //for
}

func filterFunc(m map[*net.IP]int, ip net.IP) bool {
	if m == nil {
		return false
	}
	for k, v := range m {
		if utils.CommonPrefixLen(*k, ip) >= v {
			return true
		}
	}
	return false
}
