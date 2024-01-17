#!/bin/sh
source data_manager/.env
docker build -f data_manager/Dockerfile . -t data_manager:latest \
    --build-arg SERVER_PORT=$SERVER_PORT \
    --build-arg SERVER_API_PORT=$SERVER_API_PORT \
    --build-arg DB_NAME=$DB_NAME \
    --build-arg DB_BUCKET_NAME=$DB_BUCKET_NAME
