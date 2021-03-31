init-project:
	go mod init bookstore_oauth-api

test:
	go test -race ./... -v

run:
	go run -race src/main.go 

