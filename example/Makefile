build:
	go build -o main
clean:
	go clean
run: build
	./main
fmt:
	go fmt
profile: build
	./main -cpuprofile=prof.prof
	go tool pprof main prof.prof
	rm prof.prof
