FROM golang:1.23.2-alpine3.20 AS build-stage

WORKDIR /app
COPY ./backend .
RUN ["go", "build", "-o", "/app/tdaserver"]

FROM alpine AS run-stage

WORKDIR /app
COPY --from=build-stage /app/tdaserver .
RUN ["chmod", "+x", "tdaserver"]
ENV PORT=":80"
EXPOSE 80

ENTRYPOINT ["/app/tdaserver"]
