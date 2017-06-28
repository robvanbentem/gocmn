package gocmn

import (
	"github.com/op/go-logging"
	"os"
)

var f *os.File
var Log = logging.MustGetLogger("log")

var cliFormatter = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var fileFormatter = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05} %{shortfunc} %{level:.4s} %{id:03x} %{message}`,
)

func InitLogger(loc string) {
	file, err := os.OpenFile(loc, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic("could not create/open logfile, check permissions")
	}

	file.Seek(0, 2)
	filebackend := logging.NewLogBackend(file, "", 0)
	clibackend := logging.NewLogBackend(os.Stderr, "", 0)

	cliFormatter := logging.NewBackendFormatter(clibackend, cliFormatter)
	fileFormatter := logging.NewBackendFormatter(filebackend, fileFormatter)

	logging.SetBackend(cliFormatter, fileFormatter)

	f = file
}

func ShutdownLogger() {
	f.Close()
}
