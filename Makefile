.PHONY: build
build:
	@echo "🚀 Building gtc"
	@go build -o ./dist/gtc .

.PHONY: install
install:
	@echo "📦 Installing gtc"
	@go install .

.PHONY: watch-install
watch-install:
	@cng -ik '*.go' -- make build install

.PHONY: test
test:
	@echo "🧪 Testing..."
	@go test ./...

.PHONY: watch-test
watch-test:
	@echo "🧪 Watching tests..."
	@cng -ik '**/*.go' -- make test

.DEFAULT_GOAL := watch-test