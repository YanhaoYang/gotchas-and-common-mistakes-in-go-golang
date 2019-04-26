// https://play.golang.org/p/-99cuq9fOmQ

package main

import "fmt"

type data struct {  
    name string
}

func main() {  
    m := map[string]data {"x":{"one"}}
    r := m["x"]
    r.name = "two"
    m["x"] = r
    fmt.Printf("%v",m) //prints: map[x:{two}]
}
