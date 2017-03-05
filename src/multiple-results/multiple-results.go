package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  var a, b string = "hello", "world"

  fmt.Println(a, b)
  a, b = swap(a, b)
  fmt.Println(a, b)
}
