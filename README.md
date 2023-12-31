# tracegen

A utility tool capable of generating OpenTelemetry traces. It can work in two modes:

* cli
* http server

## CLI
```
tracegen -s 10 -n 30 start
```

## HTTP Server
```
tracegen -s 10 -n 30 -p 9095 serve
```

## Example of trace
Viewed on [Tracetest](https://app.tracetest.io)
![image](https://github.com/mathnogueira/tracegen/assets/2704737/20bd1735-6db3-4a03-8bc7-8a4960723c72)


## Configuring the collector
If you pass the option `--collector <endpoint>` `tracegen` will connect to this gRPC endpoint and send traces there using the OpenTelemetry Protocol.
