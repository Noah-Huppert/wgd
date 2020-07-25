# Remote Procedural Call
Cross platform interface.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)
- [RPC Certificates](#rpc-certificates)

# Overview
GRPC is used. The [`rpc.proto`](./rpc.proto) file defines the 
interface which the client and server tools use to communicate.

GRPC generates a Go module which contains the interface bindings. This must
be built, follow instructions in the [development](#development) section.

# Development
A Go module is generated based on the interface protocol buffer definition file.

## Dependencies
Follow the [GRPC for Go](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
instructions to setup your environment.

## Build
To build the Go module run:

```
make generate
```

# RPC Certificates
To ensure secure communication between RPC client's and server's TLS 
encryption is utilized. 

Development keys with the names `dev-ca.key`, `dev-ca.cert`, `dev-server.key`,
`dev-server.csr`, and `dev-server.pem` are provided and used by the client and 
server by default.

In production custom keys are used. To generate these keys first edit the
`certs-server.conf` file. You may want to edit the common name, alternate 
names, plus `C`, `ST`, and `O` parameters.

Next run the `certs` make target. Ensure that the TLS certificate `C`, `ST`, and
`O` parameters which you set in the `certs-server.conf` file are passed to the 
make target by setting the `CERTS_COUNTRY`, `CERTS_STATE`, and 
`CERTS_ORG` variables. The `CERTS_PREFIX` variable is a prefix which will be 
placed before every generated file. As long as this prefix is not `dev-` the 
generated files will be gitignored.

```
make certs CERTS_PREFIX=prod- CERTS_COUNTRY=00 CERTS_STATE=00 CERTS_ORG=00
```
