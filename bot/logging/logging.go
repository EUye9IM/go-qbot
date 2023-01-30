package logging

import (
	"go-qbot/config"
	"log"
	"os"
)

var debug bool

const logflag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix

func init() {
	debug = config.Conf().Debug
}

type Logs struct {
	prefix string
	log    *log.Logger
}

func New(prefix string) *Logs {
	return &Logs{
		prefix: prefix + ": ",
		log:    log.New(os.Stdout, prefix+": ", logflag),
	}
}

func (l Logs) Infof(format string, v ...any) {
	l.log.Printf(format, v...)
}
func (l Logs) Infoln(v ...any) {
	l.log.Println(v...)
}
func (l Logs) Debugf(format string, v ...any) {
	if debug {
		l.log.Printf(format, v...)
	}
}
func (l Logs) Debugln(v ...any) {
	if debug {
		l.log.Println(v...)
	}
}
func (l Logs) New(prefix string) *Logs {
	return &Logs{
		prefix: l.prefix + prefix + ": ",
		log:    log.New(os.Stdout, l.prefix+prefix+": ", logflag),
	}
}
