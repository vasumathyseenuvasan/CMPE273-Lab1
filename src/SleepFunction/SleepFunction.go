package main
import (
  "time"
  "fmt"
  "validator"
  "bufio"
  "os")

func main(){
  var (
  isValid bool
  input string
  message string)

  fmt.Println("Sleep function using Time.After")
  fmt.Println("Enter the time (seconds) to sleep")
  fmt.Scanf("%s",&input)
  isValid,secondsDelay := validatorname.IntValidator(input)
  if !isValid{
  return}
  fmt.Println("Enter the message to be printed")
  in:= bufio.NewReader(os.Stdin)
  message,_ =in.ReadString('\n')
  fmt.Println("Enter the number of times your want the message to be printed")
  fmt.Scanf("%s",&input)
  isValid,number := validatorname.IntValidator(input)
  if !isValid{
  return}
  CallSleep(number,secondsDelay,message)
}

func CallSleep(number int, secondsDelay int, message string){
  for i:=0;i<number;i++{
  Sleeping(secondsDelay)
  fmt.Print(message)
  }
}

//Sleep Function using time.After
func Sleeping(secondsDelay int){
<-time.After(time.Duration(secondsDelay)*time.Second)
}
