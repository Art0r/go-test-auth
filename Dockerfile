FROM golang:1.25-rc-bookworm AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags='-w -s' cwd/api/main.go

FROM alpine:3.22.1

WORKDIR /app

COPY --from=builder /app/main /app/main

ENV PORT=3000
ENV GIN_MODE=release

EXPOSE $PORT

RUN chmod +x /app/main
CMD [ "/app/main" ]

# docker build -t teste-go  .
# docker run -d -p 3000:3000 -e PORT=3000 -e SECRET=foo -e GIN_MODE=release teste-go:latest

