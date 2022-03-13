run:
	go run main.go

build:
	goreleaser build

test:
	go test

test-scaffolding:
	mkdir ./testing
	touch ./testing/testing.md
