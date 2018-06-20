.PHONY: build
build:
	go build -o mirror main.go 

.PHONY: build-linux
build-linux:
	GOOS=linux go build -o mirror main.go

.PHONY: test
test:
	go test ./...

.PHONY: image
image:
	docker build -t mirror .

.PHONY: clean
clean:
	rm -rf mirror
