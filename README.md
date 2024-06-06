# Simple CRUD and Authentication with Go

Este é um projeto de CRUD simples para cadastro de usuário e autenticação em Go, com utilização de JWT para autenticação e Docker para a configuração do ambiente.

## Funcionalidades

- Cadastro de usuário
- Login de usuário com JWT
- Atualização de dados do usuário
- Exclusão de usuário

## Pré-requisitos

Certifique-se de ter instalado em sua máquina:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Configuração

Para configurar o projeto, crie o arquivo `.env` na raiz do projeto e defina as variáveis de ambiente necessárias. (Você irá encontra-las no arquivo `.env.example`).

## Instalação e Uso

1. Clone o repositório:

```bash
git clone https://github.com/pedro-costa22/first-crud-go.git 
```

2. Navegue até o diretório do projeto:

```bash
cd first-crud-go
```

3. Inicie o ambiente Docker:

```bash
docker-compose up -d
```


