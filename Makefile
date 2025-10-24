all:
	mkdir -p build/
	go build -o build/host cmd/host/main.go
	go build -o build/client cmd/client/main.go
