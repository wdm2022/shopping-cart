## Local deployment

To build and deploy the whole system locally, execute:

```shell
task dist:compose_up
```

The environment can be taken down with:

```shell
task dist:compose_down
```

Logs can be viewed by executing:

```shell
task dist:compose_logs
```

To display only the logs of one service - for instance the API gateway - execute:

```shell
task dist:compose_logs -- api-gateway
```

To execute arbitrary Docker Compose command `<COMMAND>`, call:

```shell
task dist:compose -- <COMMAND>
```