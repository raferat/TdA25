FROM golang:1.23.2-alpine3.20 AS build-stage

WORKDIR /app
COPY ./backend .
RUN go build -o /app/tdaserver

FROM alpine:latest AS run-stage

COPY --from=build-stage /app/tdaserver .

RUN chmod +x tdaserver

CMD ["./tdaserver"]
