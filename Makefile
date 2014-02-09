build:
	go build
	make -C backend/
	make -C gfx/
	make -C input/
fmt:
	go fmt
	make -C backend/ fmt
	make -C gfx/ fmt
	make -C input/ fmt
