build:
	go build
	make -C plumbing/
	make -C gfx/
	make -C input/
fmt:
	go fmt
	make -C plumbing/ fmt
	make -C gfx/ fmt
	make -C input/ fmt
