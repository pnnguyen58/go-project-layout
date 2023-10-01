# go-project-layout

## Description


## Prerequisites


## Technologies and Frameworks


## Directory structure
    📁 go-project-layout
    |__ 📁 api
    |__ 📁 build
        |__ 📁 ci
        |__ 📁 packages
    |__ 📁 cmd
        |__ 📁 app
        |__ 📁 workers
    |__ 📁 configs
    |__ 📁 deploy 
        |__ 📁 docker-compose
        |__ 📁 kubernetes
        |__ 📁 terraform
    |__ 📁 docs
    |__ 📁 internal 
        |__ 📁 activities // temporal activities
        |__ 📁 app // application layer
        |__ 📁 models // entitiy layer
        |__ 📁 repositories // database repositories
        |__ 📁 workflows // temporal workflows
        |__ 📁 defined // types and constants
    |__ 📁 pkg
        |__ 📁 clients
        |__ 📁 google
        |__ 📁 logger
        |__ 📁 persistence
        |__ 📁 proto
        |__ 📁 proto_generated
    |__ 📁 scripts
    |__ 📁 tests
    |__ Makefile
    |__ go.mod
    |__ go.sum
    |__ README.md

## How to run local
- Run services and workers: `docker compose up --build -d`. 
If `aspire-api` and `aspire-worker` not running, maybe initialize not done yet, please 
re-run the docker compose command again.
- Generate proto: `make`
- TODO: run services and worker separately for scaling

## How to test


## Monitoring
- Workflow status: http://localhost:8080

## Conceptual model


## TODO tasks
- Add a policy check to make sure that the customers can view them own loan only.
- Implement more unit tests
- Implement exceed amount repayment
- Fix docker network and wait dependencies healthy
- Implement error abstract


## References
- Project layout: https://github.com/golang-standards/project-layout
- Transcoding of HTTP/JSON to gRPC: https://adevait.com/go/transcoding-of-http-json-to-grpc-using-go
- Workflow framework: https://docs.temporal.io/
- Clean architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

