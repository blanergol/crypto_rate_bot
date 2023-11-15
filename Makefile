generate: ## Code generation
	@make gen_logs
	@make imports

.PHONY: fmt
fmt: ## Format source using gofmt
	@gofumpt -l -w .

imports: ## fix go imports
	@goimports -local github.com/blanergol/crypto_rate_bot -w -l ./

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)