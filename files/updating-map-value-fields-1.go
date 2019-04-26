// https://play.golang.org/p/yeIBO1tIyqd

package main

type data struct {  
    name string
}

func main() {  
    m := map[string]data {"x":{"one"}}
    m["x"].name = "two" //error
}
