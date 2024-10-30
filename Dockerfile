FROM docker.io/library/golang:1.23.2-alpine3.20 AS build-stage

COPY . .
RUN go build

FROM alpine:latest AS run-stage

COPY --from=build-stage tdaserver .

RUN chmod +x tdaserver

CMD ["./tdaserver"]
