package extract

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	. "mervin.me/Scamper4Go/warts"
	"os"
	"strings"
)

var (
	ErrNotFiles          = errors.New("not files.")
	ErrNotFindFile       = errors.New("not find file")
	ErrUnsupportFileType = errors.New("unsupport file type.")
)

func AnalysisDump(dstFile string, srcFiles ...string) {
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
			analysis(fp, fw)
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
			analysis(fp, fw)
			fp.Close()
			fr.Close()
			break
		default:
			log.Panicln(ErrUnsupportFileType)
		}

	}

}

func analysis(fp io.Reader, fw *os.File) {

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
			// log.Println(h.TypeValue)
			NewAddress(fp)
			break
		case TracerouteType:
			tr := NewTracerouteRecord()
			tr.Parsing(fp)
			// if i == 3 {
			// 	os.Exit(0)
			// }
			// i++
			log.Printf("%v==%v", tr.HopCount, tr.Options)
			if IsOldAddress {
				buf.WriteString(tr.Options.SrcIP.String())
				buf.WriteString("\t")
				buf.WriteString(tr.Options.DstIP.String())
				buf.WriteString("##")
				buf.WriteString(fmt.Sprintf("%v", tr.Options.StopReason))
				buf.WriteString("##")
				for _, v := range tr.HopRcords {
					buf.WriteString(v.HopRefAddress.String())
					buf.WriteString("\t")
				}
				buf.WriteString("\r\n")
			} else {
				buf.WriteString(tr.Options.SrcAdress.String())
				buf.WriteString("\t")
				buf.WriteString(tr.Options.DstAdress.String())
				buf.WriteString("##")
				buf.WriteString(fmt.Sprintf("%v", tr.Options.StopReason))
				buf.WriteString("##")
				for _, v := range tr.HopRcords {
					buf.WriteString(v.HopDefAddress.String())
					buf.WriteString("\t")
				}
				buf.WriteString("\r\n")
			}
			break
		default:

		} //switch
		fw.WriteString(buf.String())
	} //for
}
