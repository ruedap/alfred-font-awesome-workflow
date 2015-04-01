DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
MAKEFILE_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
WORKFLOW_DIR = ~/Dropbox/Alfred/Alfred.alfredpreferences/workflows
BUNDLE_ID = com.ruedap.font-awesome
FAW_CLI_CMD = FAW_ICONS_YAML_PATH=workflow/icons.yml ./workflow/faw

default: build cli test

ci: deps build cli coveralls

cli:
	@echo "--> Running CLI commands"
	@$(FAW_CLI_CMD) find apple
	@echo
	@$(FAW_CLI_CMD) put -name apple
	@echo
	@$(FAW_CLI_CMD) put -code apple
	@echo
	@$(FAW_CLI_CMD) put -ref apple
	@echo
	@$(FAW_CLI_CMD) put -url apple
	@echo

coveralls:
	@echo "--> Testing packages and sending coverage report"
	@goveralls -v -service=travis-ci

deps:
	@echo "--> Installing build dependencies"
	go get -d ./... $(DEPS)
	go get github.com/axw/gocov/gocov
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

build:
	@echo "--> Compiling packages and dependencies"
	@mkdir -p ./workflow/
	go build -ldflags '-s -w' -o ./workflow/faw

cov:
	@echo "--> Reporting coverages"
	@gocov test | gocov report


test:
	@echo "--> Testing packages"
	@go test -v -cover

bench:
	@echo "--> Benchmarking packages"
	@go test -bench .

clean:
	@echo "--> Cleaning workflow files"
	@- rm ./workflow/faw

link:
	@echo "--> Linking workflow files"
	@- ln -s $(MAKEFILE_DIR)/workflow $(WORKFLOW_DIR)/$(BUNDLE_ID)

unlink:
	@echo "--> Unlinking workflow files"
	@- rm $(WORKFLOW_DIR)/$(BUNDLE_ID)
	@- rm ./workflow/workflow

release: clean build unlink link

.PHONY: default ci cli cov coveralls deps build test clean link unlink release
