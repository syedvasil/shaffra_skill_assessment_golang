## Assignment 

## Code Coverage
**main**: &nbsp; [![coverage report](https://github.com/syedvasil/concurrent/blob/main/coverage.svg)](https://github.com/syedvasil/shaffra_skill_assessment_golang/blob/main/coverage/coverage.out)

## 🔧 Installation

After cloning the repository, navigate into the directory and make sure you are in **main** branch.

- Make sure you have **go and mongoDb** installed
- Run `go mod tidy`, to download dependencies.

- Run `go run ./cmd/server` to run this code on your local.

Once your server is running you can browse the swagger docs

@http://localhost:8080/docs/index.html#/


- Run `go tool cover -html="coverage.out"` to view the report on web browser.
- Run `go test -coverprofile=coverage.out ./...` to create coverage report.


## Important Points regarding the requirements

1. Efficient use of GoLang constructs (goroutines, channels, error handling). ->  Unable to Demonstrate that as part of the CURD operations, but check out this repo of mine which showcases the Efficient use https://github.com/syedvasil/concurrent/tree/main 
2. no need make requests use go routines gin server does it out of the box.