.PHONY: generate

generate:
	go generate ./...
	@echo "[OK] Files added to app box!"

security:
	gosec ./...
	@echo "[OK] Go security check completed successfully"

clean:
	go clean ./...
	@echo "[OK] Binary clean up completed"

build: clean test generate security
	go build -o ./bin ./...
	@echo "[OK] App binary was generated successfully"

run:
	@./bin/logify

test:
	reset && go test ./... -cover -short
	@echo "[OK] test execution completed successfully"