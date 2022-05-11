# Shopping cart

Shopping cart microservice system with application-side distributed transactions.

Project for Web-scale Data Management course at TU Delft (IN4331)

## Development

### Preparing the project

In order to generate files, necessary for development (eg. gRPC client/server from protobuf schema)
and perform vendoring, execute the `dev_prepare` Makefile target:

```shell
make dev_prepare
```

### Running the services

Individual services can be built and executed locally with the `run_*` Makefile targets.
For example, in order to run the API gateway service, execute:

```shell
make run_api_gateway
```