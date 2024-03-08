# Makefile

BINARY := network

BUILDDIR ?= ./build
PLUGINDIR := ./plugins

GO := go
GOFLAGS := -v

PLUGINTAG := plugin
PLUGINMODE := plugin
PLUGINPATTERN := $(wildcard $(PLUGINDIR)/*)
PLUGINBUILD := $(patsubst $(PLUGINDIR)/%,$(BUILDDIR)/plugins/%.so,$(PLUGINPATTERN))

.PHONY: all clean

all: clean $(BUILDDIR)/$(BINARY) plugins

$(BUILDDIR)/$(BINARY):
	$(GO) build $(GOFLAGS) -o $(BUILDDIR)/$(BINARY)

plugins: $(PLUGINBUILD)

$(BUILDDIR)/plugins/%.so: $(PLUGINDIR)/%
	mkdir -p $(BUILDDIR)/plugins/$*
	
	$(GO) build $(GOFLAGS) -tags $(PLUGINTAG) -buildmode=$(PLUGINMODE) -o $(BUILDDIR)/plugins/$*/plugin.so $</main.go

clean:
	$(RM) $(BUILDDIR)/$(BINARY) $(PLUGINBUILD)