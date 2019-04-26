// https://play.golang.org/p/twTKuVNDaVX

package main

import "fmt"

func main() {  
    x := "text"
    xbytes := []byte(x)
    xbytes[0] = 'T'

    fmt.Println(string(xbytes)) //prints Text
}
