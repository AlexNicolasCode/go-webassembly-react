build:
	tinygo build -o ./src/wasm/static/main.wasm -target wasm ./src/wasm/cmd/main.go

start:
	go run main.go