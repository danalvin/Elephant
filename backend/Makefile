APP = elephant 

fmt:
	go fmt ./...

run: 
	go run main.go s

build: fmt
	go mod tidy
	go build -o ${APP} .

test:
	go test tests/* -v

pay:
	go run main.go p

# docker-build:

# docker-run:
