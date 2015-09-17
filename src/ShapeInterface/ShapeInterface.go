package main
import(
  "fmt"
  "math"
  "validator")


type Shape interface{
  perimeter() float64
}

type Rectangle struct{
  length, breadth float64
}

type Circle struct{
  radius float64
}

func (rect *Rectangle) perimeter() float64{
  if rect.length <=0.0 || rect.breadth <=0.0{
    return 0.0
  }
  return 2*(rect.length + rect.breadth)
}

func (circle *Circle) perimeter() float64{
  if circle.radius <=0.0{
    return 0.0
  }
  return 2*math.Pi*circle.radius
}

func CalculatePerimeter(shapes Shape) float64{
  return  shapes.perimeter()
}

func main(){
  var input string
  var isValid bool
  var l,b,r float64
  fmt.Print("Enter the length of the rectangle: ")
  fmt.Scanf("%s\n",&input)
  isValid,l = validatorname.FloatValidator(input)
  if !isValid{
  return}
  fmt.Print("Enter the breadth of the rectangle: ")
  fmt.Scanf("%s\n",&input)
  isValid,b = validatorname.FloatValidator(input)
  if !isValid{
  return}
  fmt.Print("Enter the radius of the circle: ")
  fmt.Scanf("%s\n",&input)
  isValid,r = validatorname.FloatValidator(input)
  if !isValid{
  return}

  rectangle:=Rectangle{length:l,breadth:b}
  circle:=Circle{radius:r}
  rectanglePerimeter:=CalculatePerimeter(&rectangle)
  circlePerimeter:=CalculatePerimeter(&circle)

  fmt.Println("Perimeter of the rectangle: ",rectanglePerimeter)
  fmt.Println("Perimeter of the circle: ",circlePerimeter)
}
