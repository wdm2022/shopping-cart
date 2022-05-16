# Shopping cart

Shopping cart microservice system with application-side distributed transactions.

Project for Web-scale Data Management course at TU Delft (IN4331)

## Development

### Prerequisites

* Go 1.18 or newer
* Docker
* [Task](https://taskfile.dev/) build tool

#### For Windows

* [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install)

#### When Docker builder image is not used

* [Protocol buffer compiler](https://grpc.io/docs/protoc-installation/) 
  with its [Go plugins](https://grpc.io/docs/languages/go/quickstart/) 

### Preparing the project

In order to generate files, necessary for development (eg. gRPC client/server from protobuf schema)
and perform vendoring, execute the `dev_prepare` task:

```shell
task -- task dev_prepare
```

### Running the services

Individual services can be built and executed locally with the `run_*` Makefile targets.
For example, in order to run the API gateway service, execute:

```shell
task -- task api_gateway:run
```

### Running Docker builder image to execute arbitrary commands

Simply run:

```shell
task
```

To execute arbitrary command `COMMAND` in the container,
pass it to the Task after `--`:

```shell
task -- COMMAND
```