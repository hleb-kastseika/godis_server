[![Build Status](https://travis-ci.org/gleb-kosteiko/godis_server.svg?branch=master)](https://travis-ci.org/gleb-kosteiko/godis_server)

# Godis Server

Godis (like Go + Redis) - implementation of simple Redis-like cache for Go language training. This is the server side of Godis.

## Requirements:
- Golang 1.11.2 or higher
- Docker (for running in container)

## How to build
```
go clean && go test ./... -coverprofile coverage.out && go build
```

## How to run
#### Run in terminal
```
./godis_server
```

#### Run options:
  - -m, -mode -- storage mode of the server: "memory" or "disk" (default "memory")
  - -p, -port -- port on which the server runs (default "9090")
<!---  - -v, -verbose - turn on/off full log of client requests, possible values: "true" and "false" (default "false") -->
Example: ``` ./godis_server -p=8080 -m=disk```

#### Run in Docker container
```
docker build -t godis_server .
docker run --rm -p 9090:9090 godis_server
```

For access the server via HTTP use any HTTP client, the main resource - http://localhost:9090/storage
