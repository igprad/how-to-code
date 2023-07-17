## About
This is an example project for learning how to create rest service using go (trying to not use framework).\
The idea is to create User management service with CRUD functionalities as the MVP.\

## Project Structure
The structure of this project is following: https://github.com/qiangxue/go-rest-api \
So it became like this.
```
.
├── cmd                  main applications of the project
│   └── server           the API server application
├── config               configuration files for different environments
├── internal             private application and library code
│   ├── config           configuration library
│   ├── entity           entity definitions and domain logic
│   ├── errors           error types and handling
│   ├── healthcheck      healthcheck feature
│   ├── test             helpers for testing purpose
│   └── user             user domain layer features
├── migrations           database migrations
├── pkg                  public library code
└── testdata             test data scripts
```

## How-to
### Pre-requisite
Because still no data seed for this example, need to setup local postgres on your machine.\
Also need setup database: `training` and table: `user` for default.
### Run
Go to root then `go run cmd/server/main.go`\
Now access on your local: http://localhost:6969/users for the example.

