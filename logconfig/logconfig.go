package logconfig

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Logfid *os.File

func UserHome(localpath string) error {

	// get user home folder (~ may not work on Windows)
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	Root(filepath.Clean(homedir + "/" + localpath))

	return nil
}

func Root(abspath string) error {

	// get hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Could not retrieve hostaname")
	}

	// generate name of log file
	logfilename := fmt.Sprintf("%s-%v.txt", hostname, time.Now().Format("20060102150405"))
	logfilename = abspath + "/" + logfilename
	logfilename = filepath.Clean(logfilename)

	// open the log file
	// and defer its closing in case of panic,
	Logfid, err = os.OpenFile(logfilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panicf("Could not create log file: %v\n", err)
	}

	// configure logging
	mux := io.MultiWriter(os.Stdout, Logfid)
	log.SetOutput(mux)
	log.SetFlags(log.Lmicroseconds)

	return nil

}
