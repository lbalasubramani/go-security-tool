.PHONY: help build test run dashboard sbom clean

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	go build -o bin/security-tool ./cmd/security-tool

test: ## Run tests
	go test ./...

run: build ## Run the CLI
	./bin/security-tool

dashboard: ## Start the HTMX web dashboard
	go run cmd/security-tool/main.go dashboard

sbom: ## Example: generate and sign SBOM
	go run cmd/security-tool/main.go sbom generate nginx:latest --sign

clean: ## Clean build artifacts
	rm -rf bin/ sbom.json *.sig

lint: ## Run linters (add golangci-lint if desired)
	go vet ./...

# Future targets
# docker-build:
# k8s-deploy:
# trivy-scan: