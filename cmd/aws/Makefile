.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/list list/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get get/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/create create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/update update/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/delete delete/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
