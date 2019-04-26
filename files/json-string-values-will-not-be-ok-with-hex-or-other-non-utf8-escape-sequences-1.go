// https://play.golang.org/p/gvsIfehmdhH

package main

import (
  "fmt"
  "encoding/json"
)

type config struct {
  Data string `json:"data"`
}

func main() {
  raw := []byte(`{"data":"\xc2"}`)
  var decoded config

  if err := json.Unmarshal(raw, &decoded); err != nil {
        fmt.Println(err)
    //prints: invalid character 'x' in string escape code
    }
  
}
