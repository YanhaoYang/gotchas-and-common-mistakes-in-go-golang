// https://play.golang.org/p/C7UVunfhtp1

package main

import (
  "fmt"
  "encoding/json"
  "bytes"
)

func main() {
  data := "x < y"
  
  raw,_ := json.Marshal(data)
  fmt.Println(string(raw))
  //prints: "x \u003c y" 