GO ?= GO111MODULE=on go
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
	cd examples/pertify; fs-bundler --format=js --callback=examples *.yml > ../../$@

.PHONY: release
release:
	goreleaser --snapshot --skip-publish --rm-dist
	@echo -n "Do you want to release? [y/N] " && read ans && [ $${ans:-N} = y ]
	goreleaser --rm-dist

.PHONY: lambda-build
lambda-build: web/pertify/examples.js
	rm -rf lambda-build
	cd lambda && GOOS=linux GOARCH=amd64 $(GO) build -o ../lambda-build/pertify pertify.go

.PHONY: netlify-dev
netlify-dev: lambda-build
	netlify dev

.PHONY: netlify
netlify: _netlify_deps lambda-build install
	rm -rf web/pertify/examples
	mkdir -p web/pertify/examples
	cp ./examples/pertify/*.yml ./web/pertify/examples/
	@cd ./web/pertify/examples; for file in ./*.yml; do ( set -e; \
	  export name=`echo $$file | sed s/.yml//`; \
	  set -x; \
	  pertify -f $$file > $$name.dot; \
	  dot -Tsvg $$name.dot > $$name.svg; \
	  dot -Tpng $$name.dot > $$name.png; \
	); done
	cd ./web/pertify/examples; tree -H . --charset=utf-8 > index.html
	ls -la web/pertify/examples

.PHONY: sam-local
sam-local: lambda-build
	@echo ""
	@echo "Open: http://localhost:8080/pertify/index.html"
	@echo ""
	sam local start-api --host=0.0.0.0 --port=8080 --static-dir=web

.PHONY: netlify
netlify: _netlify-deps lambda-build

.PHONY: _netlify-deps
_netlify_deps:
	cd; go get moul.io/fs-bundler
	cd cmd/pertify; $(GO) get -v .

.PHONY: docker
docker:
	docker build \
	  --build-arg VCS_REF=`git rev-parse --short HEAD` \
	  --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	  --build-arg VERSION=`git describe --tags --always` \
	  -t $(DOCKER_IMAGE) .
