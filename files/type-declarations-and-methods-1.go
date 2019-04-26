// https://play.golang.org/p/RVUMJ5yxrHE

package main

import "sync"

type myMutex sync.Mutex

func main() {  
    var mtx myMutex
    mtx.Lock() //error
    mtx.Unlock() //error  
}
