package main
// package imagehash

import (
  "fmt"
  "log"
  "image"
  "github.com/disintegration/imaging"
)

func main() {
  // Open test image
  src, err := imaging.Open("lena_512.png")
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }

  fmt.Println(dhash(src,8))
}


// Dhash vertical
// Dhash horizontal
// Dhash from image
// Dhash from file

func dhash(img image.Image, bitLen int) []byte {
  // Width and height of the scaled-down image
  width, height := bitLen + 1, bitLen

  // Grayscale the image. Do this first for performance.
  res := imaging.Grayscale(img)

  // Downscale the image by 'bitLen' amount for a horizonal diff.
  res = imaging.Resize(res, width, height, imaging.Lanczos)

  var sig []byte  // The byte array signature that will be returned
  var prev uint32 // Variable to store the previous pixel value

  // Calculate the horizonal gradient difference
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current row, then
      // compare the gradient difference from the previous one
      if x > 0 {
        if prev < r {
          sig = append(sig, 1) // if it's smaller, append '1'
        } else {
          sig = append(sig, 0) // else append '0'
        }
      }
      prev = r // Set this current pixel value as the previous one
    }
  }

  return sig
}


func byteArrayToHex(arr []byte) {

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
