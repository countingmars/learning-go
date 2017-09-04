package log

import (
	"flag"
	"github.com/golang/glog"
)

func Path(path string) {
	flag.Set("log_dir", path)
	flag.Parse()
}

func Flush() {
	glog.Flush()
}
func Info(s string) {
	glog.Info(s)	
}
