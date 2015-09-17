package main

import "fmt"
import "time"

func Sleepfunc(x int) {

 <-time.After(time.Millisecond * time.Duration(x))
   fmt.Println("Slept for ",x," milliseconds")
    
          
}

func main() {
  var t int
  fmt.Println("Enter the time for sleep")
  fmt.Scanf("%d\n",&t)
  Sleepfunc(t)
    
}
