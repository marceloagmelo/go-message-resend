#!/usr/bin/env bash

source setenv.sh

# Rabbitmq message api
echo "Subindo o go-message-resend..."
docker run -d --name go-message-resend --network message-net  \
-p 7070:8080 \
-e API_SERVICE_URL="http://go-message-api:8080" \
-e TZ=America/Sao_Paulo \
marceloagmelo/go-message-resend

# Listando os containers
docker ps
