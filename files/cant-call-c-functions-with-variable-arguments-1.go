// https://play.golang.org/p/8sDX4tHI2gJ

package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
  "unsafe"
)

func main() {
  cstr := C.CString("go")
  C.printf("%s\n",cstr) //not ok
  C.free(unsafe.Pointer(cstr))
}
