#!make
include .env
export $(shell sed 's/=.*//' .env)

setup:
	go mod download

test:
	go test -cover ./internal/...

lint:
	golangci-lint run ./...

updb:
	POSTGRES_USER=${POSTGRES_USER} POSTGRES_PASSWORD=${POSTGRES_PASSWORD} docker-compose up -d postgres && echo For logs, run docker-compose logs -f postgres

stopdb:
	docker-compose stop postgres

loadfiles:
	go run ./scripts/populate_db/script.go > log/populate_db.log

run: updb
	go run cmd/api/main.go

upapi:
	docker-compose up -d api && docker-compose logs -f api

createDefault:
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
		\"neighborhood\": \"JARDIM\", \
		\"addr_extra_info\": \"Loren ipsum\" \
	}" -H 'Content-Type: application/json' http://localhost:8000/street_market

listCreatedByDistrict:
	make listByDistrict district=VILA%20FORMOSA page=$(page)

listCreatedByRegion5:
	make listByRegion5 region5=Leste page=$(page)

listCreatedByName:
	make listByName name=RAPOSO%20TAVARES page=$(page)

listCreatedByNeighborhood:
	make listByNeighborhood neighborhood=JARDIM%20SARAH page=$(page)

listCreatedFullFilter:
	curl -v -H 'Content-Type: application/json' 'http://localhost:8000/street_market?district=VILA%20FORMOSA&region5=Leste&name=RAPOSO%20TAVARES&neighborhood=JARDIM%20SARAH&page=$(page)'

listByDistrict:
	curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?district=$(district)&page=$(page)

listByRegion5:
	curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?region5=$(region5)&page=$(page)

listByName:
	curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?name=$(name)&page=$(page)

listByNeighborhood:
	curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?neighborhood=$(neighborhood)&page=$(page)

create:
	curl -v -d '${body}' -H 'Content-Type: application/json' http://localhost:8000/street_market

edit:
	curl -X 'PATCH' -v -d '${body}' -H 'Content-Type: application/json' http://localhost:8000/street_market/${id}

delete:
	curl -X 'DELETE' -v http://localhost:8000/street_market/${id}

list:
	curl -v -H 'Content-Type: application/json' http://localhost:8000/street_market?page=${page}
