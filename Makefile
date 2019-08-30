GOPKG ?=	moul.io/graphman
DOCKER_IMAGE ?=	moul/graphman
GOBINS ?=	./cmd/pertify

all: test install

-include rules.mk

web/pertify/examples.js: $(wildcard ./examples/pertify/*.yml)
	cd examples/pertify; fs-bundler --format=js --callback=examples *.yml > ../../$@

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
