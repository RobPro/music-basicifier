.PHONY: check build vet lint test fmt

check: fmt vet lint test build

fmt:
	gofmt -l . | tee /tmp/fmt-check.txt
	@if [ -s /tmp/fmt-check.txt ]; then echo "gofmt issues found"; exit 1; fi

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test ./...

build:
	GOOS=windows GOARCH=amd64 go build -o bin/app.exe ./cmd/app