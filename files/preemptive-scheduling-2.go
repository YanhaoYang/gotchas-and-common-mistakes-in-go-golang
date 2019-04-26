// https://play.golang.org/p/x-HvaWljfK6

package main

import "fmt"

func main() {  
    done := false

    go func(){
        done = true
    }()

    for !done {
        fmt.Println("not done!") //not inlined
    }
    fmt.Println("done!")
}
