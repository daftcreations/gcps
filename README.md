# Google Cloud Profile Switcher (GCPS)

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/) [![Go Report Card](https://goreportcard.com/badge/github.com/Naman2706/gcps)](https://goreportcard.com/report/github.com/Naman2706/gcps) [![GitHub license](https://img.shields.io/github/license/Naman2706/gcps)](https://github.com/Naman2706/gcps/blob/master/LICENSE) [![GitHub release](https://img.shields.io/github/v/tag/Naman2706/gcps)](https://github.com/Naman2706/gcps/releases)

Switch between Google Cloud Profiles easily

## Usage

<img src="docs/demo.gif" width="100%">

- Switch through list `gcps` will return configured profile.

- Switch to particular profile `gcps {profile_name}`

- Switch to previous profile `gcps -`

## Installation

Golang

```shell
go install github.com/naman2706/gcps@latest
```

Or

```shell
# For linux and mac systems
curl -Lso ./gcps https://github.com/Naman2706/gcps/releases/download/latest/gcps_$(uname -s)_$(uname -m)
chmod +x ./gcps
./gcps
```

```powershell
# For windows powershell
## x84_64/AMD64
Invoke-WebRequest https://github.com/Naman2706/gcps/releases/download/latest/gcps_Windows_x86_64.exe -O gcps.exe
## i386
Invoke-WebRequest https://github.com/Naman2706/gcps/releases/download/latest/gcps_Windows_Windows_i386.exe -O gcps.exe
## armv6
Invoke-WebRequest https://github.com/Naman2706/gcps/releases/download/latest/gcps_Windows_Windows_armv6.exe -O gcps.exe
## armv7
Invoke-WebRequest https://github.com/Naman2706/gcps/releases/download/latest/gcps_Windows_Windows_armv7.exe -O gcps.exe
## armv64
Invoke-WebRequest https://github.com/Naman2706/gcps/releases/download/latest/gcps_Windows_Windows_arm64.exe -O gcps.exe

ICACLS ".\gcps.exe" /grant:r "users:(RX)" /C
.\gcps.exe
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/Naman2706/gcps.svg)](https://starchart.cc/Naman2706/gcps)

---

> > *May the source be with you*
