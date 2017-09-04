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
func Info(args ...interface{}) {
	glog.Info(args)	
}
func Fatal(args ...interface{}) {
	glog.Fatal(args)
}
