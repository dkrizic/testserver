# Sample application for GraphQL

## Features

* Multi stage: Uses Multi stage build for building the container
* Database ownershop: Takes care of the database schema by itself
* all runtime parameters are passed either via environment variables or command line arguments
* JSON logging incuding trace_id and span_id
* OpenTelemetry: Exploses traces in OTLP to a configurable endpoint
* Version awareness: Exploses the own version via /version endpoint
* Health endpoint: Exploses a health endpoint for readiness and liveness probes

