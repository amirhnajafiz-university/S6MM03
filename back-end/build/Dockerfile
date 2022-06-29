FROM golang:1.18-alpine

WORKDIR /app

COPY * ./

RUN go build -o /docker-gs-ping

RUN go run db.go

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
