FROM golang:1.22.2 AS BUILDER

WORKDIR /app
COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY go.sum go.sum
COPY main.go main.go
COPY .env .env
COPY docs docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
  GOOS=linux go build -o meuprimeirocrudgo .

FROM golang:1.22-alpine AS RUNNER

COPY --from=BUILDER /app/meuprimeirocrudgo .
COPY --from=BUILDER /app/.env .env

EXPOSE 8080

CMD [ "./meuprimeirocrudgo" ]