FROM golang:latest

WORKDIR /usr/src/app
ENV PORT 8000
ENV HOST 0.0.0.0

COPY . .

RUN go build -o bin ./cmd/app

ENTRYPOINT [ "./bin" ]