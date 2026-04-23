build:
	go build -o ./bin/

install:
	go install

validate:
	go run ./cmd/validate/main.go

test:
	go test ./...
