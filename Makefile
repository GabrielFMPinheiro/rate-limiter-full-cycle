.PHONY: build test lint run clean

# Executa os testes
test:
	go test ./... -v

# Executa o aplicativo
run:
	go run main.go
