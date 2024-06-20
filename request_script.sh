#!/bin/bash

# Número total de requisições
total_requests=11

# URL do endpoint
endpoint="http://localhost:8080/"

# Header API_KEY
api_key="4uT1wVB0PLEnGwkT5gl1PZBjd4HCuJA61qYf7FYEORNk95ecp8ixXwTCgsHFvcMLJP9qCyxmHAaKMGC33RxxmSRWIeZujWU51506GmfAR8blDzk7TN5GixXguvx3E2WZ"

# Loop para realizar as requisições com intervalo de 300ms
for ((i = 1; i <= total_requests; i++)); do
    # Realiza a requisição HTTP usando curl com o header API_KEY
    curl -X GET "$endpoint" -H "API_KEY: $api_key" -w "\n" &
    
    # Aguarda 300ms (0.3 segundos) antes da próxima requisição
    sleep 0
done
