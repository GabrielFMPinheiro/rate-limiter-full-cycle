## Objetivo:

 - Criar um limitador de requisições por API e IP

## Como rodar ambiente de desenvolvimento:

**Docker:**

 1. Utilizar o comando: docker compose up -d --build

**Terminal:**

 1. Utilizar o comando: docker compose up redis -d
 2. Instale o air na sua máquina para utilizar o mecanismo de hot
    reloading: go install github.com/air-verse/air@latest (NÃO
    OBRIGATÓRIO)
 3. go mod download
 4. Copie e cole o arquivo .env.example e mude o nome para .env
 5. make run (outra opção também é usar o comando: air server --port
    8080 caso tenha o air instalado do passo 1 ou o simples comando go
    run main.go)

## Como rodar testes:

**Docker:**  

 1. Utilizar o comando: docker compose up -d --build      
 2. docker exec-it api make test

  
 **Terminal:**

 1. make test (outra opção também é usar o próprio go test ./... -v)

## Como utilizar o limitador de requisições:

**Por IP:**

 1. Basta disparar para a rota localhost:8080 com método GET
 2. O valor padrão de requisições por segundo são 5
 3. Caso alcançar o limite de requisições o IP é bloqueado por 1 minuto


**Por API Key:**
 1. Basta disparar para a rota localhost:8080 com método GET e header API_KEY e alguma API Key encontrada no arquivo api-key.json, utilizar o valor da propriedade key. 
2. Cada API Key no arquivo api-key.json possui seu número máximo de requisições por segundo
3. Caso alcançar o limite de requisições a API Key é bloqueada por 1 minuto
