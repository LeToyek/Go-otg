include .env
build:
	go build -v -o go-otg cmd/*.go

run:
	make build
	./go-otg

docker-run:
	docker build -t go-otg .
	docker run -p 5000:5000 --name=go-otg go-otg

docker-stop:
	docker stop go-otg

migrateup:
	migrate -path db/migrations -database "postgresql://root:handoko@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:handoko@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

postgres:
	docker run --name db-gootg-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=handoko -d postgres

createdb:
	docker exec -it db-gootg-postgres createdb --username=root --owner=root ${DB_NAME}

dropdb:
	docker exec -it db-gootg-postgres dropdb ${DB_NAME}

up:
	docker compose up --build

down:
	docker compose down --remove-orphans

pgbash:
	docker exec -it db-gootg-postgres /bin/sh


localmigrateup:
	migrate -path db/migrations -database "postgresql://postgres:handoko@localhost:5432/go-otg?sslmode=disable" -verbose up

localmigratedown:
	migrate -path db/migrations -database "postgresql://postgres:handoko@localhost:5432/go-otg?sslmode=disable" -verbose down