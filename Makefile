
build:
	mkdir -p bin
	go mod tidy
	go build -ldflags "-s -w" ./...