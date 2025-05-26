FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download
RUN go mod verify

COPY . .

RUN go build -o webserver cmd/webserver/main.go

FROM alpine

COPY --from=build /app/webserver /webserver

EXPOSE 8080

CMD [ "./webserver" ]