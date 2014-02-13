build:
	make -C util/
	make -C backend/
	make -C gfx/
	make -C input/
install:
	make -C util/ install
	make -C backend/ install
	make -C gfx/ install
	make -C input/ install
fmt:
	make -C util/ fmt
	make -C backend/ fmt
	make -C gfx/ fmt
	make -C input/ fmt
