// https://play.golang.org/p/MvPVURdUsNW

package main

import (  
    "fmt"
    "time"
)

func main() {  
    var ch chan int
    for i := 0; i < 3; i++ {
        go func(idx int) {
            ch <- (idx + 1) * 2
        }(i)
    }

    //get first result
    fmt.Println("result:",<-ch)
    //do other work
    time.Sleep(2 * time.Second)
}
