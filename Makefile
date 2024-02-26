.PHONY: build
build:
	@echo "🚀 Building gtc"
	@go build -o gtc ./cmd/gtc

.PHONY: install
install:
	@echo "📦 Installing gtc"
	@go install ./cmd/gtc

.PHONY: watch-install
watch-install:
	@cng -ik './cmd/**/*.go' -- make build install

.PHONY: test
test:
	@echo "🧪 Testing..."
	@go test ./...

.PHONY: watch-test
watch-test:
	@echo "🧪 Watching tests..."
	@cng -ik '**/*.go' -- make test

.DEFAULT_GOAL := watch-test