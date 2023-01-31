FROM golang:1.19-bullseye

WORKDIR /app

COPY * ./

RUN go build -o /runner main.go

RUN go run db.go

EXPOSE 8080

CMD [ "/runner" ]
