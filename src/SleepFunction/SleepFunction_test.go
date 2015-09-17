package main

import (
  "testing"
  "time")

type testValues struct{
  secondsDelay int
  expectedDelay int
}
var testSamples = []testValues{
  {3,3},
  {2,2}}


func TestSleepFunction(test *testing.T){
  for _,values:= range testSamples{
  start:=time.Now()
  Sleeping(values.secondsDelay)
  s:=time.Now().Sub(start)
  if int(s.Seconds())!= values.expectedDelay{
    test.Error("Expected delay :", values.expectedDelay, "Actual Delay :",int(s.Seconds()))
  }
}
}
