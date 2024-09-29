# Tracee Client

## About Tracee Client

Tracee client is a powerful for  using tracee with an api.
Tracee client let you chose how you want use Tracee for your need.
Tracee Client is a very easy to use and learn CLI tool .

## Quick Start

IN order to configure Tracee Client first you need to check for an active Tracee API server,
you can also run a Tracee Api server on you PC.

### Config Tracee API server
IN order to configure Tracee Api server  first you need to create a configure file for tracee with gRPC section on.
please refer to [Tracee Docs about Configuration files](https://aquasecurity.github.io/tracee/latest/docs/install/config/).

for example a short config yaml file:
```yaml
grpc-listen-addr: tcp:50051
```
### compile Tracee Client
In order to use Tracee client you need to compile Tracee client you can do it  by running:

```bash 
make 
```

## Tracee API Client

After you compile Tracee client you can run it

```bash
TraceeClient 
```

It will print the help commend for you to view how you can interact with the client.

you can view more about each commend in the [docs section](#add docs section)





