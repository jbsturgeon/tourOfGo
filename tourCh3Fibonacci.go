//The fibonacci series is:  x(n+2)=x(n+1)+x(n),
//where x(1) = 1 and x(0) = 0
package main

import "fmt"

//fibonacci is a funtion that returns
//a function that returns an int
func fibonacci() func() int64 {
  fmt.Println("fib called")
  a := int64(1)
  b := int64(0)

  return func() int64 {
  	fmt.Println("fib recur called")
    a, b = b, a+b
    return a
  }
}

func main() {
  f := fibonacci()

  for i:=0; i<10; i++ {
    fmt.Println(f())
  }
}

