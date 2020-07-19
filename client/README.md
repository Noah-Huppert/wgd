# Wgd Client
GUI application which helps users join and stay in a Wireguard network.

# Table Of Contents
- [Overview](#overview)
- [Development](#development)

# Overview
The goal of this tool is to help users with very little technical knowledge join
a Wireguard network and configure their machine correctly.

To onboard users into the network a static binary of this tool is sent to them 
along with a secret invite code. When the user starts the GUI they are prompted
for the invite code which will then be used to authorize the registration of a
new user in the server's registry.

The server calculates the correct topology from the network which this client 
then pulls down and uses to configure the Wireguard network interface and 
routing tables. This configuration occurs in the background and hopefully the 
user should only see a loading bar for a few seconds.

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

