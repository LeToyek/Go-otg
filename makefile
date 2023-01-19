build:
	go build -v -o go-otg cmd/*.go

run:
	make build
	./go-otg