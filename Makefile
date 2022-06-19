.PHONY: build-static
build-static:
	@echo "Building static library..."
	@rm -rf frontend/build
	@rm -rf backend/web
	@mkdir -p backend/web
	@npm --prefix ./frontend/ run build
	@cp -r frontend/build/. backend/web/

.PHONY: run
run:
	@echo "Running..."
	@cd backend && go run cmd/app/main.go


.PHONY: generate
generate:
	@echo "Generating..."
	@cd backend && go generate ./...

.PHONY: re-build
re-build: build-static run
