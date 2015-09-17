package main
import (
  "fmt"
  "validator")

func main(){
var number int
var input string

fmt.Print("Enter a number to print Fibonacci series: ")
fmt.Scanf("%s",&input)

isValid,number := validatorname.IntValidator(input)
if !isValid{
return}
fibonacciArray:=make([]int, number)
fmt.Print("Fibonacci of ", number, " is ")
fmt.Println(FibonacciSeries(number))
fmt.Print("Fibonacci Series : ")
copy(fibonacciArray,GenerateFibonacciSeries(number))
for i:=0;i<number;i++{
fmt.Print(fibonacciArray[i]," ")
}
}

func GenerateFibonacciSeries(number int) []int{
  fibonacciSeriesArray := make([]int, number)
  for i:=0;i<number;i++{
  fibonacciSeriesArray[i]= FibonacciSeries(i)
  }
  return fibonacciSeriesArray
}

// Recursive function to generate Fibonacci Series
func FibonacciSeries(number int) int{
  if number == 0 || number ==1{
    return number
  }else{
    return FibonacciSeries(number-1)+FibonacciSeries(number-2)
  }
}
