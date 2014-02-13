build:
	go build
	make -C plumbing/
	make -C gfx/
	make -C input/
install:
	go install
	make -C plumbing/ install
	make -C gfx/ install
	make -C input/ install
fmt:
	go fmt
	make -C plumbing/ fmt
	make -C gfx/ fmt
	make -C input/ fmt
