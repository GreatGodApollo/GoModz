src = $(wildcard plugins/*_mod.go)
obj = $(src:.go=.so)

PLFLAGS = -buildmode=plugin

build:
	@go build -o bin/gomodz ./cmd/gomodz/

plugins: $(obj)
	@go build $(PLFLAGS) -o $^ $(src)

%.so: ;

.PHONY: clean
clean:
	@rm -f $(obj) bin/*
	@echo "Built plugins & executables have been deleted."