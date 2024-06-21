### Event Reservation API

Esta é uma API em Go para reserva de lugares em eventos. A aplicação utiliza um banco de dados simulado em JSON para gerenciar os eventos e os lugares disponíveis para reserva.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado em sua máquina:

- Go (versão 1.16 ou superior)
- Git

## Instalação

Clone o repositório:

```bash
git clone https://github.com/seu-usuario/event-reservation-api.git
```

## Instale as dependências:

```bash
go mod tidy
```

## Execução

Dentro da pasta do projeto ir para:

```bash
cd cmd/events
```

Para executar a aplicação, utilize o seguinte comando:

```bash
go run main.go
```

A aplicação será executada em http://localhost:8000 por padrão.

## Utilização

Endpoints Disponíveis

- GET /events: Lista todos os eventos disponíveis.
- GET /events/{eventID}: Obtém detalhes de um evento específico.
- GET /events/{eventID}/spots: Lista os lugares disponíveis para um evento.
- POST /events/{eventID}/reserve: Reserva um lugar para um evento.
