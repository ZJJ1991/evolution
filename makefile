PACKAGE = github.com/JChain/suomi

# FAKE_GOPATH_SUFFIX = $(shell [ -e ".fake_gopath_suffix" ] || date +%s > .fake_gopath_suffix; cat .fake_gopath_suffix)
# FAKE_GOPATH = /tmp/thor-build-$(FAKE_GOPATH_SUFFIX)
# export GOPATH = $(FAKE_GOPATH)

# SRC_BASE = $(FAKE_GOPATH)/src/$(PACKAGE)

JChain:
	@echo "building $@..."
	@go build -v -i -o $(CURDIR)/bin/JChain  $(CURDIR)/chain
	@echo "done. executable created at 'bin/$@'"
