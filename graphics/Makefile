include $(GOROOT)/src/Make.inc

TARG=wombat2/core/graphics
GOFILES=\
	animation.go\
	canvas.go\
	color.go\
	display.go\
	flyweight.go\
	font.go\
	util.go\
	viewport.go\

CGOFILES=\
	image.go\

GOTESTFILES=\
	tests/*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
