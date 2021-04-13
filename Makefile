.PHONY:
.SILENT:


build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/battleship ./cmd/battleship/main.go

run: build
	./.bin/battleship
