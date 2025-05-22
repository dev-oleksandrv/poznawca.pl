APP_PROXY_SRC := ./cmd/app-proxy/main.go
APP_PROXY_OUT := ./bin/app-proxy
BACKOFFICE_PROXY_SRC := ./cmd/backoffice-proxy/main.go
BACKOFFICE_PROXY_OUT := ./bin/backoffice-proxy
GO_CMD := go

run-app: check
	@echo "Running the application..."
	$(GO_CMD) run $(APP_PROXY_SRC)

run-backoffice: check
	@echo "Running the application..."
	$(GO_CMD) run $(BACKOFFICE_PROXY_SRC)

vet:
	@echo "Running go vet..."
	$(GO_CMD) vet ./...

fmt:
	@echo "Formatting code..."
	$(GO_CMD) fmt ./...

test:
	@echo "Running tests..."
	$(GO_CMD) test ./... -v

build-app:
	@echo "Building the application..."
	$(GO_CMD) build -o $(APP_PROXY_OUT) $(APP_PROXY_SRC)

build-backoffice:
	@echo "Building the application..."
	$(GO_CMD) build -o $(BACKOFFICE_PROXY_OUT) $(BACKOFFICE_PROXY_SRC)

clean:
	@echo "Cleaning up..."
	@rm -f $(OUT)

deps:
	@echo "Installing dependencies..."
	$(GO_CMD) mod tidy

check: vet fmt test

.PHONY: all run-app vet fmt test build-app clean deps check run-backoffice build-backoffice