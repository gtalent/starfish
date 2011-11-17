include $(GOROOT)/src/Make.inc

TARG=wombat2/core/graphics
GOFILES=\
	animation.go\
	display.go\
	flyweight.go\
	viewport.go\

CGOFILES=\
	canvas.go\
	color.go\
	font.go\
	image.go\
	util.go\

GOTESTFILES=\
	tests/*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
