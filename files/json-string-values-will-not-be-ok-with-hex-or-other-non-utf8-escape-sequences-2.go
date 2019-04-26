// https://play.golang.org/p/_f9OxD8x12S

package main

import (
  "fmt"
  "encoding/json"
)

type config struct {
  Data string `json:"data"`
}

func main() {
  raw := []byte(`{"data":"\\xc2"}`)
  
  var decoded config
  
  json.Unmarshal(raw, &decoded)
  
  fmt.Printf("%#v",decoded) //prints: main.config{Data:"\\xc2"}
  //todo: do your own hex escape decoding for decoded.Data  
}
