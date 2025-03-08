# Simple Wallet Service

Simple Wallet Service is a golang microservice to serve disburse wallet balance

Maintainer : @nurlailatul

## Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install)
- Mysql database

### Setup

1. Install Go

2. Install MySql

3. Install dependencies
    ```
    make deps
    ```

4. Create database `wallet`

5. Copy and adjust the `config.yml`
    ```sh
    cp config.yml.tmpl config.yml
    ```

6. Migrate database
    ```
    make migrate
    ```

### Run the App

1. Load swagger
    ```
    make swagger
    ```

2. Compile and Run
    ```sh
    make run
    ```

3. Health check API http://localhost:8081/ping

4. Open Swagger in Browser http://localhost:8081/ping/swagger/index.html

5. Test di API in Swagger

## Other make commands

Run `make <command>` to build/test/run locally. Refer to the `Makefile` for available commands.

```
make test       # run tests (excluding E2E)
make e2e-test   # run end-to-end tests
make repo-test  # run unit tests for repo layer
make migrate    # run DB migrations
make run        # run the service at local machine on port 8081
```
