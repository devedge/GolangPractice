package main

import (
  "fmt"
  "image"
  "log"
  "github.com/disintegration/imaging"
)

// "encoding/hex"
// "image/color"

func main() {
  // Open test image
  src, err := imaging.Open("lena_512.png")
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }

  dhash(src, 8)

  // Crop the image to 350x350 px using the Center anchor
  //src = imaging.CropAnchor(src, 350, 350, imaging.Center)

  // Resize the cropped image to a width of 256px, but preserving aspect ratio
  //src = imaging.Resize(src, 256, 0, imaging.Lanczos)

  // Create a blurred version of the image
  //img1 := imaging.Blur(src, 2)

  // Create a grayscaled version of the image, and heighten the contrast & sharpness
  //img2 := imaging.Grayscale(src)
  //img2 = imaging.AdjustContrast(img2, 20)
  //img2 = imaging.Sharpen(img2, 2)

  // Create an inverted version of the image
  //img3 := imaging.Invert(src)

  // Create an embossed version of the image using a convolution filter
  //img4 := imaging.Convolve3x3(
  //  src,
  //  [9]float64{
  //    -1, -1,  0,
  //    -1,  1,  1,
  //     0,  1,  1,
  //  },
  //  nil,
  //)

  // Create a new image and paste the four produced images into it
  //dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
  //dst = imaging.Paste(dst, img1, image.Pt(0, 0))
  //dst = imaging.Paste(dst, img2, image.Pt(0, 256))
  //dst = imaging.Paste(dst, img3, image.Pt(256, 0))
  //dst = imaging.Paste(dst, img4, image.Pt(256, 256))

  // Save the resulting image using JPEG format
  //err = imaging.Save(dst, "example_out.jpg")
  //if err != nil {
  //  log.Fatalf("Save failed: %v", err)
  //}
}


func dhash(img image.Image, bitLen int) {
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
      // r = r / 257 // Not needed, </> comparison still works the same

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

  fmt.Println(sig)

  // Initialize the difference matrix that will be returned
  // DiffMatrix := make([][]int, bitLen)
  // for i := range DiffMatrix {
  //   DiffMatrix[i] = make([]int, bitLen)
  // }
  // DiffMatrix[x-1][y] = 1 // DiffMatrix is 8x8, not 9x8
  // fmt.Print("1")
  // row = append(row, 1)


  /*
  // var diff [hashlen][hashlen]int
  // diff := make([][]int, hashlen)
  var pixels [][]uint32
  bounds := res.Bounds()

  s := ""

  for y := 0; y < bounds.Max.Y; y++ {
    var row []uint32
    for x := 0; x < bounds.Max.X; x++ {
      r,_,_,_ := res.At(x,y).RGBA()
      row = append(row, r/257)
    }
    fmt.Println(row)
    pixels = append(pixels, row)
  }

  // extract the first row NOPE

  // fmt.Println(diff)

  // For each row
  for i := 0; i < 8; i++ {

    // For each column, except the last
    for j := 0; j < 9 - 1; j++ {
      if pixels[i][j] < pixels[i][j+1] {
        // fmt.Print("1")
        s += "1"
      } else {
        // fmt.Print("0")
        s += "0"
      }
      // fmt.Println(pixels[i][j], pixels[i][j+1])
      // fmt.Println(i,j)
    }
  }

  dst := []byte(s)
  encStr := hex.EncodeToString(dst)

  // for y := 0; y < 8; y++ {
  //   for x := 0; x < 9; x++ {
      // if pixels[x-1][y] < pixels[x][y] {
      //   // diff[x][y] = 1
      //   fmt.Print("1")
      // } else {
      //   // diff[x][y] = 0
        // fmt.Print("0")
      // }
  //   }
  // }
  fmt.Printf("%s\n", encStr)
  // fmt.Println(diff)
  */
}

// Generate a pixel array that can be iterated over
/*func grayscalePixelArray(img image.Image) [][]uint32 {
  var pixels [][]uint32

  // Determine the bounds
  bounds := img.Bounds()

  // For each row
  for x := 0; x < bounds.Max.X; x++ {
    var row []uint32

    // For each column
    for y := 0; y < bounds.Max.Y; y++ {
      // Since it's grayscale, r = g = b
      r,_,_,_ := img.At(x,y).RGBA()

      // Append this column to a 'row' array
      row = append(row, r / 257)
    }

    // append that 'row' array to the pixel array
    pixels = append(pixels, row)
  }
  return pixels
}*/

// do a 'rotate' Go function and a 'scale' one
