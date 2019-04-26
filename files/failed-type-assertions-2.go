// https://play.golang.org/p/_PA90RR37Xx

package main

import "fmt"

func main() {  
    var data interface{} = "great"

    if res, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",res)
    } else {
        fmt.Println("[not an int] value =>",data) 
        //prints: [not an int] value => great (as expected)
    }
}
