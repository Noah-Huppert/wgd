# Rhapso
Builds quality of life features on top of Wireguard.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)

# Overview
This suite of tools tries to make it easy for users to join and be part of 
a Wireguard network as well as for administrators to manage Wireguard networks.

The suite is composed of two tools, each meant for different types of users:

- **Server**: Centralized registry of user information. Makes administering
  the network easier. 
- **Client**: GUI tool to help peers join and stay part of the network.

Rhapso is not for everyone. It is a set of opinionated tools which try to make
it easier for technical and non-technical users alike to use Wireguard. It is
designed for users who already trust each other and simply want to connect over 
a VPN. Think of Rhapso as a self hosted alternative to LogMeIn Hamachi.

Linux, Windows, and Mac OS X will be supported.

[Rhapso was a minor goddess of sewing](https://en.wikipedia.org/wiki/Rhapso) and
this tool weaves together a Wireguard network, hence the name.

# Development
See [`client/README.md`](./client/README.md) for instructions specific to the 
client. As well as [`server/README.md`](./server/README.md) for the server.
