install:
	gmake -C util/ install
	gmake -C graphics/ install
clean:
	gmake -C util/ clean
	gmake -C graphics/ clean
