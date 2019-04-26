// https://play.golang.org/p/yGcW55n0BTo

package main

import "fmt"

type data struct {  
    num int
    key *string
    items map[string]bool
}

func (this *data) pmethod() {  
    this.num = 7
}

func (this data) vmethod() {  
    this.num = 8
    *this.key = "v.key"
    this.items["vmethod"] = true
}

func main() {  
    key := "key.1"
    d := data{1,&key,make(map[string]bool)}

    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
    //prints num=1 key=key.1 items=map[]

    d.pmethod()
    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items) 
    //prints num=7 key=key.1 items=map[]

    d.vmethod()
    fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
    //prints num=7 key=v.key items=map[vmethod:true]
}
