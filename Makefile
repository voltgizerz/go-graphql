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
	go test -v -cover ./...

cover-out:
	go test ./...  -coverpkg=./... -coverprofile ./coverage.out

gqlgen_upgrade:
	go get github.com/99designs/gqlgen

tidy:
	go mod tidy

lint:
	revive ./...

check:
	gofmt -w ./..
	revive -config revive.toml ./...

generate-mock:
	mockgen -source=./service/type.service.go -destination=./mocks/mock_TypeService.go 
	mockgen -source=./service/pokemon.service.go -destination=./mocks/mock_PokemonService.go 