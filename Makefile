MAKEFILE_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
WORKFLOW_DIR = ~/Dropbox/Alfred/Alfred.alfredpreferences/workflows
BUNDLE_ID = com.ruedap.font-awesome
FAW_EXEC_CMD = FAW_ICONS_YAML_PATH=workflow/icons.yml ./workflow/faw

ci: build exec test
exec:
	@$(FAW_EXEC_CMD) find apple
	@echo
	@$(FAW_EXEC_CMD) put -name apple
	@echo
	@$(FAW_EXEC_CMD) put -code apple
	@echo
	@$(FAW_EXEC_CMD) put -ref apple
	@echo
	@$(FAW_EXEC_CMD) put -url apple
	@echo
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
