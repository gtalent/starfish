include $(GOROOT)/src/Make.inc

TARG=wombat2/core/graphics
GOFILES=\
	animation.go\
	flyweight.go\
	viewport.go\

CGOFILES=\
	canvas.go\
	color.go\
	display.go\
	font.go\
	image.go\
	util.go\

GOTESTFILES=\
	tests/*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
