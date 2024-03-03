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

all: $(BUILDDIR)/$(BINARY) plugins

$(BUILDDIR)/$(BINARY):
	$(GO) build $(GOFLAGS) -o $(BUILDDIR)/$(BINARY)

plugins: $(PLUGINBUILD)

$(BUILDDIR)/plugins/%.so: $(PLUGINDIR)/%
	$(GO) build $(GOFLAGS) -tags $(PLUGINTAG) -buildmode=$(PLUGINMODE) -o $@ $</main.go

clean:
	$(RM) $(BUILDDIR)/$(BINARY) $(PLUGINBUILD)