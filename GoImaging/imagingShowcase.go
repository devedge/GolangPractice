package main

import (
  "image"
  "image/color"
  "log"
  "github.com/disintegration/imaging"
)

func main() {
  // Open test image
  src, err := imaging.Open("lena_512.png")
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }

  // Crop the image to 350x350 px using the Center anchor
  src = imaging.CropAnchor(src, 350, 350, imaging.Center)

  // Resize the cropped image to a width of 256px, but preserving aspect ratio
  src = imaging.Resize(src, 256, 0, imaging.Lanczos)

  // Create a blurred version of the image
  img1 := imaging.Blur(src, 2)

  // Create a grayscaled version of the image, and heighten the contrast & sharpness
  img2 := imaging.Grayscale(src)
  img2 = imaging.AdjustContrast(img2, 20)
  img2 = imaging.Sharpen(img2, 2)

  // Create an inverted version of the image
  img3 := imaging.Invert(src)

  // Create an embossed version of the image using a convolution filter
  img4 := imaging.Convolve3x3(
    src,
    [9]float64{
      -1, -1,  0,
      -1,  1,  1,
       0,  1,  1,
    },
    nil,
  )

  // Create a new image and paste the four produced images into it
  dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})
  dst = imaging.Paste(dst, img1, image.Pt(0, 0))
  dst = imaging.Paste(dst, img2, image.Pt(0, 256))
  dst = imaging.Paste(dst, img3, image.Pt(256, 0))
  dst = imaging.Paste(dst, img4, image.Pt(256, 256))

  // Save the resulting image using JPEG format
  err = imaging.Save(dst, "example_out.jpg")
  if err != nil {
    log.Fatalf("Save failed: %v", err)
  }
}
