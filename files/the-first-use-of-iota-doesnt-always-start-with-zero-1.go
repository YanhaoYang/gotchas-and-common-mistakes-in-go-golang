// https://play.golang.org/p/GiS_V9WCQKz

package main

import (
  "fmt"
)

const (
  azero = iota
  aone  = iota
)

const (
  info  = "processing"
  bzero = iota
  bone  = iota
)

func main() {
  fmt.Println(azero,aone) //prints: 0 1
  fmt.Println(bzero,bone) //prints: 1 2
}
