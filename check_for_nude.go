package main

import (
    "fmt"
    // go get github.com/koyachi/go-nude
    "github.com/koyachi/go-nude"
    "flag"
)

var (
  IMAGE string = ""
)

func main() {
  flag.StringVar(&IMAGE, "img", IMAGE, "файл с изображением")
  flag.Parse()

  if IMAGE != "" {
    isNude, err := nude.IsNude(IMAGE)
    if err != nil {
        // fmt.Println(err)
        fmt.Println("Error")
    } else {
      if isNude {
        fmt.Println("True")
      } else {
        fmt.Println("False")
      }
    }
  } else {
    fmt.Println("NoFile")
  }
}
