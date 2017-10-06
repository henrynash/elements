SHELL=/bin/bash

AN_DIRS?=an

INPUT_DIRS=$(AN_DIRS) defaults
PACKAGE=$(shell go list .)
ELEMENT_PACKAGE=repos.antha.com/elements

# Compile after downloading dependencies
all: update_deps fmt_json compile

# Compile using current state of working directories
current: fmt_json compile

clean:
	rm -rf .staging

test: all
	go test -v `go list ./... | grep -v vendor`

fmt_json:
	go run cmd/format-json/main.go -inPlace $(INPUT_DIRS)

update_deps:
	mkdir -p .staging
	go build -o .staging/antha-s1 ./vendor/github.com/antha-lang/antha/cmd/antha
	go get -f -u -d -v ./cmd/... || true
	./.staging/antha-s1 compile \
	  --outdir=vendor/$(ELEMENT_PACKAGE) \
	  --outputPackage $(ELEMENT_PACKAGE) \
	  $(AN_DIRS)
	go get -f -u -d -v ./cmd/... || true

antha:
	mkdir -p .staging
	go build -o .staging/antha-s1 ./vendor/github.com/antha-lang/antha/cmd/antha

compile: gen_comp
	go install $(PACKAGE)/cmd/antha

gen_comp: antha
	rm -rf "vendor/$(ELEMENT_PACKAGE)"
	./.staging/antha-s1 format -w $(AN_DIRS)
	./.staging/antha-s1 compile \
	  --outdir=vendor/$(ELEMENT_PACKAGE) \
	  --outputPackage $(ELEMENT_PACKAGE) \
	  $(AN_DIRS)
