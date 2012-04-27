build:
	make -C util/
	make -C graphics/
	make -C input/
fmt:
	make -C util/ fmt
	make -C graphics/ fmt
	make -C input/ fmt
