#!/usr/bin/make -f
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

build-linux: go.sum
	GOOS=linux GOARCH=amd64 $(MAKE) build

build-arm32:go.sum
	GOOS=linux GOARCH=arm GOARM=7 $(MAKE) build

build-arm64: go.sum
	GOOS=linux GOARCH=arm64 $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/