package main

// The large comment block below is processed by cgo.
// See https://golang.org/cmd/cgo/

/*
#cgo CFLAGS: -I VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include "VBoxCAPIGlue.h"
#include "VBoxCAPIGlue.c"

#include <stdio.h>

void VirtualBoxInGo() {
  printf("VBoxCGlueInit HRESULT: %x\n", VBoxCGlueInit());
  printf("VirtualBox Version: %d\n", g_pVBoxFuncs->pfnGetVersion());
  printf("API Version: %d\n", g_pVBoxFuncs->pfnGetAPIVersion());

  IVirtualBoxClient* client;
  printf("pfnClientInitialize HRESULT: %x\n",
         g_pVBoxFuncs->pfnClientInitialize(NULL, &client));
  printf("IVirtualBoxClient: %p\n", client);
}

*/
import "C"  // cgo's virtual package

func main() {
  C.VirtualBoxInGo()
}
