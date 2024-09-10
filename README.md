# tracegen

A utility tool capable of generating OpenTelemetry traces. It can work in two modes:

* cli
* http server

Both modes generates traces and send them to `http://localhost:4317` (you can override this by passing the option `-c <your-collector-endpoint`

## Options

| flag | shortcut | default | description |
| ---- | -------- | ------- | ----------- |
| `-c`   | `--collector` | `localhost:4317` | OpenTelemetry collector gRPC endpoint where the traces will be sent to |
| `-s`   | `--services` | `1` | number of fake services used to generate the trace |
| `-n` | `--spans` | `10` | approximated number of spans you want to generate. Each service will have (spans / services) spans |
| ` ` | `--tracestate` | ` ` | list of values injected into the tracestate of each trace |

## Run modes

### CLI
Generates a single trace every time you run it.
```
tracegen -s 10 -n 30 start
```

### HTTP Server
Creates an HTTP Server that generates a new server everytimes the endpoint `GET /` is called.
```
tracegen -s 10 -n 30 -p 9095 serve
```

In this mode, there's an extra option: `--port` or `-p` which specifies the number of the port that will be used by the web server

## Example of trace
Viewed on [Tracetest](https://app.tracetest.io)
![image](https://github.com/mathnogueira/tracegen/assets/2704737/20bd1735-6db3-4a03-8bc7-8a4960723c72)


## Configuring the collector
If you pass the option `--collector <endpoint>` `tracegen` will connect to this gRPC endpoint and send traces there using the OpenTelemetry Protocol.
