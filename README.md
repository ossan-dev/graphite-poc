# Graphite POC

## webserver

To run it:

1. `cd ./cmd/webserver`
2. `go run .`

To test it:

1. `curl http://localhost:8080/todos` => working
2. `curl http://localhost:8080/todo?id=1` => working
3. `curl http://localhost:8080/todo?id=11` => not found

To build it:

1. `docker build -t webserver .`

To run a container:

1. `docker run -p 8080:8080 webserver`

## graphite

To run it:

1. `docker run -d  --name graphite  --restart=always  -p 80:80  -p 2003-2004:2003-2004  -p 2023-2024:2023-2024  -p 8125:8125/udp  -p 8126:8126  graphiteapp/graphite-statsd`

To query metrics:

1. `curl "http://0.0.0.0:80/render?target=consolidateBy(webserver.get_todo_by_id.errors.not_found%2C%27sum%27)&from=-1h&format=json"`

## docker-compose

To build it:

1. `docker-compose build --no-cache`

To run it:

1. `docker-compose run --service-ports webserver`
