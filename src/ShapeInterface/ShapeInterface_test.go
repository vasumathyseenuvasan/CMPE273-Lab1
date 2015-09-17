package main
import "testing"

type testValues struct{
  rect Rectangle
  circle Circle
  rectanglePerimeter float64
  circlePerimeter float64
}

var testSamples = []testValues{
  {Rectangle{2,3},Circle{5},10,31.41592653589793},
  {Rectangle{2.5,3.4},Circle{4.5},11.8,28.274333882308138},
  {Rectangle{2.5,-2},Circle{5.2},0.0,32.67256359733385},
  {Rectangle{4,3},Circle{-1},14,0.0},
}

func TestShapeInterface(t *testing.T){
for _,values:=range testSamples{
  rectPerimeter:=CalculatePerimeter(&values.rect)
  circlePerimeter:=CalculatePerimeter(&values.circle)
  if rectPerimeter!=values.rectanglePerimeter{
    t.Error("For length ",values.rect.length, "and breadth ",values.rect.length,
      "Expected perimeter of the rectangle:  ",values.rectanglePerimeter ,"Actual:", rectPerimeter)
  }
  if circlePerimeter!=values.circlePerimeter{
    t.Error("For length ",values.circle.radius, "Expected perimeter of the circle: ",values.circlePerimeter,"Actual:", circlePerimeter)
  }
}
}
