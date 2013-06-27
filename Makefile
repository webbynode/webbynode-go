.PHONY: all

all:
	@echo all

fmt:
	@gofmt -s -l -w -tabs=false -tabwidth=2 .
