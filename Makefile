_build:
	go build -o dft main.go
	chmod +x dft
	sudo mv dft /usr/local/bin

lint:
	clear
	golanci-lint run

test:
	go test -v ./...