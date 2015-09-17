package main
import (
  "testing")

type testValues struct{
  number int
  fibonacciSeries []int
}

var testSamples = []testValues{
  {3,[]int{0,1,1}},
  {5,[]int{0,1,1,2,3}},
  {0,[]int{0}},
}

func TestFibonacciSeries(test *testing.T){
for _, sampleValue:= range testSamples{
  fibonacciArrayActual:=make([]int, sampleValue.number)
  fibonacciArrayExpected:=make([]int, sampleValue.number)
  copy(fibonacciArrayActual,GenerateFibonacciSeries(sampleValue.number))
  copy(fibonacciArrayExpected, sampleValue.fibonacciSeries)
  for i, v := range fibonacciArrayExpected {
        if v != fibonacciArrayActual[i] {
            test.Error("For Input ",sampleValue.number,"Expected: ",fibonacciArrayExpected,"Actual:",fibonacciArrayActual)
        }
    }
}
}
