// https://play.golang.org/p/jEJ4PIzSWTb

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    resp, err := http.Get("https://api.ipify.org?format=json")
    defer resp.Body.Close()//not ok
    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}
