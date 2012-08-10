build:
	make -C util/
	make -C backend/
	make -C gfx/
	make -C input/
fmt:
	make -C util/ fmt
	make -C backend/ fmt
	make -C gfx/ fmt
	make -C input/ fmt
