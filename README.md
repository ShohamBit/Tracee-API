# Tracee-API

you can run tracee with an api written in `grpc-go`.

you can run tracee with a grpc option and then you can run tracee as an api.

## About Tracee API

you can config tracee api as you need it, for quick start view the rest of the docs


## Quick Start

you can run Tracee  api in 3 different ways but all of them are the same structure

first write a config file with the grpc section

```yaml
grpc:
    - tcp:50051
```
run this config file in your prefer way

### Commend Line

```bash
tracee --config /path/to/config
```

### Docker

```bash
docker run tracee --config /path/to/config
```

### Kubernetes

## Tracee API Client

After you run tracee with grpc option on, you need to tracee client.

you can tracee client like

```bash
how to run tracee client 
```





