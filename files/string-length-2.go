// https://play.golang.org/p/VYL8ixHjblw

package main

import (  
    "fmt"
    "unicode/utf8"
)

func main() {  
    data := "♥"
    fmt.Println(utf8.RuneCountInString(data)) //prints: 1
