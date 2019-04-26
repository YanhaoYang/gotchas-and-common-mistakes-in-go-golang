// https://play.golang.org/p/KzlBGMdTmd-

package main

import "fmt"

func main() {  
    loop:
        for {
            switch {
            case true:
                fmt.Println("breaking out...")
                break loop
            }
        }

    fmt.Println("out!")
}
