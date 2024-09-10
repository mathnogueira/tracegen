# tracegen

A utility tool capable of generating OpenTelemetry traces. It can work in two modes:

* cli
* http server

Both modes generates traces and send them to `http://localhost:4317` (you can override this by passing the option `-c <your-collector-endpoint`

## CLI
Generates a single trace every time you run it.
```
tracegen -s 10 -n 30 start
```

## HTTP Server
Creates an HTTP Server that generates a new server everytimes the endpoint `GET /` is called.
```
tracegen -s 10 -n 30 -p 9095 serve
```

## Example of trace
Viewed on [Tracetest](https://app.tracetest.io)
![image](https://github.com/mathnogueira/tracegen/assets/2704737/20bd1735-6db3-4a03-8bc7-8a4960723c72)


## Configuring the collector
If you pass the option `--collector <endpoint>` `tracegen` will connect to this gRPC endpoint and send traces there using the OpenTelemetry Protocol.
