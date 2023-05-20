.PHONY: doc doc-build doc-serve doc-clean

HUGO ?= hugo
DOC_DIR ?= doc
DOC_OUT ?= $(DOC_DIR)/public
HUGO_BUILD_FLAGS ?= --buildDrafts
HUGO_SERVER_FLAGS ?= --buildDrafts --disableFastRender

doc: doc-build

doc-build:
	$(HUGO) --source $(DOC_DIR) --destination $(abspath $(DOC_OUT)) $(HUGO_BUILD_FLAGS)

doc-serve:
	$(HUGO) server --source $(DOC_DIR) $(HUGO_SERVER_FLAGS)

doc-clean:
	$(RM) -r $(DOC_OUT) $(DOC_DIR)/resources
