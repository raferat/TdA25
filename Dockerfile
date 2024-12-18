FROM golang:1.23.2-alpine3.20 AS build-backend

WORKDIR /app
COPY ./backend .
RUN ["apk", "add", "gcc", "musl-dev", "binutils"]
RUN ["go", "build", "-o", "/app/tdaserver"]

######################## FRONTEND #####################################
FROM node:23.4-alpine3.21 AS build-frontend

WORKDIR /app
COPY ./frontend .
RUN ["npm", "install"]
RUN ["npm", "run", "build"]

######################## FINAL IMAGE ##################################
FROM alpine AS run-stage

WORKDIR /app
COPY --from=build-backend /app/tdaserver .
RUN ["chmod", "+x", "tdaserver"]
ENV PORT=":4242"
EXPOSE 4242

COPY --from=build-frontend /app/build ./website

ENTRYPOINT ["/app/tdaserver"]
