# Unico API Test

Essa API tem como objetivo administrar os eventos de feira de uma determinada cidade. Nela é possível criar, editar, excluir e listar feiras.

## Requisitos
- Golang 1.18+
- Make

## Setup do projeto
Rodar `make setup`. Esse comando vai baixar e instalar as dependências.

## Cobertura de testes
Rodar `make test`. Esse comando vai rodar o `go test` para todos os arquivos dentro de internal e dar a cobertura de cada diretório.

## Rodando a API
### Via build
Rodar `make run`. Vai rodar o `go run` no main que sobe a API e subir o container do banco de dados. Essa forma permite que alterações sejam testadas sem necessidade de um novo build de container. **O `.env` deve ser ajustado para a variável correta de `host`**
### Via docker-compose
Rodar `make upapi`. Vai rodar o `docker-compose` da API e deixar observando os logs.

## Populando base para testes
É possível inserir uma base de dados inicias com dados de feiras da cidade de São Paulo, que são fornecidas via CSV no site da mesma. A massa de dados porém, apresenta alguns problemas de encoding além de terem sofrido alguns ajustes.

Rodar `make loadfiles`

## Testando a API
O Makefile do projeto tem diversos exemplos de requisições que pode ser feitas para a API. **É importante que a etapa de [rodando a api](#rodando-a-api) tenha sido feita**
___


## Documentação da API
Essa API não conta com autenticação para ser utilizada.

Todas as rodas tem [comandos make](#testando-a-api) que podem ser utilizados para testar rapidamente o comportamento da rota.

- Feira
  - [Criação](#criação)
  - [Edição](#edição)
  - [Exclusão](#exclusão)
  - [Listar](#listar)

### Criação
|  	|  	|
|---	|---	|
| **Método** 	| Post 	|
| **Caminho** 	| /street_market 	|
| **Cabeçalho** 	| `Content-Type: application/json` 	|

**Corpo**

Json com o seguinte schema:

| chave  	| tipo  	| descrição  	|
|---	|---	|---	|
| long  	| float  	| Longitude da localização do estabelecimento no território do Município, conforme MDC  	|
| lat  	| float  	| Latitude da localização do estabelecimento no território do Município, conforme MDC  	|
| sect_cens  	| string  	| Setor censitário conforme IBGE  	|
| area  	| string  	| Área de ponderação (agrupamento de setores censitários) conforme IBGE 2010  	|
| id_dist  	| string  	| Código do Distrito Municipal conforme IBGE  	|
| district  	| string  	| Nome do Distrito Municipal  	|
| id_sub_th  	| string  	| Código de cada uma das 31 Subprefeituras (2003 a 2012  	|
| subtownhall  	| string  	| Nome da Subprefeitura (31 de 2003 até 2012  	|
| region_5  	| string  	| Região conforme divisão do Município em 5 áreas  	|
| region_8  	| string  	| Região conforme divisão do Município em 8 áreas  	|
| name  	| string  	| Denominação da feira livre atribuída pela Supervisão de Abastecimento  	|
| register  	| string  	| Número do registro da feira livre na PMSP  	|
| street  	| string  	| Nome do logradouro onde se localiza a feira livre  	|
| number  	| string  	| Um número do logradouro onde se localiza a feira livre  	|
| neighborhood  	| string  	| Bairro de localização da feira livre  	|
| addr_extra_info  	| string  	| Ponto de referência da localização da feira livre  	|


**Resposta**

**[Resposta de erro](#resposta-de-erro)**

**Resposta de sucesso**

Sem conteúdo

**Headers**
- `Location`

#### Exemplo de edição
```bash
  curl -v -d "{ \
		\"long\": -46548146, \
		\"lat\": -23568390, \
		\"sect_cens\": \"355030885000019\", \
		\"area\": \"3550308005040\", \
		\"id_dist\": \"87\", \
		\"district\": \"VILA FORMOSA\", \
		\"id_sub_th\": \"26\", \
		\"subtownhall\": \"ARICANDUVA\", \
		\"region_5\": \"Leste\", \
		\"region_8\": \"Leste 1\", \
		\"name\": \"RAPOSO TAVARES\", \
		\"register\": \"1129-0\", \
		\"street\": \"Rua dos Bobos\", \
		\"number\": \"500\", \
		\"neighborhood\": \"JARDIM SARAH\", \
		\"addr_extra_info\": \"Loren ipsum\" \
	}" -H 'Content-Type: application/json' http://localhost:8000/street_market
```

#### Exemplo de resposta
```json
{
  "error": "Input is invalid. Neighborhood is required"
}
```

#### Teste via make
`make createDefault`.
___
### Edição
|  	|  	|
|---	|---	|
| **Método** 	| Patch 	|
| **Caminho** 	| /street_market/{ID} 	|
| **Cabeçalho** 	| `Content-Type: application/json` 	|

**Parâmetros do caminho**
|chave   	| descrição   	|
|---	|---	|
| ID  	| ID do recurso que quer ser atualizado  	|

**Corpo**

Json com o seguinte schema:

| chave  	| tipo  	| descrição  	|
|---	|---	|---	|
| long  	| float  	| Longitude da localização do estabelecimento no território do Município, conforme MDC  	|
| lat  	| float  	| Latitude da localização do estabelecimento no território do Município, conforme MDC  	|
| sect_cens  	| string  	| Setor censitário conforme IBGE  	|
| area  	| string  	| Área de ponderação (agrupamento de setores censitários) conforme IBGE 2010  	|
| id_dist  	| string  	| Código do Distrito Municipal conforme IBGE  	|
| district  	| string  	| Nome do Distrito Municipal  	|
| id_sub_th  	| string  	| Código de cada uma das 31 Subprefeituras (2003 a 2012  	|
| subtownhall  	| string  	| Nome da Subprefeitura (31 de 2003 até 2012  	|
| region_5  	| string  	| Região conforme divisão do Município em 5 áreas  	|
| region_8  	| string  	| Região conforme divisão do Município em 8 áreas  	|
| name  	| string  	| Denominação da feira livre atribuída pela Supervisão de Abastecimento  	|
| register  	| string  	| Número do registro da feira livre na PMSP  	|
| street  	| string  	| Nome do logradouro onde se localiza a feira livre  	|
| number  	| string  	| Um número do logradouro onde se localiza a feira livre  	|
| neighborhood  	| string  	| Bairro de localização da feira livre  	|
| addr_extra_info  	| string  	| Ponto de referência da localização da feira livre  	|

**Resposta**

**[Resposta de erro](#resposta-de-erro)**

**Resposta de sucesso**

Sem conteúdo


#### Exemplo de edição
```bash
  curl -X 'PATCH' -v -d '{\"number\": \"999\"}' -H 'Content-Type: application/json' http://localhost:8000/street_market/{id}
```

#### Exemplo de resposta
```json
{
  "error": "Input is invalid. ID is an invalid UUID."
}
```

#### Teste via make
`make edit body="{\"number\": \"999\"}" id=` complete com um ID, que pode ser obtido através da [lista](#teste-via-make-listagem).
____
### Exclusão
|  	|  	|
|---	|---	|
| **Método** 	| Delete 	|
| **Caminho** 	| /street_market/{ID} 	|
| **Cabeçalho** 	| `Content-Type: application/json` 	|

**Parâmetros do caminho**
|chave   	| descrição   	|
|---	|---	|
| ID  	| ID do recurso que quer ser atualizado  	|

**Resposta**

**[Resposta de erro](#resposta-de-erro)**

**Resposta de sucesso**

Sem conteúdo

#### Exemplo de exclusão
```bash
  curl -X 'DELETE' -v http://localhost:8000/street_market/{ID}
```

#### Exemplo de resposta
```json
{
  "error": "Input is invalid. ID is an invalid UUID."
}
```

#### Teste via make
`make delete id=` complete com um ID, que pode ser obtido através da [lista](#teste-via-make-listagem)
____
### Listar
Rota da listar feiras. As ferias serão ordenadas de forma decrescente considerando sua data de criação.

Essa rota é paginada, e utiliza o parâmetro page identificar qual a pagina está sendo solicitada. A paginação termina quando não são retornados mais dados para uma determinada pagina.

|  	|  	|
|---	|---	|
| **Método** 	| Delete 	|
| **Caminho** 	| /street_market 	|
| **Cabeçalho** 	| `Content-Type: application/json` 	|

**Parâmetros de query**
| nome  	| descrição  	|
|---		|---	|
| name  	| Nome da feira  	|
| district  	| Distrito da feita  	|
| region5  	| Região conforme divisão do Município em 5 áreas da feira  	|
| neighborhood  	| Bairro da feira  	|
| page  	| pagina a ser buscada  	|

**Resposta**

**[Resposta de erro](#resposta-de-erro)**

**Resposta de sucesso**
| nome  	| tipo  	| descrição  	|
|---	|---	|---	|
| data   	| lista de [feira](#feira)  	|   	|
#### Feira
| chave  	| tipo  	| descrição  	|
|---	|---	|---	|
| id  	| string (uuid)  	| UUID que identifica aquele recurso na API  	|
| long  	| float  	| Longitude da localização do estabelecimento no território do Município, conforme MDC  	|
| lat  	| float  	| Latitude da localização do estabelecimento no território do Município, conforme MDC  	|
| sect_cens  	| string  	| Setor censitário conforme IBGE  	|
| area  	| string  	| Área de ponderação (agrupamento de setores censitários) conforme IBGE 2010  	|
| id_dist  	| string  	| Código do Distrito Municipal conforme IBGE  	|
| district  	| string  	| Nome do Distrito Municipal  	|
| id_sub_th  	| string  	| Código de cada uma das 31 Subprefeituras (2003 a 2012  	|
| subtownhall  	| string  	| Nome da Subprefeitura (31 de 2003 até 2012  	|
| region_5  	| string  	| Região conforme divisão do Município em 5 áreas  	|
| region_8  	| string  	| Região conforme divisão do Município em 8 áreas  	|
| name  	| string  	| Denominação da feira livre atribuída pela Supervisão de Abastecimento  	|
| register  	| string  	| Número do registro da feira livre na PMSP  	|
|  addr_extra_info  	| string  	| Ponto de referência da localização da feira livre  	|

#### Exemplo de consulta
```bash
  curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?page=1&region5=Leste
```

#### Exemplo de resposta
```json
{
  "data":[{
    "id":"dc9c826f-9a05-41d1-a8ec-9cc37e9fac66",
    "long":-46548146,
    "lat":-23568390,
    "sect_cens":"355030885000019",
    "area":"3550308005040",
    "id_dist":"87",
    "district":"VILA FORMOSA",
    "id_sub_th":"26",
    "subtownhall":"ARICANDUVA",
    "region_5":"Leste",
    "region_8":"Leste 1",
    "name":"RAPOSO TAVARES",
    "register":"1129-0",
    "street":"Rua dos Bobos",
    "number":"999",
    "neighborhood":"JARDIM SARAH",
    "addr_extra_info":"Loren ipsum"
  }]
}
```

#### Teste via make listagem
`make list page=` complete com a pagina desejada, ou deixe em brando para pagina 1.
___
### Resposta de erro

Json com o seguinte esquema:
| chave  	| tipo  	| descrição   	|
|---	|---	|---	|
|  error 	| string  	| Mensagem de erro  	|



## Todo
- Test dos middlewares
- Controle de level no Logger
- Opção configurável de pretty no log.
- Adicionar ID como chave no contexto e adicionar no log.
