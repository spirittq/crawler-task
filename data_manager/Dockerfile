FROM golang:1.21-alpine as build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

COPY data_manager/ ./data_manager/
COPY shared/ ./shared/

WORKDIR /src/data_manager

RUN go build -o data_manager

FROM alpine:3.19

ARG SERVER_PORT
ENV SERVER_PORT=${SERVER_PORT}
ARG SERVER_API_PORT
ENV SERVER_API_PORT=${SERVER_API_PORT}
ARG DB_NAME
ENV DB_NAME=${DB_NAME}
ARG DB_BUCKET_NAME
ENV DB_BUCKET_NAME=${DB_BUCKET_NAME}

RUN apk update && apk add --no-cache ca-certificates curl

WORKDIR /app

COPY --from=build /src/data_manager/data_manager .

EXPOSE 3000
EXPOSE 50051

CMD [ "./data_manager" ]
