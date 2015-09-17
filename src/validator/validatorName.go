package validatorname
import ("strconv"
        "fmt")

// To validate if user enters float value
func FloatValidator(input string) (bool, float64){
  value,err := strconv.ParseFloat(input,64)
  if err != nil{
  fmt.Print("Please enter a valid input")
  return false, value
  }else{
    return true, value
  }
}

// To validate if user enters int value
func IntValidator(input string) (bool,int){
  value,err := strconv.Atoi(input)
  if err != nil{
  fmt.Print("Please enter a valid input")
  return false, value
  }else{
    return true, value
  }
}
