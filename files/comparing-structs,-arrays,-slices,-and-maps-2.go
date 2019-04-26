// https://play.golang.org/p/E2lw-mvxw-A

package main

import "fmt"

type data struct {  
    num int                //ok
    checks [10]func() bool //not comparable
    doit func() bool       //not comparable
    m map[string] string   //not comparable
    bytes []byte           //not comparable
}

func main() {  
    v1 := data{}
    v2 := data{}
    fmt.Println("v1 == v2:",v1 == v2)
}
