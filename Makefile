GOX := $(shell which goreleaser)
BIN := copy
SRC := cmd/copy
OUT := dist
PREFIX := /usr/local

release-local:
	$(GOX) release --snapshot --clean

release:
	$(GOX) release

clean:
	@rm -rf $(OUT)

.PHONY: clean install
