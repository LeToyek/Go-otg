build:
	go build -v -o go-otg cmd/*.go

run:
	make build
	./go-otg

docker-run:
	docker build -t go-otg .
	docker run -d -p 5000:5000 --name=go-otg go-otg

docker-stop:
	docker stop go-otg

migrateup:
	migrate -path db/migrations -database "postgresql://root:handoko@localhost:5432/db-gootg?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:handoko@localhost:5432/db-gootg?sslmode=disable" -verbose down

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=handoko -d postgres

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root db-gootg

dropdb:
	docker exec -it postgres15 dropdb db-gootg
