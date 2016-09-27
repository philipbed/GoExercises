package main

import "fmt"

func fibRec(n int) int {
  if (n == 0 || n == 1) {
    return n
  }else{
    return fibRec(n-1) + fibRec(n-2)
  }
}

func fibIter(n int) int {
  var i, n1, n2 int = 0, 0, 1
  var copy int
  for (i < n) {
    copy = n1
    n1 = n2
    n2 += copy
    i += 1
  }
  return n1
}

func main(){
  n := 23
  fmt.Println(fibRec(n))
  fmt.Println(fibIter(n))
}
