package main

import "testing"
import "fmt"

func TestFib(t *testing.T) {
  var v int
  fmt.Println("Checking the 3rd value of fibonacci series")
  v = fib(3)
  if v != 2 {
    t.Error("Expected 1.5, got ", v)
  }
  fmt.Println("Third value is 2 as expected ")
}