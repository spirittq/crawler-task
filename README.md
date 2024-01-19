# Crawler Task

Consists of 2 apps, `data_manager` (gRPC server) and `crawler` (gRPC client)

## data_manager

Requires these env variables:

```
SERVER_PORT=
SERVER_API_PORT=
DB_NAME=
DB_BUCKET_NAME=
```

Upon running, creates db file (based on DB_NAME) and bucket, where all data is stored.

Also has a simple API for health_check & fetching all db items (for convenience, final result is provided in `results.json` file)

## crawler

Requires these environment variables:

```
SCRAPE_URL=
ALLOWED_DOMAIN=
SERVER_PORT=
ASYNC_COUNT=
SERVER_DOMAIN=
```

It is a one-time task that exits upon finishing scraping data.

---
Both apps can be build and run with `docker-compose up` command and deployed and run with `kubectl apply -f deployment.yaml` command.

---
# Future Improvements

1. Increase test coverage (database integration, gRPC communication, crawler functionality)
2. Potentially, API server could be a separate micro-service, as well use [gin framework](https://github.com/gin-gonic/gin)
3. Currently, crawler is a kubernetes cronjob that triggers once an hour, to implement initial run starting from deployment, additional job had to be configured.
4. Crawler is dedicated to the webpage in question, need to check if there is a possibility to make it more generic (register as a handler?)
