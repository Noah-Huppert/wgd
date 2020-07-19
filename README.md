# wgd
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

Wgd is not for everyone. It is a set of opinionated tools which try to make it 
easier for technical and non-technical users alike to use Wireguard. It is
designed for users who already trust each other and simply want to connect over 
a VPN. Think of Wgd as a self hosted alternative to LogMeIn Hamachi.

Linux, Windows, and Mac OS X will be supported.

# Development
Go is used due to its cross platform nature.

Currently these steps have only been confirmed on Linux. Windows support
is coming soon. Followed by Mac OSX.

## Dependencies
Most dependencies are managed by the Go modules automatically.

For the GUI to be built on Linux the following headers must be installed:

| Dependency                  | Void Linux Package  |
| --                          | --                  |
| `X11/Xcursor/Xcursor.h`     | `libXcursor-devel`  |
| `X11/extensions/Xrandr.h`   | `libXrandr-devel`   |
| `X11/extensions/Xinerama.h` | `libXinerama-devel` |
| `X11/extensions/XInput2.h`  | `libXi-devel`       |
| `GL/glx.h`                  | `libglvnd-devel`    |
| Xxf86vm                     | `libXxf86vm-devel`  |

## Run
Execute:

```
go run .
```
