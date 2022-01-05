format:
	gofmt -w ./..

gr:
	go run github.com/99designs/gqlgen generate
	go run server.go

generate:
	go run github.com/99designs/gqlgen generate

run:
	gofmt -w ./..
	go run server.go

test:
	go test -v ./...

gqlgen_upgrade:
	go get github.com/99designs/gqlgen

tidy:
	go mod tidy

lint:
	revive ./...

check:
	gofmt -w ./..
	revive ./...
