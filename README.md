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

6. Load swagger
    ```
    make swagger
    ```

7. Migrate database
    ```
    make migrate
    ```

### Run the App

1. Compile and Run
    ```sh
    make run
    ```

2. Health check API http://localhost:8081/ping

3. Open Swagger in Browser http://localhost:8081/ping/swagger/index.html
   You can use browser extension to modify request header, to add `x-api-key`.
   In chrome, you can use ModHeader https://modheader.com/?ref=me&product=ModHeader&version=7.0.7&browser=chrome

4. Test di API in Swagger

## Other make commands

Run `make <command>` to build/test/run locally. Refer to the `Makefile` for available commands.

```
make test           # run tests
make migrate-down   # rollback DB migrations
make vet            # run go vet
```
