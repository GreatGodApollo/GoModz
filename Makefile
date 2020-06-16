src = $(wildcard plugins/*_mod.go)
obj = $(src:.go=.so)

PLFLAGS = -buildmode=plugin

build:
	@echo "Building executable"
	@go build -o bin/gomodz ./cmd/gomodz/

plugins: $(obj)
	@echo "Building plugins"
	@go build $(PLFLAGS) -o $^ $(src)

all: build plugins

run: all
	@echo "----"
	@echo
	@echo
	@bin/gomodz

%.so: ;

.PHONY: clean
clean:
	@rm -f $(obj) bin/*
	@echo "Built plugins & executables have been deleted."