// https://play.golang.org/p/KQPoVTEfRgc

package main

import "fmt"

func main() {  
    doit := func(arg int) interface{} {
        var result *struct{} = nil

        if(arg > 0) {
            result = &struct{}{}
        }

        return result
    }

    if res := doit(-1); res != nil {
        fmt.Println("good result:",res) //prints: good result: <nil>
        //'res' is not 'nil', but its value is 'nil'
    }
}
