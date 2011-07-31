install:
	gmake -C util/ install
	gmake -C graphics/ install
	gmake -C input/ install
clean:
	gmake -C util/ clean
	gmake -C graphics/ clean
	gmake -C input/ clean
