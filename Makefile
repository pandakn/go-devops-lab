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

.PHONY: build run stop help

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
	@echo "Examples:"
	@echo "  make run CONTAINER_NAME=my-app"
	@echo "  make run CONTAINER_NAME=my-app HOST_PORT=3000"
	@echo "  make stop CONTAINER_NAME=my-app"

## Build: Build docker image
build:
	@echo "$(BLUE)Building Docker image: $(IMAGE_NAME)$(NC)"
	@docker build -t $(IMAGE_NAME) .
	@echo "$(GREEN)Build complete: $(IMAGE_NAME)$(NC)"

## Run: Run container with specified name (use: make run CONTAINER_NAME=my-app)
run:
	@echo "$(BLUE)Starting container: $(CONTAINER_NAME)$(NC)"
	@docker run -d \
		--name $(CONTAINER_NAME) \
		-p $(HOST_PORT):$(CONTAINER_PORT) \
		$(IMAGE_NAME)
	@echo "$(GREEN)Container started: $(CONTAINER_NAME)$(NC)"
	@echo "$(GREEN)Access at: http://localhost:$(HOST_PORT)$(NC)"

## Stop: Stop and remove container
stop:
	@echo "$(BLUE)Stopping container: $(CONTAINER_NAME)$(NC)"
	@docker stop $(CONTAINER_NAME) || true
	@docker rm $(CONTAINER_NAME) || true
	@echo "$(GREEN)Container stopped and removed: $(CONTAINER_NAME)$(NC)"