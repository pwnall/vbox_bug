# VirtualBox XPCOM Bug Repro

This repository demonstrates a bug in VirtualBox's XPCOM initialization when
used in Go on Mac OSX.


## Bug Reproduction

```bash
go build
./vbox_bug
```

On OSX, I get the output below, which indicates a failure.

```
VBoxCGlueInit HRESULT: 0
VirtualBox Version: 4003028
API Version: 4003
pfnClientInitialize HRESULT: 80004005
IVirtualBoxClient: 0x0
```

The code works on Fedora Linux, whre `pfnClientInitialize` returns 0. When
using a distribution-packaged VirtualBox, `VBOX_APP_HOME` most likely needs to
be set, e.g.

```bash
VBOX_APP_HOME=/usr/lib64/virtualbox ./vbox_bug
```


## Source Code

The reproduction code is all in `vbox_bug.go`. Everything else is boilerplate
needed to get the code to build.


## Copyright Notice

For ease of reproduction, this repository includes portions of the VirtualBox
SDK, in the `VirtualBoxSDK` directory. The files in there have copyright
notices.
