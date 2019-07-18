.PHONY: test
test:
	go test -v ./...
	cd examples; go test -v ./...

.PHONY: install
install:
	cd cmd/pertify; go install

.PHONY: lint
lint:
	golangci-lint run --verbose ./...
