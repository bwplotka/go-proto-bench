# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.8. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for buf variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(BUF)
#	@echo "Running buf"
#	@$(BUF) <flags/args..>
#
BUF := $(GOBIN)/buf-v1.28.1
$(BUF): $(BINGO_DIR)/buf.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/buf-v1.28.1"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=buf.mod -o=$(GOBIN)/buf-v1.28.1 "github.com/bufbuild/buf/cmd/buf"

COPYRIGHT := $(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60
$(COPYRIGHT): $(BINGO_DIR)/copyright.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=copyright.mod -o=$(GOBIN)/copyright-v0.0.0-20230505153745-6b7392939a60 "github.com/efficientgo/tools/copyright"

GOIMPORTS := $(GOBIN)/goimports-v0.16.0
$(GOIMPORTS): $(BINGO_DIR)/goimports.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/goimports-v0.16.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=goimports.mod -o=$(GOBIN)/goimports-v0.16.0 "golang.org/x/tools/cmd/goimports"

GOLANGCI_LINT := $(GOBIN)/golangci-lint-v1.55.2
$(GOLANGCI_LINT): $(BINGO_DIR)/golangci-lint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golangci-lint-v1.55.2"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=golangci-lint.mod -o=$(GOBIN)/golangci-lint-v1.55.2 "github.com/golangci/golangci-lint/cmd/golangci-lint"

PROTOC_GEN_FASTMARSHAL := $(GOBIN)/protoc-gen-fastmarshal-v0.26.0
$(PROTOC_GEN_FASTMARSHAL): $(BINGO_DIR)/protoc-gen-fastmarshal.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-fastmarshal-v0.26.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-fastmarshal.mod -o=$(GOBIN)/protoc-gen-fastmarshal-v0.26.0 "github.com/CrowdStrike/csproto/cmd/protoc-gen-fastmarshal"

PROTOC_GEN_GO_VTPROTO := $(GOBIN)/protoc-gen-go-vtproto-v0.5.0
$(PROTOC_GEN_GO_VTPROTO): $(BINGO_DIR)/protoc-gen-go-vtproto.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go-vtproto-v0.5.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go-vtproto.mod -o=$(GOBIN)/protoc-gen-go-vtproto-v0.5.0 "github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto"

PROTOC_GEN_GO := $(GOBIN)/protoc-gen-go-v1.31.0
$(PROTOC_GEN_GO): $(BINGO_DIR)/protoc-gen-go.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-go-v1.31.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-go.mod -o=$(GOBIN)/protoc-gen-go-v1.31.0 "google.golang.org/protobuf/cmd/protoc-gen-go"

PROTOC_GEN_GOGOFAST := $(GOBIN)/protoc-gen-gogofast-v1.3.3-0.20221024144010-f67b8970b736
$(PROTOC_GEN_GOGOFAST): $(BINGO_DIR)/protoc-gen-gogofast.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/protoc-gen-gogofast-v1.3.3-0.20221024144010-f67b8970b736"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=protoc-gen-gogofast.mod -o=$(GOBIN)/protoc-gen-gogofast-v1.3.3-0.20221024144010-f67b8970b736 "github.com/gogo/protobuf/protoc-gen-gogofast"

