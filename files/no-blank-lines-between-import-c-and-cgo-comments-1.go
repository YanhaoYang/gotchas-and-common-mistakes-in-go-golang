// https://play.golang.org/p/LGd2CEcC65S

package main

/*
#include <stdlib.h>
*/

import "C"

import (
  "unsafe"
)

func main() {
  cs := C.CString("my go string")
  C.free(unsafe.Pointer(cs))
}
