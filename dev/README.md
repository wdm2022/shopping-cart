# Development

## Prerequisites

* Go 1.18 or newer
* Docker 20.10 or newer
* [Task](https://taskfile.dev/) build tool

### For Windows

* [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install)
    * Docker tasks work natively only under Linux and (possibly) macOS.
      On Windows use WSL2.

### When Docker builder image is not used

* [Protocol buffer compiler](https://grpc.io/docs/protoc-installation/)
  with its [Go plugins](https://grpc.io/docs/languages/go/quickstart/)

## Preparing the project

In order to generate files, necessary for development (eg. gRPC client/server from protobuf schema)
and perform vendoring, execute the `dev_prepare` task:

```shell
task -- task prepare
```

## Starting runtime dependencies

Runtime dependencies (eg. MongoDB) are launched as Docker containers and are managed by Docker Compose.
The Compose operations are wrapped into Task tasks.

Make sure that the Docker daemon is running and execute:

```shell
task dev:compose:up
```

In order to remove the containers and their data, run:

```shell
task dev:compose:down
```

To execute arbitrary Docker Compose command `<COMMAND>`, call:

```shell
task dist:compose -- <COMMAND>
```

### Connecting to MongoDB

Since the local MongoDB instance is run is the replica set mode and advertises the internal Docker hostname 
to other containers, in order to connect to it from the host machine it is necessary to configure the client to 
ignore the advertised hostname and connect to the instance directly, 
by specifying the `directConnection` option in the connection string:

```
mongodb://root:password@localhost:27017/?directConnection=true
```

## Running the services

Individual services can be built and executed locally with the `run_*` Task tasks.
For example, in order to run the API gateway service, execute:

```shell
task -- task api_gateway:run
```

## Running Docker builder image to execute arbitrary commands

In order to launch Docker container with the project's directory mounted, execute:

```shell
task
```

To execute arbitrary command `COMMAND` in the container,
pass it to the Task after `--`:

```shell
task -- COMMAND
```
