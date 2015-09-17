package main
import "math"
import "fmt"

type shape interface{

 perimeter() float64

}

type rect struct{
 h float64
 b float64

}
func (r rect) perimeter(h float64,b float64) float64{
return 2* (h + b)

}

type circle struct{

radius float64
}


func (c circle) perimeter(radius float64) float64{
return 2 * math.Pi * radius

}

func main(){
var height,breadth,rad float64

fmt.Println("Enter the height of rectangle")
fmt.Scanf("%f\n",&height)
fmt.Println("Enter the Breadth of rectangle")
fmt.Scanf("%f\n",&breadth)
fmt.Println("Enter radius of circle")
fmt.Scanf("%f\n",&rad)
r := rect{h: height,b: breadth}
c := circle{radius: rad}

fmt.Println("The perimeter of the rectangle is: ",r.perimeter(r.h,r.b))
fmt.Println("The perimeter of the circle is: ",c.perimeter(c.radius))



}