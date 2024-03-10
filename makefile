BINARY := network

BACKENDDIR := ./
BUILDDIR ?= ./build
PLUGINDIR := $(BACKENDDIR)/plugins

GO := go
GOFLAGS := -v

PLUGINTAG := plugin
PLUGINMODE := plugin
PLUGINPATTERN := $(wildcard $(PLUGINDIR)/*)
PLUGINBUILD := $(patsubst $(PLUGINDIR)/%,$(BUILDDIR)/plugins/%.so,$(PLUGINPATTERN))

.PHONY: all clean

all: clean $(BUILDDIR)/$(BINARY) plugins

$(BUILDDIR)/$(BINARY):
	cd $(BACKENDDIR) && $(GO) build $(GOFLAGS) -o $(abspath $(BUILDDIR))/$(BINARY)

plugins: clean_plugins
	mkdir -p $(BUILDDIR)/plugins
	$(MAKE) $(PLUGINBUILD)

$(BUILDDIR)/plugins/%.so: $(PLUGINDIR)/%
	mkdir -p $(BUILDDIR)/plugins/$*
	cd $(BACKENDDIR) && $(GO) build $(GOFLAGS) -tags $(PLUGINTAG) -buildmode=$(PLUGINMODE) -o $(abspath $(BUILDDIR))/plugins/$*/plugin.so $</main.go

clean: clean_plugins

clean_plugins:
	$(RM) -r $(BUILDDIR)/plugins
