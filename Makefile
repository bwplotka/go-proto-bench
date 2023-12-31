include .bingo/Variables.mk
FILES_TO_FMT   ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)
PROTO_VERSIONS ?= $(shell ls ./proto/prw)

PB_VERSIONS ?= v2testvtproto v2testcsproto #v1testgogo v2testgogo-opt v2 #v2testgogo-proto might be not interesting at this point.
BENCHMARKS ?= BenchmarkPRWSerialize BenchmarkPRWDeserialize BenchmarkPRWDeserializeToBase

BENCH_RESULT_DIR ?= ./results

GOBIN ?= $(firstword $(subst :, ,${GOPATH}))/bin
PATH := $(PATH):$(GOBIN)

# Tools.
GIT  ?= $(shell which git)

# Support gsed on OSX (installed via brew), falling back to sed. On Linux
# systems gsed won't be installed, so will use sed as expected.
SED ?= $(shell which gsed 2>/dev/null || which sed)

define require_clean_work_tree
	@git update-index -q --ignore-submodules --refresh

    @if ! git diff-files --quiet --ignore-submodules --; then \
        echo >&2 "cannot $1: you have unstaged changes."; \
        git diff-files --name-status -r --ignore-submodules -- >&2; \
        echo >&2 "Please commit or stash them."; \
        exit 1; \
    fi

    @if ! git diff-index --cached --quiet HEAD --ignore-submodules --; then \
        echo >&2 "cannot $1: your index contains uncommitted changes."; \
        git diff-index --cached --name-status -r --ignore-submodules HEAD -- >&2; \
        echo >&2 "Please commit or stash them."; \
        exit 1; \
    fi

endef

help: ## Displays help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: all
all: format

.PHONY: format
format: ## Formats Go code including imports and proto
format: $(GOIMPORTS) $(BUF)
	@echo ">> formatting code"
	@$(GOIMPORTS) -w $(FILES_TO_FMT)
	@echo ">> formatting proto"
	@$(BUF) format -w ./proto

.PHONY: proto
proto: ## Regenerate Go from proto
proto: $(BUF) $(PROTOC_GEN_GOGOFAST) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_VTPROTO) $(PROTOC_GEN_FASTMARSHAL)
	@for version in $(PROTO_VERSIONS); do \
    	echo ">> regenerating $$version" ; \
    	$(BUF) generate --template proto/prw/$$version/buf.gen.yaml --path proto/prw/$$version proto ; \
	done

.PHONY: check-git
check-git:
ifneq ($(GIT),)
	@test -x $(GIT) || (echo >&2 "No git executable binary found at $(GIT)."; exit 1)
else
	@echo >&2 "No git binary found."; exit 1
endif

VER_EXTRA ?=
PB ?= v2
BENCH_NAME ?= BenchmarkPRWSerialize
VER := $(PB)$(VER_EXTRA)
.PHONY: bench
bench:
	@mkdir -p $(BENCH_RESULT_DIR)/$(BENCH_NAME)
	@echo ">> benchmarking $(VER)"
	go test ./prw -run '^$$' -bench '^$(BENCH_NAME)$$' -benchtime 5s -count 6 \
		--proto $(PB) \
 		-cpu 4 \
 		-benchmem \
 		-memprofile=$(BENCH_RESULT_DIR)/$(BENCH_NAME)/$(VER).mem.pprof -cpuprofile=$(BENCH_RESULT_DIR)/$(BENCH_NAME)/$(VER).cpu.pprof \
 		| tee $(BENCH_RESULT_DIR)/$(BENCH_NAME)/$(VER).txt;

.PHONY: bench-all
bench-all:
	@for version in $(PB_VERSIONS); do \
		for bench in $(BENCHMARKS); do \
        		echo ">> benchmarking $$bench for $$version" ; \
        		$(MAKE) bench PB=$$version BENCH_NAME=$$bench ; \
        	done \
	done

.PHONY: cmp
cmp: $(BENCHSTAT)
	@cd $(BENCH_RESULT_DIR) && for bench in $(BENCHMARKS); do \
  		cd $$bench ; \
		echo ">> comparing $$bench" ; \
		$(BENCHSTAT) *.txt ; \
		cd .. ; \
	done

# PROTIP:
# Add
#      --cpu-profile-path string   Path to CPU profile output file
#      --mem-profile-path string   Path to memory profile output file
# to debug big allocations during linting.
lint: ## Runs various static analysis against our code.
lint: $(GOLANGCI_LINT) $(COPYRIGHT) format check-git
	$(call require_clean_work_tree,"detected not clean master before running lint")
	@echo ">> linting all of the Go files"
	@$(GOLANGCI_LINT) run
	@echo ">> linting proto"
	@$(BUF) lint ./proto
	@echo ">> ensuring Copyright headers"
	@$(COPYRIGHT) $(FILES_TO_FMT)
	$(call require_clean_work_tree,"detected white noise or/and files without copyright; run 'make lint' file and commit changes.")
