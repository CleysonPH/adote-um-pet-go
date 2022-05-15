# Adote um Pet

Imeplentação feita em Go da API Adote um Pet desenvolida no workshop Multi-stack da TreinaWeb.

## Requisitos

- Go >= 1.17

## Pacotes utilizados

- Gin 1.7.7
- Gorm 1.23.5

## Como começar

Clone o repositório e entre na pasta do projeto
```sh
git clone https://github.com/CleysonPH/adote-um-pet-go.git
cd adote-um-pet-go
```

Instale as dependências do projeto
```sh
go mod tidy
```

Levante o banco de dados no Docker
```sh
docker-compose up -d
```

E por fim execute o projeto
```sh
go run main.go
```

Então acesse a aplicação em [http://localhost:8000](http://localhost:8000)

## Rotas disponívels

| Métodos HTTP | Rota         |
|--------------|--------------|
| GET, POST    | /api/pets    |
| GET, POST    | /api/adocoes |

## Estrutura do projeto

```
.
├── README.md
├── api
│   ├── adoption
│   │   ├── controllers
│   │   │   ├── adoption.controller.go
│   │   │   └── controllers.go
│   │   ├── dtos
│   │   │   └── adoption.dto.go
│   │   ├── mappers
│   │   │   ├── adoption.mapper.go
│   │   │   └── mappers.go
│   │   └── validators
│   │       ├── adoption.validator.go
│   │       └── validator.go
│   └── pet
│       ├── controllers
│       │   ├── controllers.go
│       │   └── pet.controller.go
│       ├── dtos
│       │   └── pet.dto.go
│       └── mappers
│           ├── mappers.go
│           └── pet.mapper.go
├── core
│   ├── config
│   │   └── validator.go
│   ├── database
│   │   └── postgres.database.go
│   ├── models
│   │   ├── adoption.model.go
│   │   └── pet.model.go
│   ├── repositories
│   │   ├── adoption.repository.go
│   │   ├── pet.repository.go
│   │   └── repositories.go
│   └── utils
│       └── utils.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```
