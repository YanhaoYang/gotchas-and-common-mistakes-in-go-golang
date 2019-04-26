// https://play.golang.org/p/hi0dt9vq5Mp

package main

import (
  "fmt"
)

func main() {
  i := 1
  defer func (in *int) { fmt.Println("result =>", *in) }(&i)
  
  i = 2
  //prints: result => 2
}
