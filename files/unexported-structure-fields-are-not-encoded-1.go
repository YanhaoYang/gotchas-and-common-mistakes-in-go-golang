// https://play.golang.org/p/RamGD8MqD2G

package main

import (  
    "fmt"
    "encoding/json"
)

type MyData struct {  
    One int
    two string
}

func main() {  
    in := MyData{1,"two"}
    fmt.Printf("%#v\n",in) //prints main.MyData{One:1, two:"two"}

    encoded,_ := json.Marshal(in)
    fmt.Println(string(encoded)) //prints {"One":1}

    var out MyData
    json.Unmarshal(encoded,&out)

    fmt.Printf("%#v\n",out) //prints main.MyData{One:1, two:""}
}
