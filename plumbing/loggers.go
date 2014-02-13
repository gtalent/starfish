package plumbing

import (
	"io/ioutil"
	l "log"
	"os"
)

var log = l.New(ioutil.Discard, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
var errlog = l.New(ioutil.Discard, "ERROR: starfish backend: ", l.Ldate|l.Ltime)

func LogOn(on bool) {
	if on {
		log = l.New(os.Stdout, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		log = l.New(ioutil.Discard, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	}
}

func ErrLogOn(on bool) {
	if on {
		errlog = l.New(os.Stdout, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		errlog = l.New(ioutil.Discard, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	}
}
