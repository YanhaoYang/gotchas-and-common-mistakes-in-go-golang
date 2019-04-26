// https://play.golang.org/p/X8l0A2t40Dr

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    tr := &http.Transport{DisableKeepAlives: true}
    client := &http.Client{Transport: tr}

    resp, err := client.Get("http://golang.org")
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp.StatusCode)

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(len(string(body)))
}
