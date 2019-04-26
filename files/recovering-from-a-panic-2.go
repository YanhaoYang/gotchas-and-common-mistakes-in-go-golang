// https://play.golang.org/p/LBwnYATIUYA

package main

import "fmt"

func main() {  
    defer func() {
        fmt.Println("recovered:",recover())
    }()

    panic("not good")
}
