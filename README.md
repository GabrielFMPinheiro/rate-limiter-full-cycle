## Objetivo:

- Criar um limitador de requisições por API e IP

## Como rodar ambiente de desenvolvimento:

**Docker:**

1. O arquivo de env utilizado será: .env.docker
2. Utilizar o comando: docker compose up -d --build

**Terminal:**

1.  Copie e cole o arquivo .env.example e mude o nome para .env
2.  Utilizar o comando: docker compose up redis -d
3.  go mod download
4.  make run ou go run main.go

## Como rodar testes:

**Docker:**

1.  O arquivo de env utilizado será: .env.docker
2.  Utilizar o comando: docker compose up -d --build
3.  docker exec -it api make test (outra opção também é usar o próprio go test ./... -v)

**Terminal:**

1.  Utilizar o comando: docker compose up redis -d
2.  make test (outra opção também é usar o próprio go test ./... -v)

## Como utilizar o limitador de requisições:

**Por IP:**

1.  Basta disparar para a rota localhost:8080 com método GET
2.  O valor padrão de requisições por segundo são 5
3.  Caso alcançar o limite de requisições o IP é bloqueado por 1 minuto
4.  Caso queira alterar o número máximo de requisições ou o tempo de bloqueio, basta mudar as envs LIMIT_REQUEST_PER_SECOND_DEFAULT e CACHE_EXPIRATION respectivamente

**Por API Key:**

1. Basta disparar para a rota localhost:8080 com método GET e header API_KEY e alguma API Key encontrada no arquivo api-key.json, utilizar o valor da propriedade key
2. Cada API Key no arquivo api-key.json possui seu número máximo de requisições por segundo
3. Caso alcançar o limite de requisições a API Key é bloqueada por 1 minuto
4. Caso queira alterar o número máximo de requisições ou o tempo de bloqueio, basta mudar as envs LIMITER_REQUEST_PER_SECOND_API_KEY e CACHE_EXPIRATION respectivamente

## Dica:

- Caso queira fazer o teste de carga, existe um arquivo para facilitar o teste, basta executar o arquivo request_script.sh. Você pode alterá-lo para testar todos os cenários desejados.
