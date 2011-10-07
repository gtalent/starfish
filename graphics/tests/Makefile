include $(GOROOT)/src/Make.inc

TARG=wombat/base/graphics
GOFILES=\
	../*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $(GOTESTFILES) $(GOFILES) 
