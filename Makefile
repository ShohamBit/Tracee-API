.PHONY: all | env
# Default target
all: tracee-client

#
#tools
#
CMD_MKDIR ?= mkdir
CMD_GO ?= go

.PHONY: env
env:
	@echo "CMD_MKDIR                $(CMD_MKDIR)"
.PHONY: help
help:

#
# Variables
#

SRC_DIR=./...
#
# output dir
#

OUTPUT_DIR =./dist

$(OUTPUT_DIR):
#
	@$(CMD_MKDIR) -p $@



.PHONY: tracee-client
tracee-client:
	$(CMD_GO) build \
	-o $(OUTPUT_DIR) $(SRC_DIR)
	$(CMD_GO) install $(SRC_DIR)
	

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT_DIR)
