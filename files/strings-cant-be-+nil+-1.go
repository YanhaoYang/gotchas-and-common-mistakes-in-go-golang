// https://play.golang.org/p/HOIlrhat8Wo

package main

func main() {  
    var x string = nil //error

    if x == nil { //error
        x = "default"
    }
}
