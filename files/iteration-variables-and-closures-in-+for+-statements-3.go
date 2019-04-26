// https://play.golang.org/p/Ae3jaVSIC8m

package main

import (  
    "fmt"
    "time"
)

func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func(in string) {
            fmt.Println(in)
        }(v)
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
