build:
	go build -o bin/main *.go

build-wasm:
	 GOOS=js GOARCH=wasm go build -o playground/static/main.wasm *.go

run:
	go run *.go