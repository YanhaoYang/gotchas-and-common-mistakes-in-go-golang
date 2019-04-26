// https://play.golang.org/p/jMvfSGLnBhZ

package main

import "fmt"

func main() {  
    data := []int{1,2,3}
    for i,_ := range data {
        data[i] *= 10
    }

    fmt.Println("data:",data) //prints data: [10 20 30]
}
