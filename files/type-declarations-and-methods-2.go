// https://play.golang.org/p/jOJhwpIlWEB

package main

import "sync"

type myLocker struct {  
    sync.Mutex
}

func main() {  
    var lock myLocker
    lock.Lock() //ok
    lock.Unlock() //ok
}
