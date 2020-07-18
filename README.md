# wgd
Wireguard configuration tool.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)

# Overview
This tool tried to make it easy to manage a Wireguard virtual private network.
It is meant both for end users and network administrators.

Wgd is composed of a server which keeps track of user information and a client
which configures peer's Wireguard interfaces to join the network.

The purpose of this tool is to make it easy to bring non-technical users into
a Wireguard network.

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
