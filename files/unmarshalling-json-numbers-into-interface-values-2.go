// https://play.golang.org/p/sux7Jud97vR

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

  var status = uint64(result["status"].(float64)) //ok
  fmt.Println("status value:",status)
}
