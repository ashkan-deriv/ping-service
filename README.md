# Ping Service

A simple Go service that responds to `/ping` requests with "pong".

## Running the service

```bash
go run main.go
```

The service will start on port 8080. You can test it using:

```bash
curl http://localhost:8080/ping
```