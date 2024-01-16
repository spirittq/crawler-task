FROM golang:1.21-alpine as build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o crawler-task

FROM alpine:3.19

ARG SCRAPE_URL
ENV SCRAPE_URL=${SCRAPE_URL}
ARG ALLOWED_DOMAIN
ENV ALLOWED_DOMAIN=${ALLOWED_DOMAIN}
ARG INTERVAL_SECONDS
ENV INTERVAL_SECONDS=${INTERVAL_SECONDS}
ARG DB_NAME
ENV DB_NAME=${DB_NAME}
ARG DB_BUCKET_NAME
ENV DB_BUCKET_NAME=${DB_BUCKET_NAME}
ARG ASYNC_COUNT
ENV ASYNC_COUNT=${ASYNC_COUNT}

RUN apk update && apk add --no-cache ca-certificates curl

WORKDIR /app

COPY --from=build /src/crawler-task .

EXPOSE 3000

CMD [ "./crawler-task" ]
