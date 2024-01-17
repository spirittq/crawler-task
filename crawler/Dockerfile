FROM golang:1.21-alpine as build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

COPY crawler/ ./crawler
COPY shared/ ./shared

WORKDIR /src/crawler

RUN go build -o crawler

FROM alpine:3.19

ARG SCRAPE_URL
ENV SCRAPE_URL=${SCRAPE_URL}
ARG ALLOWED_DOMAIN
ENV ALLOWED_DOMAIN=${ALLOWED_DOMAIN}
ARG SERVER_PORT
ENV SERVER_PORT=${SERVER_PORT}
ARG ASYNC_COUNT
ENV ASYNC_COUNT=${ASYNC_COUNT}
ARG SERVER_DOMAIN
ENV SERVER_DOMAIN=${SERVER_DOMAIN}

RUN apk update && apk add --no-cache ca-certificates curl

WORKDIR /app

COPY --from=build /src/crawler/crawler .

CMD [ "./crawler" ]
