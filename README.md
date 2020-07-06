![Go](https://github.com/geoffjay/7g-tooling/workflows/Go/badge.svg)
![release](https://github.com/geoffjay/7g-tooling/workflows/release/badge.svg)
[![CircleCI](https://circleci.com/gh/geoffjay/7g-tooling.svg?style=shield)](https://app.circleci.com/pipelines/github/geoffjay/7g-tooling)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/geoffjay/7g-tooling)
[![Go Report Card](https://goreportcard.com/badge/github.com/geoffjay/7g-tooling)](https://goreportcard.com/report/github.com/geoffjay/7g-tooling)

---

# Tooling Project

[WIP] pre-alpha, doesn't do anything, might burn your house down.

## CLI

### `deploy`

#### `initialize`

```shell
7g deploy initialize
```

#### `branch`

```shell
7g deploy branch feat/nonsense
```

#### `destroy`

```shell
7g deploy destroy
```

## Service

Building and running the web service during development is done using:

```shell
make setup
make
SG_CONFIG=configs/config.yml ./target/7g daemon
```

## API

### Swagger

Install the `swag` utility.

```shell
go get -u github.com/swaggo/swag/cmd/swag
```

To build Swagger documentation for the REST API.

```shell
make docs
make
```

Now when running the documentation should be available at http://localhost:3000/swagger/index.html.

### Ping

http://localhost:3000/v1/ping

### Deployment

#### Initialize

[priority]: 3

http://localhost:3000/v1/deploy/initialize

#### Branch

[priority]: 3

http://localhost:3000/v1/deploy/branch/<branch_name>

#### Network

[priority]: 1

http://localhost:3000/v1/deploy/network

_should there be something for domain specification_ /<domain_name> (? or id)

GET

returns?

POST

takes xlsx file?

### Automate

[priority]: 2

http://localhost:3000/v1/automate/network

GET

returns?

POST

takes xlsx file?

#### Configuration

[priority]: 2

- API key
- server settings?

http://localhost:3000/v1/config/apikey/

GET

POST

## Contributing

To contribute to development ...

```shell
python3 -m venv .venv
pip install pre-commit
pre-commit install
```
