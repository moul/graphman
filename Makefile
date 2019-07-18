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

.PHONY: release
release:
	goreleaser --snapshot --skip-publish --rm-dist
	@echo -n "Do you want to release? [y/N] " && read ans && [ $${ans:-N} = y ]
	goreleaser --rm-dist

.PHONY: lambda-build
lambda-build:
	rm -rf lambda-build
	cd lambda && GO111MODULE=on go build -o ../lambda-build/pertify pertify.go

.PHONY: netlify-dev
netlify-dev: lambda-build
	netlify dev
