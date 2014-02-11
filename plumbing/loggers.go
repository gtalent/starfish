package plumbing

import (
	l "log"
	"os"
)

var log = l.New(os.Stdout, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
var errlog = l.New(os.Stderr, "ERROR: starfish backend: ", l.Ldate|l.Ltime)

func LogOn(on bool) {
	if on {
		log = l.New(os.Stdout, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		log = l.New(nil, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	}
}

func ErrLogOn(on bool) {
	if on {
		errlog = l.New(os.Stdout, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		errlog = l.New(nil, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	}
}
