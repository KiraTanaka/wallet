FROM golang:1.23.1-bullseye AS build

WORKDIR /wallet

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN ls

RUN go build ./cmd/main.go

EXPOSE 8080

RUN chmod +x ./entypoint.sh

CMD ["./main"]

