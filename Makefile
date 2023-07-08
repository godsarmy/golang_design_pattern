
# go source files, ignore vendor directory
PATTERNS = $(shell find cmd -mindepth 1 -maxdepth 1 -type d -printf '%f\n')
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

GO := go
GOFMT := gofmt
GO_GCFLAGS =
GO_BUILD_FLAGS= 
GO_LDFLAGS = -ldflags "-s -w"

BINARIES = $(addprefix target/,$(PATTERNS))

FORCE: 

define BUILD_BINARY
@echo "Build $@"
$(GO) build ${GO_GCFLAGS} ${GO_BUILD_FLAGS} -o $@ ${GO_LDFLAGS} ./$<
endef

target/%: cmd/% FORCE
	$(call BUILD_BINARY)

.PHONY: precommit
precommit: clean format lint compile

.PHONY: commit
commit: clean compile
	ls -lha target/

.PHONY: clean
clean:
	rm -rf target

target:
	mkdir target

.PHONY: compile
compile: target $(BINARIES) ## build binaries
	@echo "Build $@"

.PHONY: format
format:
	@$(GOFMT) -l -w $(SRC)

.PHONY: lint
lint:
	golangci-lint run

