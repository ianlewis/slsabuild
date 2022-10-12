REVISION_ID ?= $(shell git rev-parse --verify HEAD)
VERSION = $(subst /,_,$(subst heads/,,$(patsubst v%,%,$(subst tags/,,$(shell if [ "${TAG_NAME}" == "" ]; then git describe --match "v*" --all --long --dirty --broken --always; else echo ${TAG_NAME}-0-${SHORT_SHA}; fi)))))
RELEASE = "v$(call NORMALIZE,$(VERSION))"


.PHONY: help
help: ## Shows all targets and help from the Makefile (this message).
	@echo "slsabuild Makefile"
	@echo "Usage: make [COMMAND]"
	@echo ""
	@grep --no-filename -E '^([/a-z.A-Z0-9_%-]+:.*?|)##' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = "(:.*?|)## ?"}; { \
			if (length($$1) > 0) { \
				printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2; \
			} else { \
				if (length($$2) > 0) { \
					printf "%s\n", $$2; \
				} \
			} \
		}'

slsabuild: ## Build the slsabuild binary.
	VERSION=$(VERSION) go run ./*.go run

clean: ## Clean up
	rm -f slsabuild
	rm -f slsabuild.intoto.jsonl
