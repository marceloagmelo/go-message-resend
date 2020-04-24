# Receber Mensagem usando Golang, RabbitMQ e MySQL

Aplicação ler as mensagens que ainda não foram lidas pelo o **RabbitMQ** e reenvia, esta aplicação utiliza os serviços  [Message API](https://github.com/marceloagmelo/go-message-api) e com acesso ao serviço do RabbitMQ . Esta aplicação possue a seguinte funcionalidade.

- [Ler e Gravar Mensagem](#ler-e-gravar-mensagem)

----

# Instalação

```
go get -v github.com/marceloagmelo/go-message-resend
```
```
cd go-message-resend
```

## Build da Aplicação

```
./image-build.sh
```

## Iniciar a Aplicação
```
./start.sh
```

## Finalizar a Aplicação
```
./stop.sh
```

# Fucionalidades
Lista das funcionalidas:

### Ler e Gravar Mensagem
- Passo 01: A aplicação ler todas as mensagens que não foram lidas pelo RabbitMQ
- Passo 02: A aplicação reenvia as mensagens para o RabbitMQ