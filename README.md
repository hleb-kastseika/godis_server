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

#### REST API documentation
For access the server via HTTP use any HTTP client, There is the list of avalaible resources and methods:

---
##### Get value by key
`GET`  http://localhost:9090/storage?key=key_value

_Headers_: 'Content-Type:application/json'

_URL Params_: key (string, required)

_Success Response_: Code - 200 Ok, Content - {"key":"test","value":"test value"}

_Error Response_: Code - 404 Not Fount

---
##### Get all values
`GET`  http://localhost:9090/storage

_Headers_: 'Content-Type:application/json'

_Success Response_: Code - 200 Ok, Content - [{"key":"","value":""},{"key":"","value":""},{"key":"test","value":"test value"},{"key":"test2","value":"test value 2"}]

---
##### Set value
`POST`  http://localhost:9090/storage

_Headers_: 'Content-Type:application/json'

_Body_: '{"key":"test","value":"test value"}'

_Success Response_: Code - 200 Ok, Content - {"key":"test","value":"test value"}

_Error Response_: Code - 400 Bad Request, Content - 'Pass the data in next format: {"key":"","value":""}'

---
##### Delete value
`DELETE` http://localhost:9090/storage?key=key_value

_Headers_: 'Content-Type:application/json'

_URL Params_: key (string, required)

_Success Response_: Code - 200 Ok

---
##### Find keys that matches to expression

`GET` http://localhost:9090/storage/keys?match=value*

_Headers_: 'Content-Type:application/json'

_URL Params_: match (string, required, should contain '*')

_Success Response_: Code - 200 Ok, Content - ["test1","test2","test3"]