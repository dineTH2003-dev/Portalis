.PHONY: all help test lint build run-server run-gateway run-agent run-client clean

all: help

help:
	@echo "Portalis Development Automation"
	@echo "==============================="
	@echo "Available commands:"
	@echo "  make test         - Run test suites across all modules"
	@echo "  make lint         - Check Go formatting and TypeScript types"
	@echo "  make build        - Build all Go binaries and client bundle"
	@echo "  make run-server   - Start Control Plane in development mode (Bun)"
	@echo "  make run-gateway  - Start Go Ingress Gateway"
	@echo "  make run-agent    - Start Go Agent CLI"
	@echo "  make run-client   - Start React Dashboard in dev mode (Vite)"
	@echo "  make clean        - Remove build artifacts and temporary files"

test:
	@if [ -f "frames/go.mod" ]; then echo "==> Testing frames..."; cd frames && go test -v ./...; fi
	@if [ -f "gateway/go.mod" ]; then echo "==> Testing gateway..."; cd gateway && go vet ./...; fi
	@if [ -f "agent/go.mod" ]; then echo "==> Testing agent..."; cd agent && go vet ./...; fi
	@if [ -f "apps/server/package.json" ]; then echo "==> Testing server..."; cd apps/server && bun test; fi

lint:
	@echo "==> Checking Go formatting..."
	@UNFORMATTED=$$(gofmt -l frames gateway agent 2>/dev/null || true); \
	if [ -n "$$UNFORMATTED" ]; then \
		echo "Files requiring gofmt:"; echo "$$UNFORMATTED"; exit 1; \
	else \
		echo "Go formatting is clean."; \
	fi
	@if [ -f "apps/server/package.json" ]; then \
		echo "==> Typechecking Control Plane..."; \
		cd apps/server && bun run typecheck; \
	fi

build:
	@echo "==> Building Gateway..."
	@mkdir -p gateway/bin && cd gateway && go build -o bin/portalis-gateway ./cmd/gateway
	@echo "==> Building Agent CLI..."
	@mkdir -p agent/bin && cd agent && go build -o bin/portalis ./cmd/agent
	@echo "==> Building React Dashboard..."
	@cd apps/client && bun run build

run-server:
	@cd apps/server && bun run dev

run-gateway:
	@cd gateway && go run ./cmd/gateway

run-agent:
	@cd agent && go run ./cmd/agent

run-client:
	@cd apps/client && bun run dev

clean:
	@rm -rf gateway/bin agent/bin apps/client/dist apps/server/*.db*
