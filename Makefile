.PHONY: test
test:
	go test -v ./...
	cd examples; go test -v ./...
