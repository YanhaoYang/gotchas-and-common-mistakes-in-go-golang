// https://play.golang.org/p/_PSKCLcpvGA

package main

import (  
    "fmt"
    "time"
)

type field struct {  
    name string
}

func (p *field) print() {  
    fmt.Println(p.name)
}

func main() {  
    data := []field{{"one"},{"two"},{"three"}}

    for _,v := range data {
        v := v
        go v.print()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
