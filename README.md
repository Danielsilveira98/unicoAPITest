# Unico API Test

Essa API tem como objetivo administrar os eventos de feira de uma determinada cidade. Nela é possível criar, editar, excluir e listar feiras.

## Requisitos
- Golang 1.18+
- Make

## Setup do projeto
Rodar `make setup`

## Cobertura de testes
Rodar `make test`

## Rodando a API
### Via docker-compose
### Via build
Rodar `make run`

## Populando base para testes
É possível inserir uma base de dados inicias com dados de feiras da cidade de São Paulo, que são fornecidas via CSV no site da mesma. A massa de dados porém, apresenta alguns problemas de encoding além de derem que sofre alguns ajustes.

Rodar `make loadfiles`

## Documentação da API
