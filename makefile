build:
	go build -o ./bin/

install:
	go install

dev:
	watchexec -e go -- go install

validate:
	go run ./cmd/validate/main.go

test:
	go test ./...
