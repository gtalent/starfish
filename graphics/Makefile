include $(GOROOT)/src/Make.inc

TARG=wombat/core/graphics
GOFILES=\
	*.go\

GOTESTFILES=\
	tests/*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go
