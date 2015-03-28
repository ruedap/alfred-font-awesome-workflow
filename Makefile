MAKEFILE_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
WORKFLOW_DIR = ~/Dropbox/Alfred/Alfred.alfredpreferences/workflows
BUNDLE_ID = com.ruedap.font-awesome

ci: build test
build:
	go build -o ./workflow/faw
test:
	go test -v
clean:
	- rm ./workflow/faw
link:
	- ln -s $(MAKEFILE_DIR)/workflow $(WORKFLOW_DIR)/$(BUNDLE_ID)
unlink:
	- rm $(WORKFLOW_DIR)/$(BUNDLE_ID)
	- rm ./workflow/workflow

.PHONY: ci build test clean link unlink
