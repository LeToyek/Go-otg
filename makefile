build:
	go build -v -o go-otg cmd/*.go

run:
	make build
	./go-otg

docker-run:
	docker build -t go-otg .
	docker run -d -p 5000:5000 go-otg