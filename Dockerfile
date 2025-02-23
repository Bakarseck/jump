FROM golang:latest

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o jump -ldflags '-extldflags "-static"'

FROM scratch

COPY --from=0 /app/jump /jump

CMD ["./jump"]