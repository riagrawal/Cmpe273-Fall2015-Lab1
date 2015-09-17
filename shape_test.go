package main

import "testing"
import "fmt"

func TestShapeRect(t *testing.T) {
  
  r:= rect{h:3,b:4}
  fmt.Println("Check for perimeter with height = 3 and width = 4")
  if r.perimeter(3,4) != 14 {
    t.Error("Expected 18, got ", r.perimeter(3,4))
  }
  fmt.Println("perimeter of Rectangle is 18 as expected ")
} 

func TestShapeCircle(t *testing.T) {
  c := circle{radius:2}
  fmt.Println("Check for perimeter with radius = 2")
  if c.perimeter(2) != 12.566370614359172   {
    t.Error("Expected 12.566370614359172, got ", c.perimeter(2))
  }
  fmt.Println("perimeter of circle is 12.566370614359172 as expected ")
}