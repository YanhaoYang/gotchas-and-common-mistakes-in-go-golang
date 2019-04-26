// https://play.golang.org/p/En9Y1AqukYC

package main

func main() {  
    x := 2
    y := 4

    table := make([][]int,x)
    for i:= range table {
        table[i] = make([]int,y)
    }
}
