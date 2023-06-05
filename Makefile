all: mod test build run

test:
	go test ./...

mod:
	go mod tidy
	go mod vendor

build:
	go build -o bin/salva main.go

run:
	chmod +x bin/salva
	./bin/salva

install:
	go install -v ./...