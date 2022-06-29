FROM golang:1.18-bullseye

WORKDIR /app

COPY * ./

RUN go build -o /docker-gs-ping main.go

RUN go run db.go

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
