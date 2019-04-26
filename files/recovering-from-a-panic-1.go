// https://play.golang.org/p/PFKW8ozRKvN

package main

import "fmt"

func main() {  
    recover() //doesn't do anything
    panic("not good")
    recover() //won't be executed :)
    fmt.Println("ok")
}
