FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy && go mod download

COPY . .

RUN go build -o webserver cmd/webserver/main.go

FROM alpine

COPY --from=build /app/webserver /webserver

EXPOSE 8080

CMD [ "./webserver" ]