.PHONY: test
test:
	go test -v ./...
	cd examples; go test -v ./...

.PHONY: tidy-all
tidy-all:
	for file in `find . -name go.mod`; do (dir=`dirname $$file`; set -xe; cd $$dir && go mod tidy); done

web/pertify/examples.js: $(wildcard ./examples/pertify/*.yml)
	@# go get moul.io/fs-bundler
	cd examples/pertify; fs-bundler --format=js --callback=examples *.yml > ../../$@

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
lambda-build: web/pertify/examples.js
	rm -rf lambda-build
	cd lambda && GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ../lambda-build/pertify pertify.go

.PHONY: netlify-dev
netlify-dev: lambda-build
	netlify dev

.PHONY: sam-local
sam-local: lambda-build
	@echo ""
	@echo "Open: http://localhost:3000/pertify/index.html"
	@echo ""
	sam local start-api --static-dir=web

.PHONY: _netlify-deps
_netlify_deps:
	cd; go get moul.io/fs-bundler
