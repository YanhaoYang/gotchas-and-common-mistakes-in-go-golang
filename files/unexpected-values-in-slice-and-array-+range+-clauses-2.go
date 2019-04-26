// https://play.golang.org/p/aJP8_Z1su-z

package main

import "fmt"

func main() {  
    x := []string{"a","b","c"}

    for _, v := range x {
        fmt.Println(v) //prints a, b, c
    }
}
