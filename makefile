BINARY := net-work
BACKENDDIR := ./
BUILDDIR ?= ./build
PLUGINDIR := $(BACKENDDIR)/plugins

GO := go
GOFLAGS := -v

PLUGINTAG := plugin
PLUGINMODE := plugin
PLUGINPATTERN := $(wildcard $(PLUGINDIR)/*)
PLUGINBUILD := $(patsubst $(PLUGINDIR)/%,$(BUILDDIR)/plugins/%.so,$(PLUGINPATTERN))

.PHONY: all build run clean clean_plugins

all: clean build plugins run

build:
	cd $(BACKENDDIR) && $(GO) build $(GOFLAGS) -o $(abspath $(BUILDDIR))/$(BINARY)

run:
	$(BUILDDIR)/$(BINARY)

plugins: clean_plugins
	mkdir -p $(BUILDDIR)/plugins
	$(MAKE) $(PLUGINBUILD)

$(BUILDDIR)/plugins/%.so: $(PLUGINDIR)/%
	mkdir -p $(BUILDDIR)/plugins/$*
	cd $(BACKENDDIR) && $(GO) build $(GOFLAGS) -tags $(PLUGINTAG) -buildmode=$(PLUGINMODE) -o $(abspath $(BUILDDIR))/plugins/$*/plugin.so $</main.go

clean:
	$(RM) -r $(BUILDDIR)

clean_plugins:
	$(RM) -r $(BUILDDIR)/plugins
