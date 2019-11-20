package main

import (
  "fmt"
  "github.com/ckbball/api-gin"
)

func main() {
  fmt.Println("False : False")
  fmt.Println(ValidateParams("string", "description"))

  fmt.Println("False : True")
  fmt.Println(ValidateParams("string", "desc"))

  fmt.Println("True : False")
  fmt.Println(ValidateParams("reads", "description"))

  fmt.Println("True : True")
  fmt.Println(ValidateParams("reads", "asc"))
}
