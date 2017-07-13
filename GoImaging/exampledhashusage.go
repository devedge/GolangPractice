package main

import (
  "fmt"
  "encoding/hex"
  "./imagehash"
)

func main() {
  // Open test image
  src,_ := imagehash.OpenImg("lena_512.png")
  // if err != nil {
  //   log.Fatalf("Open failed: %v", err)
  // }

  hash,_ := imagehash.Dhash(src, 8)

  fmt.Println(hex.EncodeToString(hash))
}
