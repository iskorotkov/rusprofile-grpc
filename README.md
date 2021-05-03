# gRPC wrapper for rusprofile.ru

[![Go](https://github.com/iskorotkov/rusprofile-grpc/actions/workflows/go.yml/badge.svg)](https://github.com/iskorotkov/rusprofile-grpc/actions/workflows/go.yml)
[![CodeQL](https://github.com/iskorotkov/rusprofile-grpc/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/iskorotkov/rusprofile-grpc/actions/workflows/codeql-analysis.yml)

The wrapper provides access to [rusprofile.ru](https://www.rusprofile.ru/) data via gRPC. HTTP API is available via HTTP-to-gRPC gateway.

- [gRPC wrapper for rusprofile.ru](#grpc-wrapper-for-rusprofileru)
  - [Run](#run)
  - [Test](#test)
    - [gRPC](#grpc)
    - [HTTP](#http)
    - [Browser](#browser)
  - [Project structure](#project-structure)

## Run

To run Docker container, execute the following:

```shell
docker run -it -p 8080:8080 -p 8888:8888 --rm iskorotkov/rusprofile-grpc:v1.0.1
```

## Test

### gRPC

Use `grpcurl` (`curl` for gRPC) to test gRPC API:

```shell
# Returns 'not found'
grpcurl -plaintext -d '{"INN": "123"}' localhost:8888 rusprofile.v1.CompanyFinder/ByINN

# Returns 'Xsolla'
grpcurl -plaintext -d '{"INN": "5902879646"}' localhost:8888 rusprofile.v1.CompanyFinder/ByINN
```

### HTTP

Use `curl` to test HTTP API:

```shell
# Returns 'not found'
curl localhost:8080/v1/company/123

# Returns 'Xsolla'
curl localhost:8080/v1/company/5902879646
```

### Browser

Open [Swagger UI](http://localhost:8080/swagger-ui/) in your browser.

## Project structure

- api - proto files and buf configuration.
  - openapiv2 - generated swagger.json file.
- build - Dockerfile.
- cmd - main package.
- pkg - generated Go files and implementation of gRPC server for getting data from rusprofile.ru.
- static/web - static files for Swagger UI page.
- tools - imports for tools used in code generation.
