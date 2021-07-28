run:
	clear
	go fmt ./...
	GOARCH=wasm GOOS=js go build -o ./bin/main.wasm ./cmd/gox/main.go
	GOARCH=amd64 GOOS=linux go build -race -o ./bin/srv-race ./cmd/srv/main.go
	cd ./bin && \
	./srv-race

build:
	clear
	go fmt ./...
	GOARCH=wasm GOOS=js go build -o ./bin/main.wasm ./cmd/gox/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/srv ./cmd/srv/main.go
	cd ./bin && \
	strip -s ./srv && \
	upx -f ./srv