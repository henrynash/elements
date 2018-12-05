SHELL=/bin/bash

AN_DIRS?=an

INPUT_DIRS=$(AN_DIRS)
PACKAGE=$(shell go list .)
ELEMENT_PACKAGE=repos.antha.com/elements

# Compile after downloading dependencies
all: compile_with_deps fmt_json fmt_go

# Compile using current state of working directories
current: compile fmt_json fmt_go

clean:
	rm -rf .staging

test: all
	go test -v `go list ./... | grep -v vendor`

fmt_json:
	go run cmd/format-json/main.go -inPlace $(INPUT_DIRS)
	
fmt_go:
	gofmt -w $(INPUT_DIRS)

compile:
	rm -rf "vendor/$(ELEMENT_PACKAGE)"
	mkdir -p .staging
	go build -o .staging/antha-s1 github.com/antha-lang/antha/cmd/antha
	./.staging/antha-s1 format -w $(AN_DIRS)
	./.staging/antha-s1 compile \
	  --outdir=vendor/$(ELEMENT_PACKAGE) \
	  --outputPackage $(ELEMENT_PACKAGE) \
	  $(AN_DIRS)
	go install $(PACKAGE)/cmd/antha

compile_with_deps:
	go get -d ./cmd/... || true
	rm -rf "vendor/$(ELEMENT_PACKAGE)"
	mkdir -p .staging
	go build -o .staging/antha-s1 github.com/antha-lang/antha/cmd/antha
	./.staging/antha-s1 format -w $(AN_DIRS)
	./.staging/antha-s1 compile \
	  --outdir=vendor/$(ELEMENT_PACKAGE) \
	  --outputPackage $(ELEMENT_PACKAGE) \
	  $(AN_DIRS)
	go get -f -u -d ./cmd/... || true
	go install $(PACKAGE)/cmd/antha
