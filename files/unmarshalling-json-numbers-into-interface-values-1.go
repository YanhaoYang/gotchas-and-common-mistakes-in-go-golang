// https://play.golang.org/p/MKAjloaT6hs

package main

import (  
  "encoding/json"
  "fmt"
)

func main() {  
  var data = []byte(`{"status": 200}`)

  var result map[string]interface{}
  if err := json.Unmarshal(data, &result); err != nil {
    fmt.Println("error:", err)
    return
  }

  var status = result["status"].(int) //error
  fmt.Println("status value:",status)
}
