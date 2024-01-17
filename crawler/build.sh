#!/bin/sh
source crawler/.env
docker build -f crawler/Dockerfile . -t crawler:latest \
    --build-arg SCRAPE_URL=$SCRAPE_URL \
    --build-arg ALLOWED_DOMAIN=$ALLOWED_DOMAIN \
    --build-arg SERVER_PORT=$SERVER_PORT \
    --build-arg ASYNC_COUNT=$ASYNC_COUNT \
    --build-arg SERVER_DOMAIN=$SERVER_DOMAIN
