.PHONY: build test lint run clean

# Executa os testes
test:
	go test ./...

# Executa o aplicativo
run:
	air server --port 8080
