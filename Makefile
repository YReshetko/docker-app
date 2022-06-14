.PHONY: build-static
build-static:
	@echo "Building static library..."
	@rm -rf frontend/build
	@rm -rf backend/resources
	@mkdir -p backend/resources
	@npm --prefix ./frontend/ run build
	@cp -r frontend/build/. backend/resources/

.PHONY: run
run:
	@echo "Running..."
	@go run backend/main.go