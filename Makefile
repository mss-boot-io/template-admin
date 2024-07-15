PROJECT:=chainide-admin

.PHONY: build
build:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o application .
generate:
	go generate ./...
test:
	go test ./...

.PHONY: check
check:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o application .
