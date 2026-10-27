# gtfs2postgis


### enviroments variable

- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_DB
- POSTGRES_USER
- POSTGRES_PASSWORD
- SLACK_URL

### health

The worker listens on port 8080 and answers `GET /health` with `{"status":"ok"}`; the deployment's readiness and liveness probes use it.
