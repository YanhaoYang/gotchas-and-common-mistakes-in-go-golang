// https://play.golang.org/p/nlz-7SHQBj2

package main

import "fmt"

func main() {  
    m := map[string]int{"one":1,"two":2,"three":3,"four":4}
    for k,v := range m {
        fmt.Println(k,v)
    }
}
