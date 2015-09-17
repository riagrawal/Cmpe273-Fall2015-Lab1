package main

import "testing"
import "time"
import "fmt"
func TestSleepfunc(t *testing.T) {
   var checkduration time.Duration = 999 * time.Microsecond
   t0 := time.Now().UTC()
   Sleepfunc(10)
   t1 := time.Now().UTC()
   var duration1 time.Duration = t1.Sub(t0)
   fmt.Println("duration for time.After function is:",duration1)
   t2 := time.Now().UTC()
   time.Sleep(time.Millisecond * time.Duration(10))
   t3:= time.Now().UTC()
   var duration2 time.Duration = t3.Sub(t2)
   fmt.Println("duration for Sleep funtion is :",duration2)
  if duration2-duration1 >  checkduration  {
    t.Error("Expected to sleep for 10 milliseconds")
  }
}

