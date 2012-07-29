package backend

/*
#include "capi.h"
*/
import "C"

//export draw
func draw() {

}

func OpenDisplay() {
	C.openDisplay()
}

func CloseDisplay() {
	C.closeDisplay()
}

func Main() {
	C.mainBlock()
}

func GetDisplayTitle() string {
	return C.GoString(C.getDisplayTitle())
}

func SetDisplayTitle(t string) {
	C.setDisplayTitle(C.CString(t))
}

func GetDisplayWidth() int {
	return int(C.getDisplayWidth())
}

func GetDisplayHeight() int {
	return int(C.getDisplayHeight())
}

func SetDisplayWidth(w int) {
	C.setDisplayWidth(C.int(w))
}

func SetDisplayHeight(h int) {
	C.setDisplayHeight(C.int(h))
}

func SetDisplaySize(w, h int) {
	C.setDisplaySize(C.int(w), C.int(h))
}

func SetFullscreen(full bool) {
	//yes, this is sloppy, but I have no choice here
	if full {
		C.setFullscreen(1)
	} else {
		C.setFullscreen(0)
	}
}
