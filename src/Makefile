.PHONY: test bins clean cover cover_ci
PROJECT_ROOT = github.com/ContinuumLLC/boaring/src

export PATH := $(GOPATH)/bin:$(PATH)

# default target
default: test

PROGS = boaring
TEST_ARG ?= -race -v -timeout 10m
BUILD := ./build
INTEG_TEST_ROOT=./host
INTEG_TEST_DIR=host

vendor/glide.updated: glide.lock glide.yaml
	glide install
	touch vendor/glide.updated

# Automatically gather all srcs
ALL_SRC := $(shell find . -name "*.go" | grep -v -e Godeps -e vendor \
	-e ".*/\..*" \
	-e ".*/_.*" \
	-e ".*/mocks.*")

# all directories with *_test.go files in them
TEST_DIRS := $(sort $(dir $(filter %_test.go,$(ALL_SRC))))

# all tests other than integration test fall into the pkg_test category
PKG_TEST_DIRS := $(filter-out $(INTEG_TEST_ROOT)%,$(TEST_DIRS))

GOCOVERPKG_ARG := -coverpkg="$(PROJECT_ROOT)/common/...,$(PROJECT_ROOT)/service/..."

test: vendor/glide.updated bins
	@rm -f test
	@rm -f test.log
	@for dir in $(TEST_DIRS); do \
		go test -timeout 15m -race -coverprofile=$@ "$$dir" | tee -a test.log; \
	done;

cover_profile:
	@mkdir -p $(BUILD)
	@echo "mode: atomic" > $(BUILD)/cover.out

	@echo Running integration test
	@mkdir -p $(BUILD)/$(INTEG_TEST_DIR)
	@time go test $(INTEG_TEST_ROOT) $(TEST_ARG) $(GOCOVERPKG_ARG) -coverprofile=$(BUILD)/$(INTEG_TEST_DIR)/coverage.out || exit 1;
	@cat $(BUILD)/$(INTEG_TEST_DIR)/coverage.out | grep -v "mode: atomic" >> $(BUILD)/cover.out

	@echo Running package tests:
	@for dir in $(PKG_TEST_DIRS); do \
		mkdir -p $(BUILD)/"$$dir"; \
		go test "$$dir" $(TEST_ARG) -coverprofile=$(BUILD)/"$$dir"/coverage.out || exit 1; \
		cat $(BUILD)/"$$dir"/coverage.out | grep -v "mode: atomic" >> $(BUILD)/cover.out; \
	done;

cover: cover_profile
	go tool cover -html=$(BUILD)/cover.out;

lint: vendor/glide.updated
	@echo Running linter
	@lintFail=0; for file in $(ALL_SRC); do \
		golint "$$file"; \
		if [ $$? -eq 1 ]; then lintFail=1; fi; \
	done; \
	if [ $$lintFail -eq 1 ]; then exit 1; fi;
	@OUTPUT=`gofmt -l $(ALL_SRC) 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "Run 'make fmt'. gofmt must be run on the following files:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

clean:
	rm -f test
	rm -f test.log
	rm -Rf $(BUILD)

fmt:
	@gofmt -w $(ALL_SRC)

	