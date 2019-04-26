// https://play.golang.org/p/O1QUglcVjsu

package main

/*
#include <stdlib.h>
*/
import (
  "C"
  "unsafe"
)

func main() {
  cs := C.CString("my go string")
  C.free(unsafe.Pointer(cs))
}
