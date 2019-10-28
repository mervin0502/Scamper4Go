package extract

import (
	"errors"
)
import (
	"github.com/golang/glog"
)

var (
	//ErrNotFiles no files
	ErrNotFiles = errors.New("not files.\n ")
	//ErrNotFindFile not found files
	ErrNotFindFile = errors.New("not find file\n ")
	//ErrUnsupportFileType unsport file type
	ErrUnsupportFileType = errors.New("unsupport file type.\n ")
)

func init() {
	glog.Info("extract init")
}
