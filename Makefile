.PHONY: setup build install

help:
	@echo "Usage:"
	@echo "  make setup   # install required modules"
	@echo "  make build   # gobuild build"

setup:
	go get gopkg.in/yaml.v2

build:
	go build

install:
	go install
