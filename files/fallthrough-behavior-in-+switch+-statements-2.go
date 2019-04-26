// https://play.golang.org/p/qfMdB3XnGmW

package main

import "fmt"

func main() {  
    isSpace := func(ch byte) bool {
        switch(ch) {
        case ' ', '\t':
            return true
        }
        return false
    }

    fmt.Println(isSpace('\t')) //prints true (ok)
    fmt.Println(isSpace(' '))  //prints true (ok)
}
