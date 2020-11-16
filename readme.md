# efishery-task
Built with Golang, PostgreSQL, Redis, Echo, Bcrypt, and other third-party libraries.
## Installation
Make sure Golang and is already installed.
Clone the repo to the directory ```go/src/github.com/ilhamrobyana/bts-test```
Install the dependendices with ```dep ensure```
## Peresquite
Before running the application, set the dot env file first using .env.example as the layout.

PostgreSQL and Redis needed to be setup first.

And then migrate the data structure to PostgreSQL

```cd auth-api/migration```

```go run main.go```

## Running
Go to the root folder of the project and enter the following:
```go run auth.go```

## Documentation
The API documentation can be found in the following link:
```https://documenter.getpostman.com/view/8102951/TVep9Tn7```