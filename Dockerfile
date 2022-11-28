FROM golang:1.19.3-bullseye

WORKDIR /app

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/web/*.go

CMD ["app"]