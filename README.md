# notable-take-home

This project contains a simple webserver for notable take home project. The Go server uses MongoDB to store data.

## Features

- [x] Go API using [Gorilla mux](github.com/gorilla/mux)

- [x] [MongoDB](github.com/mongodb/mongo-go-driver) integration
- [x] [JWT](github.com/dgrijalva/jwt-go) authentication
- [x] Go server configuration via [.yml](gopkg.in/yaml.v2) file
- [x] Using Go Modules for dependency management
- [x] Makefile based project

## Building and running

The **notable-take-home** contains a Go server which serves an RESTful API which handles all request. To run the Go server run the following commands:

```sh
git clone https://github.com/micahli/notable-take-home.git

# Run the Go server on localhost:8081
cd notable-take-home

make build
./bin/starter
```

## Usage and configuration

Befor you start the **notable-take-home** edit the `config.yml` file:

```yaml
# Port where the notable-take-home server should listen on
listen_address: :8081

api:
  # Domain is for the frontend
  domain: http://localhost:8081
  # Signing secret for the jwt authentication
  signing_secret: "your-super-secret"

database:
  # Connection URL for MongoDB and the name of the MongoDB database
  mongodb:
    connection_uri: mongodb://localhost:27017
    database_name: notable-test


```
Usage of starter:
  -config.file string
      Path to the configuration file. (default "config.yml")
  -debug
      Show debug information.
```

## Dependencies

The Go dependencies are:

- [crypto - Go supplementary cryptography libraries](golang.org/x/crypto)
- [gorilla/mux - A powerful URL router and dispatcher for golang](github.com/gorilla/mux)
- [jwt-go - Golang implementation of JSON Web Tokens (JWT)](github.com/dgrijalva/jwt-go)
- [logrus - Structured, pluggable logging for Go](github.com/sirupsen/logrus)
- [mongo-go-driver - The Go driver for MongoDB](github.com/mongodb/mongo-go-driver)
- [yaml.v2 - YAML support for the Go language](gopkg.in/yaml.v2)
