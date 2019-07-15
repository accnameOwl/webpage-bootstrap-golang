GO := go
GORUN := $(GO) run
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test
GOGET := $(GO) get
ENTRY_FILE := main.go
PR_DIR := ./

CPU_PR_FILE := ./tmp/compile_cpu-profile.txt
MEMFILE := ./tmp/compile_mem-profile.txt
FULLBUILD:= -cpuprofile $(CPU_PR_FILE) -memprofile $(MEMFILE)

build:
	$(GOBUILD) $(ENTRY_FILE)

fbuild:
	$(GOBUILD) $(ENTRY_FILE) $(FULLBUILD)

test:
	$(GOTEST) $(ENTRY_FILE)

run:
	$(GORUN) $(ENTRY_FILE)