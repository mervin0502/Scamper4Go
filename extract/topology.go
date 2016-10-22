package extract

import (
	"bytes"
	"compress/gzip"
	// "errors"
	// "fmt"
	"io"
	"log"
	. "mervin.me/Scamper4Go/warts"
	"mervin.me/util/ip"
	"os"
	"strconv"
	"strings"
)

func TopologyDump(dstFile string, srcFiles ...string) {
	if len(srcFiles) == 0 {
		log.Panicln(ErrNotFiles)
	}
	//write
	fw, err := os.Create(dstFile)
	if err != nil {
		log.Panicln(err)
	}
	defer fw.Close()
	for _, srcFile := range srcFiles {
		fInfo, err := os.Stat(srcFile)
		//not exist
		if err != nil && os.IsNotExist(err) {
			log.Panicf("%s:%s", ErrNotFindFile, srcFile)
		}
		//file type
		fName := fInfo.Name()
		ext := fName[strings.LastIndex(fName, ".")+1 : len(fName)]
		// var fp io.Reader
		switch ext {
		case "creating":
			fallthrough
		case "warts":

			fp, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
			if err != nil {
				log.Panicln(err)
			}
			// defer fp.Close()
			analysisPath(fp, fw)
			fp.Close()
			break
		case "gz":
			fr, err := os.Open(srcFile)
			if err != nil {
				log.Panicln(err)
			}
			// defer fr.Close()
			fp, err := gzip.NewReader(fr)
			if err != nil {
				log.Panicln(err)
			}

			// defer fr.Close()
			analysisPath(fp, fw)
			fp.Close()
			fr.Close()

			break
		default:
			log.Panicln(ErrUnsupportFileType)
		}
	}

}

func analysisPath(fp io.Reader, fw *os.File) {

	//read
	// i := 0
	for {
		var buf bytes.Buffer
		h := NewHeader(fp)
		switch h.TypeValue {
		case ListType:
			l := NewListRecord()
			l.Parsing(fp)
			break
		case CycleStartType:
		case CycleDefinitionType:
			o := NewCycleDefinitionRecord()
			o.Parsing(fp)
			break
		case AddrType:
			NewAddress(fp)
			break
		case TracerouteType:
			tr := NewTracerouteRecord()
			tr.Parsing(fp)
			if IsOldAddress {
				var pre uint64 = 0
				var cur uint64 = 0
				for _, v := range tr.HopRcords {
					if !ip.IsPublicIp(v.HopRefAddress.String()) {
						pre = 0
						continue
					}
					if pre == 0 {
						pre = v.HopRefAddress.Uint()
						continue
					}

					cur = v.HopRefAddress.Uint()
					buf.WriteString(strconv.FormatUint(pre, 10))
					buf.WriteString("\t")
					buf.WriteString(strconv.FormatUint(cur, 10))
					buf.WriteString("\r\n")
					// log.Printf("%d==%d\n", pre, cur)
					pre = cur
				}
			} else {
				var pre uint64 = 0
				var cur uint64 = 0
				for _, v := range tr.HopRcords {
					if !ip.IsPublicIp(v.HopDefAddress.String()) {
						log.Println(v.HopDefAddress.String())
						pre = 0
						continue
					}
					if pre == 0 {
						pre = v.HopDefAddress.Uint()
						continue
					}
					cur = v.HopDefAddress.Uint()
					buf.WriteString(strconv.FormatUint(pre, 10))
					buf.WriteString("\t")
					buf.WriteString(strconv.FormatUint(cur, 10))
					buf.WriteString("\r\n")
					// log.Printf("%d==%d\n", pre, cur)
					pre = cur
				}
			}
			break
		default:

		} //switch
		fw.WriteString(buf.String())
	} //for
}
