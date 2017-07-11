package main

import (
  "fmt"
  "log"
  "./imagehash.go"
  "github.com/disintegration/imaging"
)

func main() {
  // Open test image
  src, err := imaging.Open("lena_512.png")
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }

  fmt.Println(imagehash.Dhash(src,8))
}
