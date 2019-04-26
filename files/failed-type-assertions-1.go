// https://play.golang.org/p/HZHklPHFJuj

package main

import "fmt"

func main() {  
    var data interface{} = "great"

    if data, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",data)
    } else {
        fmt.Println("[not an int] value =>",data) 
        //prints: [not an int] value => 0 (not "great")
    }
}
