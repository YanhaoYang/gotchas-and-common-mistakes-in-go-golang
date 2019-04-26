// https://play.golang.org/p/Nv8UuONUQzn

package main

import (
  "fmt"
  "encoding/json"
)

type config struct {
  Data []byte `json:"data"`
}

func main() {
  raw := []byte(`{"data":"wg=="}`)
  var decoded config
  
  if err := json.Unmarshal(raw, &decoded); err != nil {
          fmt.Println(err)
      }
  
  fmt.Printf("%#v",decoded) //prints: main.config{Data:[]uint8{0xc2}}
}
