NAME := route53-update
VERSION := v0.0.1

build:
	go build -ldflags "-X main.Version=$(VERSION)"

compile:
	@rm -rf build/
	@gox -ldflags "-X main.Version=$(VERSION)" \
	-osarch="darwin/amd64" \
	-osarch="linux/amd64" \
	-osarch="freebsd/amd64" \
	-output "build/{{.Dir}}_$(VERSION)_{{.OS}}_{{.Arch}}/$(NAME)"

install:
	go install -ldflags "-X main.Version=$(VERSION)"

deps:
	go get github.com/mitchellh/gox

dist: compile
	$(eval FILES := $(shell ls build))
	@rm -rf dist && mkdir dist
	@for f in $(FILES); do \
		(cd $(shell pwd)/build/$$f && gzip -9 -c route53-update > ../../dist/$$f.gz); \
		(cd $(shell pwd)/dist && shasum -a 512 $$f.gz > $$f.sha512); \
		(cd $(shell pwd)/dist && gpg --armor --output $$f.asc --detach-sig $$f.gz); \
		echo $$f; \
	done

.PHONY: build compile install deps dist
