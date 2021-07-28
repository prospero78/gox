run:
	clear
	go fmt ./...
	go build -race -o ./bin/srv-race ./cmd/srv/main.go
	cd ./bin && \
	./srv-race

build:
	clear
	go fmt ./...
	go build -o ./bin/srv ./cmd/srv/main.go
	cd ./bin && \
	strip -s ./srv && \
	upx -f ./srv