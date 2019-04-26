// https://play.golang.org/p/0tBzwYY6PAV

package main

import "fmt"

type data struct {  
    name string
}

func main() {  
    s := []data {{"one"}}
    s[0].name = "two" //ok
    fmt.Println(s)    //prints: [{two}]
}
