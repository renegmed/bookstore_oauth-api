init-project:
	go mod init bookstore_oauth-api
up:
	docker-compose up -d 

down:
	docker-compose down 
	
test:
	go test -race ./... -v -cover

build:
	go build -race -o oauth-api src/main.go 

run: build
	./oauth-api -server localhost:8081 

go-run:
	go run -race src/main.go --server localhost:8081 

get:
	curl localhost:8081/oauth/access_token/123123123

post-invalid:
	curl -XPOST localhost:8081/oauth/access_token -d '{"access_token": "abc123", "user_id": 1, "client_id": 1}'	


add-user:
	curl -XPOST localhost:8081/oauth/access_token -d '{"id":1, "first_name": "John", "last_name": "Doe", "email": "johndoe@test.com"}'  #', "date_created": "2021-03-21"}'