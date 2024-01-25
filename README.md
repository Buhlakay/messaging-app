# Project Overview

A side project to learn Go & PostgreSQL. This is a messaging application (Facebook/Telegram) clone with minimal UI.

Uses WebSockets to communicate with two send/receive message services, who are publishers/subscribers, respectively to a Kafka broker. Persists messaging data to a PostgreSQL instance.

![High-level design diagram](docs/messaging-app.png)

# Setup
Create an `.env` file at the root directory of this project, and define the `POSTGRES_USER`, `POSTGRES_PASSWORD`, and `POSTGRES_DB` environment variables to be referenced by the PostgreSQL database instance. 