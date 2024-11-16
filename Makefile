IMAGE_NAME := go-hello-api
CONTAINER_NAME ?= $(IMAGE_NAME)-container

# Port configuration
HOST_PORT ?= 8080
CONTAINER_PORT := 8080

# Color codes
BLUE := \033[0;34m
GREEN := \033[0;32m
RED := \033[0;31m
NC := \033[0m # No Color

DEFAULT_GOAL := help


.PHONY: help
## Help: Show this help menu
help:
	@echo "Usage:"
	@echo "  make [target] [CONTAINER_NAME=name] [HOST_PORT=port]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  $(BLUE)%-15s$(NC) %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@echo ""

.PHONY: lint
## Lint: Run the Go linter with golangci-lint
lint:
	@echo "$(BLUE)Running linter...$(NC)"
	@out=$$(golangci-lint run); \
	if [ -n "$$out" ]; then \
		echo "$$out"; \
		echo "$(RED)Linter found issues!$(NC)"; \
		exit 1; \
	else \
		echo "$(GREEN)Linter complete - no issues found.$(NC)"; \
	fi
	
.PHONY: build
## Build: Build Go binary
build: lint
	@echo "$(BLUE)Building Go binary$(NC)"
	@go build -o bin/go-hello-api main.go
	@echo "$(GREEN)Build complete: bin/go-hello-api$(NC)"

.PHONY: docker-build
## Docker Build: Build docker image
doker-build:
	@echo "$(BLUE)Building Docker image: $(IMAGE_NAME)$(NC)"
	@docker build -t $(IMAGE_NAME) .
	@echo "$(GREEN)Build complete: $(IMAGE_NAME)$(NC)"

.PHONY: docker-run
## Docker Run: Run container with specified name (use: make run CONTAINER_NAME=my-app)
docker-run:
	@echo "$(BLUE)Starting container: $(CONTAINER_NAME)$(NC)"
	@docker run -d \
		--name $(CONTAINER_NAME) \
		-p $(HOST_PORT):$(CONTAINER_PORT) \
		$(IMAGE_NAME)
	@echo "$(GREEN)Container started: $(CONTAINER_NAME)$(NC)"
	@echo "$(GREEN)Access at: http://localhost:$(HOST_PORT)$(NC)"

.PHONY: docker-stop
## Docker Stop: Stop and remove container
docker-stop:
	@echo "$(BLUE)Stopping container: $(CONTAINER_NAME)$(NC)"
	@docker stop $(CONTAINER_NAME) || true
	@docker rm $(CONTAINER_NAME) || true
	@echo "$(GREEN)Container stopped and removed: $(CONTAINER_NAME)$(NC)"