# notable-take-home

This project contains a simple webserver for notable take home project. The Go server store data in memory.

To run the server, you should install go runtime on the local machine.

## APIs: 
1. get all the doctors
http://localhost:8081/api/v1/doctor
2. get one doctor's appointments
http://localhost:8081/api/v1/doctor/001
3. add an appointment to a doctor
POST request:

http://localhost:8081/api/v1/doctor/appointment

{
  "doctoruid": "001",
  "patientfirstname": "micassh",
  "patientlastname": "zhdddo",
  "datetime": "2020-01-20 15:15",
  "kind": 0
}

4. cancel an appointment

DELETE request:
http://localhost:8081/api/v1/doctor/appointment

{
  "uid": "7a161f5c64f548e1985a7897e00e346a",
}


## Features

- [x] Go API using [Gorilla mux](github.com/gorilla/mux)
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
- [logrus - Structured, pluggable logging for Go](github.com/sirupsen/logrus)
- [yaml.v2 - YAML support for the Go language](gopkg.in/yaml.v2)
