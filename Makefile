default: test

test:
	@go test . $(OPTS)

sure:
	@go test -race .
	@go fmt .
	@go vet .
	@golint .
	@go build .

.PHONY: test
