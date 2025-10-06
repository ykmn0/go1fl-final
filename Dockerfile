FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o scheduler main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

COPY --from=builder /app/scheduler .
COPY --from=builder /app/web ./web

ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db

EXPOSE 7540

RUN mkdir -p /data

CMD ["./scheduler"]
