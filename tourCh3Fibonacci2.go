//The fibonacci series is:  x(n+2)=x(n+1)+x(n),
//where x(1) = 1 and x(0) = 0
package main

import "fmt"

//fibonacci is a funtion that returns
//a function that returns an int
func fibonacci(x int64) int64 {
  if x == 0 {
    return 0
  } else if x == 1 {
    return 1
  } else {
    return fibonacci(x-1) + fibonacci(x-2)
  }
}

func main() {
  for i:=0; i<50; i++ {
    fmt.Println("i = ", i, " ", fibonacci(int64(i)))
  }
}

