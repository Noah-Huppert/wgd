# Remote Procedural Call
Cross platform interface.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)

# Overview
GRPC is used. The [`interface.proto`](./interface.proto) file defines the 
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
make
```
