# Go Stress Test CLI

Este projeto implementa um sistema de linha de comando (CLI) escrito em Go para realizar testes de carga em um serviço web. Ele permite que você forneça a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas. Após a execução dos testes, o sistema gera um relatório com informações específicas sobre o desempenho.

## Estrutura do Projeto

```
stress-test/
├── Dockerfile
├── go.mod
├── go.sum
└── cmd/
    └── main.go
```

## Funcionalidades

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.
- Gerar um relatório ao final dos testes contendo:
  - Tempo total gasto na execução.
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Pré-requisitos

- Docker instalado na sua máquina.

## Como rodar o projeto

### 1. Clonar o Repositório

Clone o repositório do projeto para a sua máquina local.

```sh
git clone https://github.com/rzeradev/stress-test-cli
cd stress-test-cli
```

### 2. Construir a Imagem Docker

Construa a imagem Docker usando o Dockerfile fornecido.

```sh
docker build -t stress-test .
```

### 3 Executar o Teste de Carga

Execute o container Docker com os parâmetros desejados. Por exemplo:

```sh
docker run stress-test --url=http://google.com --requests=1000 --concurrency=10
```

### Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.
- `--logs`: true/false (opcional). Se definido como `true`, exibirá os logs de cada request.

### Exemplo de Uso

```sh
docker run stress-test --url=http://example.com --requests=500 --concurrency=20
```

## Relatório

Ao final da execução dos testes, o programa gera um relatório com as seguintes informações:

- **Tempo total gasto na execução**: O tempo total que levou para concluir todos os requests.
- **Quantidade total de requests realizados**: O número total de requests que foram feitos.
- **Quantidade de requests com status HTTP 200**: O número de requests que retornaram com status HTTP 200 (OK).
- **Distribuição de outros códigos de status HTTP**: A quantidade de outros códigos de status HTTP recebidos (por exemplo, 404, 500, etc.).

## Arquivo `Dockerfile`

```dockerfile
FROM golang:1.21.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/

RUN go build -o /stress-test ./cmd/main.go

ENTRYPOINT ["/stress-test"]
```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
