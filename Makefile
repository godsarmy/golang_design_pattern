
# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: precommit
precommit: clean format lint compile

.PHONY: commit
commit: clean cross-compile
	ls -lha target/

.PHONY: clean
clean:
	rm -rf target

target:
	mkdir target

.PHONY: format
format:
	@gofmt -l -w $(SRC)

.PHONY: lint
lint:
	golangci-lint run

