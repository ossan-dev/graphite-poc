# Graphite POC

## webserver

To run it:

1. `cd ./cmd/webserver`
2. `go run .`

To test it:

1. `curl <http://localhost:8080/todos>` => working
2. `curl <http://localhost:8080/todo?id=1>` => working
3. `curl <http://localhost:8080/todo?id=11>` => not found

To build it:

1. `docker build -t webserver .`

To run a container:

1. `docker run -p 8080:8080 webserver`
