package main
// package imagehash

import (
  "fmt"
  "log"
  "image"
  "strings"
  "github.com/disintegration/imaging"
)

func main() {
  // Open test image
  src, err := imaging.Open("lena_512.png")
  //lena_512.png
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }
  dhash(src,8)

  // fmt.Println(dhash(src,8))
  // fmt.Println(hex.EncodeToString(dhash(src,8)))
  // byteArrayToHex(dhash(src,8))

  // fmt.Printf("%x\n", dhash(src,8))
}


// Dhash vertical
// Dhash horizontal
// Dhash from image
// Dhash from file

func dhash(img image.Image, bitLen int) {
  // Width and height of the scaled-down image
  width, height := bitLen + 1, bitLen

  // Grayscale the image. Do this first for performance.
  res := imaging.Grayscale(img)

  // Downscale the image by 'bitLen' amount for a horizonal diff.
  res = imaging.Resize(res, width, height, imaging.Lanczos)

  sig := make([]string, bitLen * bitLen) // The byte array signature that will be returned
  i := 0
  var prev uint32 // Variable to store the previous pixel value
  var test uint8

  // Calculate the horizonal gradient difference
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current row, then
      // compare the gradient difference from the previous one
      if x > 0 {
        if prev < r {
          sig[i] = "1" // if it's smaller, append '1'
        } else {
          sig[i] = "0" // else append '0'
        }
        i++ // increment the index into the 'sig' array
      }
      prev = r // Set this current pixel value as the previous one
    }
  }
  fmt.Println(sig)
  // return matchHex(sig)
}

func horizontalGradient(img image.Image, bitLen int, width height int)  {
  cbyte := make([]byte, bitLen)
  var mask byte = 0x01
  i := 0

  for y := 0; y < height; y++ {
    cbyte[i] = 0x00 // init current byte

    for x := 0; x < width; x++ {
      r,_,_,_ := res.At(x,y).RGBA()
      if x > 0 {

        // need to do an 'append action'
        if prev < r {
          sig[i] = "1"
          mask <<  8 - i
        } else {
          sig[i] = "0"
        }
        i++
      }
      prev = r
    }

    // add the latest byte to the byte array
  }
}

// func arrayToByte(arr []string)  {
//   len := len(arr)
//   if len % 4 != 0 { fmt.Println("err") }
//
//   res := make([]byte, len)
//   for i := 0; i < len; i++ {
//
//   }
// }



func matchHex(arr []string) string {
  out := ""
  for i := 0; i < len(arr) / 4; i++ {
    switch strings.Join(arr[i:i+4], "") {
      case "0000":
        out += "0"
      case "0001":
        out += "1"
      case "0010":
        out += "2"
      case "0011":
        out += "3"
      case "0100":
        out += "4"
      case "0101":
        out += "5"
      case "0110":
        out += "6"
      case "0111":
        out += "7"
      case "1000":
        out += "8"
      case "1001":
        out += "9"
      case "1010":
        out += "a"
      case "1011":
        out += "b"
      case "1100":
        out += "c"
      case "1101":
        out += "d"
      case "1110":
        out += "e"
      case "1111":
        out += "f"
    }
  }
  return out
}


// func binaryArrayToHex(arr []byte) {
  // def b2hx(arr):
  // ...     h = 0
  // ...     s = []
  // ...     for i, v in enumerate(arr):
  // ...             if v:
  // ...                     h += 2**(i % 8)
  // ...             if (i % 8) == 7:
  // ...                     s.append(hex(h)[2:].rjust(2, '0'))
  // ...                     h = 0
  // ...     return "".join(s)
// }
