.PHONY: help 
all: help

wc: ## Run word count with time
	time wc -w data.txt

build: 
	go build -o words

gc: build ## Run go program count with time
	time ./words data.txt

cpu: ## Run pprof view for cpu profile
	go tool pprof -http=:8080 cpu.pprof

mem: ## Run pprof view for mem profile
	go tool pprof -http=:8080 mem.pprof

trace: ## Run trace view for trace profile
	go tool trace trace.out

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'