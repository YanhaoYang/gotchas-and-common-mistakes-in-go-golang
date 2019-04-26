// https://play.golang.org/p/hgtiHYRPx3_m

package main

import (  
    "fmt"
    "reflect"
    "encoding/json"
)

func main() {  
    var str string = "one"
    var in interface{} = "one"
    fmt.Println("str == in:",str == in,reflect.DeepEqual(str, in)) 
    //prints: str == in: true true

    v1 := []string{"one","two"}
    v2 := []interface{}{"one","two"}
    fmt.Println("v1 == v2:",reflect.DeepEqual(v1, v2)) 
    //prints: v1 == v2: false (not ok)

    data := map[string]interface{}{
        "code": 200,
        "value": []string{"one","two"},
    }
    encoded, _ := json.Marshal(data)
    var decoded map[string]interface{}
    json.Unmarshal(encoded, &decoded)
    fmt.Println("data == decoded:",reflect.DeepEqual(data, decoded)) 
    //prints: data == decoded: false (not ok)
}
