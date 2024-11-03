FROM golang:1.23.2-alpine3.20 AS build-stage

WORKDIR /app
COPY ./backend .
RUN ["go", "build", "-o", "/app/tdaserver"]

FROM postgres:alpine AS run-stage

ENV POSTGRES_PASSWORD="password123"


#setup postgresql tables and logic
WORKDIR /docker-entrypoint-initdb.d
COPY ./postgresql.conf.d .
RUN ["chmod", "+x", "init.sh"]


WORKDIR /app
COPY --from=build-stage /app/tdaserver .
RUN ["chmod", "+x", "tdaserver"]
COPY tda25-entrypoint.sh .
RUN ["chmod", "+x", "tda25-entrypoint.sh"]
ENV PORT=":80"
EXPOSE 80

ENTRYPOINT ["/app/tda25-entrypoint.sh"]
