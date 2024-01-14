#!/bin/sh
source .env
docker build . -t crawler-task \
    --build-arg SCRAPE_URL=$SCRAPE_URL \
    --build-arg ALLOWED_DOMAIN=$ALLOWED_DOMAIN \
    --build-arg INTERVAL_SECONDS=$INTERVAL_SECONDS \
    --build-arg DB_NAME=$DB_NAME \
    --build-arg DB_BUCKET_NAME=$DB_BUCKET_NAME \
    --build-arg ASYNC_COUNT=$ASYNC_COUNT