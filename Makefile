PROJECT:=chainide-admin

.PHONY: build
build:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o application .
generate:
	go generate ./...
test:
	go test -coverprofile=coverage.out ./...
deps:
	go mod download

.PHONY: check
check:
	go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-s -w" -o application .
