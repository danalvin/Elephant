APP = elephant 

run: 
	go run main.go s

build:  
	go fmt ./...
	go mod tidy
	go build -o ${APP} .

test:
	go test tests/* -v

# docker-build:

# docker-run:
