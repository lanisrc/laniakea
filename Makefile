.PHONY: build check clean test vet doc doc-build doc-serve doc-clean

GO ?= go
GO_OUT ?= bin
GO_CACHE ?= .cache/go-build

HUGO ?= hugo
DOC_DIR ?= doc
DOC_OUT ?= $(DOC_DIR)/public
HUGO_BUILD_FLAGS ?= --buildDrafts
HUGO_SERVER_FLAGS ?= --buildDrafts --disableFastRender

build:
	GOCACHE=$(abspath $(GO_CACHE)) $(GO) build -o $(GO_OUT)/laniakea ./cmd/laniakea

test:
	GOCACHE=$(abspath $(GO_CACHE)) $(GO) test ./...

vet:
	GOCACHE=$(abspath $(GO_CACHE)) $(GO) vet ./...

fmt:
	@test -z "$$($(GO)fmt -l $$(find . -name '*.go' -not -path './.cache/*'))"

check: test vet

clean:
	$(RM) -r $(GO_OUT)

doc: doc-build

doc-build:
	$(HUGO) --source $(DOC_DIR) --destination $(abspath $(DOC_OUT)) $(HUGO_BUILD_FLAGS)

doc-serve:
	$(HUGO) server --source $(DOC_DIR) $(HUGO_SERVER_FLAGS)

doc-clean:
	$(RM) -r $(DOC_OUT) $(DOC_DIR)/resources
