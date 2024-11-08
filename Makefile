_build:
	go build -o dft main.go
	chmod +x dft

lint:
	clear
	golanci-lint run

test:
	go test -v ./...