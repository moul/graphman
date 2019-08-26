GO ?= go
DOCKER_IMAGE ?= moul/graphman

.PHONY: install
install:
	cd cmd/pertify; $(GO) install .

.PHONY: test
test: unittest lint tidy

.PHONY: unittest
unittest:
	echo "" > /tmp/coverage.txt
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  $(GO) test -mod=readonly -v -cover -coverprofile=/tmp/profile.out -covermode=atomic -race ./...; \
	  if [ -f /tmp/profile.out ]; then \
	    cat /tmp/profile.out >> /tmp/coverage.txt; \
	    rm -f /tmp/profile.out; \
	  fi); done
	mv /tmp/coverage.txt .

.PHONY: lint
lint:
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  golangci-lint run --verbose ./...; \
	); done

.PHONY: tidy
tidy:
	set -e; for dir in `find . -type f -name "go.mod" | sed 's@/[^/]*$$@@' | sort | uniq`; do ( set -xe; \
	  cd $$dir; \
	  $(GO)	mod tidy; \
	); done



web/pertify/examples.js: $(wildcard ./examples/pertify/*.yml)
	@# go get moul.io/fs-bundler
	cd examples/pertify; fs-bundler --format=js --callback=examples *.yml > ../../$@

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

.PHONY: docker
docker:
	docker build \
	  --build-arg VCS_REF=`git rev-parse --short HEAD` \
	  --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	  --build-arg VERSION=`git describe --tags --always` \
	  -t $(DOCKER_IMAGE) .
