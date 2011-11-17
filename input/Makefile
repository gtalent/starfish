include $(GOROOT)/src/Make.inc

TARG=wombat/core/input
GOFILES=\
	keyboard.go\
	keyevent.go\
	keylisteners.go\
	mouse.go\
	mouseevent.go\
	mouselisteners.go\
	quit.go

CGOFILES=\
	keys.go\
	input.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
link: package
	$(O)l -o main _go_.$(O)
c:
	make -C ../Clib
