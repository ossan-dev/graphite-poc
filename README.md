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

1. `docker run  --name graphite  --restart=always  -p 80:80  -p 2003-2004:2003-2004  -p 2023-2024:2023-2024  -p 8125:8125/udp  -p 8126:8126  graphiteapp/graphite-statsd`

To query metrics:
Enconding:

- "'" => %27
- "," => %2C

1. `curl "http://0.0.0.0:80/render?target=consolidateBy(webserver.get_todos.success%2C%27sum%27)&from=-1h&format=json"`
1. `curl "http://0.0.0.0:80/render?target=removeEmptySeries(webserver.get_todos.success%2C0.1)&from=-1h&format=json"`
1. `curl "http://0.0.0.0:80/render?target=consolidateBy(webserver.get_todos.success%2C%27sum%27)&from=-10min&format=json"`
1. `curl "http://0.0.0.0:80/render?target=maxSeries(webserver.get_todos.success)&from=-10min&format=json"`
1. `curl "http://0.0.0.0:80/render?target=removeBelowValue(webserver.get_todos.success%2C1)&from=-15min&format=json"`
1. `curl "http://0.0.0.0:80/render?target=removeBelowValue(transformNull(webserver.get_todos.success)%2C0)&from=-60min&format=json"`
1. `curl "http://0.0.0.0:80/render?target=removeBelowValue(transformNull(webserver.get_todos.success)%2C%200)&format=json"`

## docker-compose

To build it:

1. `docker-compose build --no-cache`

To run it:

1. `docker-compose up -d`

To stop it:

1. `docker-compose down --remove-orphans`

## Integration Tests

To run it:

1. `go test -v -tags=integration ./...`
1. `go test ./tests -tags=integration`

## Run containers after conflicts

### Repro the issue

1. `docker-compose up -d`
2. `docker ps` => you should see the two containers up & running
3. `go test ./tests -tags=integration`

    ```text
    --- FAIL: TestGetTodos (0.86s)
        get_todos_test.go:17: 
                    Error Trace:    $HOME/graphite-poc/tests/container.go:24
                                                            $HOME/graphite-poc/tests/get_todos_test.go:17
                    Error:          Received unexpected error:
                                    compose up: Error response from daemon: Conflict. The container name "/graphite" is already in use by container "ceb36087cd8f58a82bdc07d755c36b7e1d043b5e4a670a37cf44c1947f6c7826". You have to remove (or rename) that container to be able to reuse that name.
                    Test:           TestGetTodos
    FAIL
    FAIL    github.com/ossan-dev/graphitepoc/tests  7.805s
    FAIL
    ```

### Fix the issue

The fix for the code is contained in the commit with the message `fixConflictContainers`. We try to spin up the containers again in case of an issue.
