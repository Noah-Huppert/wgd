# Wgd Server
Backend server which stores user information and calculates the correct 
network topology.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)

# Overview
This tool is meant to be run by the administrator of the Wireguard network. It 
is a centralized service which clients contact to get information about other 
users on the network and network topology information.

A pure Wireguard network setup does not require a centralized service like this.
Instead it is decentralized with no single authority. By adding this centralized
service the root of trust in the network becomes the network administrator. The
use case this tool is designed for assumes that the administrator is someone 
users know personally, and whom has is trusted. The administrator is probably 
running the Wireguard network so people they know can benefit.

# Development
Go is used with Mongodb as a data store and GRPC as a transport utility.

## Dependencies
Protocol buffers are used for client server communication. Follow all
instructions in the
[`../rpc/README.md`](../rpc/README.md#Development) file 
development section.

## Execute
To run the server:

```
go run .
```
