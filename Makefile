.PHONY: all dev fmt lint test integration stress clean

BINARY_NAME=motor-seo

all: fmt lint test integration

dev:
	go run main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

# Corre solo pruebas unitarias puras en memoria, sin colgarse por falta de red
test:
	go test -v -race -cover -tags=!integration ./...

# Corre las pruebas que requieren el tag de integración
integration:
	go test -v -tags=integration ./tests/...

stress:
	chmod +x stress_test.sh
	./stress_test.sh

clean:
	rm -rf bin/