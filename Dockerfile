FROM golang:1.17.1-alpine as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/bin

FROM scratch

WORKDIR /

COPY --from=build /app/bin /app

ENTRYPOINT ["/app"]