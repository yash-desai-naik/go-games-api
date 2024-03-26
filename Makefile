build:
	docker build -t go-games-api .
run:
	docker run --rm -p 8080:8080 go-games-api
test:
	docker run --rm go-games-api go test -v ./...