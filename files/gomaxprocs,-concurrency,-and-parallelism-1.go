// https://play.golang.org/p/4hFCnD2aytK

package main

import (  
    "fmt"
    "runtime"
)

func main() {  
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: X (1 on play.golang.org)
    fmt.Println(runtime.NumCPU())       //prints: X (1 on play.golang.org)
    runtime.GOMAXPROCS(20)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 20
    runtime.GOMAXPROCS(300)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 256
}
