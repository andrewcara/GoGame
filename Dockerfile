FROM golang:1.23.5

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

EXPOSE 8080

COPY . .

WORKDIR /app/srv

CMD ["go", "run", "."]