build:
	go build -o bin/main src/*.go

build-wasm:
	 GOOS=js GOARCH=wasm go build -o playground/static/main.wasm src/*.go

run:
	go run src/*.go