.PHONY: generate run test all

all: run

generate:
	go generate ./db/ent

run: test 
	go run ./cmd/your_app_name

test:
	go test ./...
