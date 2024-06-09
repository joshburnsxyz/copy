GOX := $(shell which go)
BIN := copy
SRC := cmd/copy
OUT := dist
PREFIX := /usr/local

copy:
	@mkdir -p $(OUT)
	$(GOX) build \
		-v \
		-x \
		-o $(OUT)/$(BIN) \

install:
	@cp $(OUT)/$(BIN) $(PREFIX)/bin/$(BIN)

clean:
	@rm -rf $(OUT)

.PHONY: clean install
