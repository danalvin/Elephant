## Elephant Microservice
- Backend Service for scheduling, calculating and depositing payments. 

## Description
Using preconfigured Payment Gateway of Choice (MPesa/VISA/MasterCard), This service will
automate payments of active workers on weekly basis. The design should facilitate user to
add additional constraints.



```go
    go func(){
        log.Println("Where's my money?")
    }()
``` 
## Core Features
- [ ] Running cron jobs
- [ ] Payment Infrastructure Integration (MPesa)
- [ ] Download reports 
- [ ] Unit & Integration testing

### Local Set Up  
+ Clone the [repo](https://github.com/vonmutinda/Elephant.git) 
+ Create `config.toml` using `config.toml.example` file. 
+ RUN - `make run` or `go run main.go serve`

### API Routes 
We won't be needed API endpoints much.
1. FOO             - `http://localhost:6060/api/v1/foo`  

## Technology Stack 
A list of technologies used in this project:
- [Golang version `go1.13.12`](https://golang.org) 
- [Gin](https://github.com/gin-gonic/gin) Gin Framework is a cool lightweight alternative for GorillaMux.
- [Viper](github.com/spf13/viper) Reading environment variables.
- [Cobra](github.com/spf13/cobra) Commandline tooling
- [Gorm.io](github.com/jinzhu/gorm) Golang's ORM
