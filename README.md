# 🍊 Orange Portfolio Backend

Bem-vindo ao repositório do backend do Orange Portfolio! Este projeto é uma api REST desenvolvido utilizando Golang, e tem como objetivo fornecer dados para o frontend do projeto Orange Portfolio.

## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas em seu ambiente de desenvolvimento:

- [Docker](https://www.docker.com/)

## Clonando o Repositório

Para começar, clone o repositório para o seu ambiente local e acesse a pasta do projeto:

```bash
git clone git@github.com:Grupo-38-Orange-Juice/orange-portfolio-back.git
cd orange-portfolio-back
```
## Suba o container docker

Acesse o diretório do projeto, garanta que seu docker esteja rodando e suba o container docker, certifique-se de que as porta 5432 e 8080 estejam disponíveis em seu ambiente local:

```bash
sudo service docker start
docker compose up -d --build
```

## Documentação da API

Após alguns segundos, garante que sua aplicação esteja rodando em http://localhost:8080, e acesse a documentação da API em http://localhost:8080/swagger/index.html. A partir daqui, você pode testar os endpoints da API e verificar os modelos de dados.
Caso a porta 8080 esteja ocupada em seu ambiente local, você pode alterar a porta no arquivo docker-compose.yml.
Caso a aplicação não esteja rodando, você pode verificar o status do container com o comando e executar o passo anterior novamente:

```bash
docker ps
```
## Rodando os testes

Para rodar os testes, acesse o diretório do projeto e execute o seguinte comando:

```bash
go test ./...
```

## Padrões de Projeto Utilizados

- [Clean Architecture (simplificado)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [SOLID](https://en.wikipedia.org/wiki/SOLID)
- [Domain Driven Design (simplificado)](https://en.wikipedia.org/wiki/Domain-driven_design)
- [Repository Pattern](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/ff649690(v=pandp.10)?redirectedfrom=MSDN)
- [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection)

## Camadas da Aplicação
- Domain: Camada responsável por definir e implementar as entidades e as regras de negócio da aplicação e as interfaces de manipulação de dados.
  - entity: Camada responsável por definir as entidades da aplicação
  - usecase: Camada responsável por implementar as regras de negócio da aplicação

- Data: Camada responsável por implementar as interfaces definidas no domínio, realizar manipulação direta de dados
  - http: Camada responsável por implementar as interfaces definidas no domínio, realizar manipulação de dados via http usando o gin
  - postgres: Camada responsável por implementar as interfaces definidas no domínio, realizar manipulação de dados via postgres
  - crypto: Camada responsável por cryptografar e descriptografar dados sensíveis usando bcrypt
  - token: Camada responsável por gerar e validar tokens JWT

## Melhorias Futuras

- Implementar testes de integração
- Desacoplar a camada http de dependências externas
