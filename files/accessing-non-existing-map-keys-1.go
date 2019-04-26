// https://play.golang.org/p/NUo8F0AvNjv

package main

import "fmt"

func main() {  
    x := map[string]string{"one":"a","two":"","three":"c"}

    if v := x["two"]; v == "" { //incorrect
        fmt.Println("no entry")
    }
}
