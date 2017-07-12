package main
// package imagehash

import (
  "os"
  "fmt"
  "log"
  "image"
  "errors"
  "encoding/hex"
  "github.com/disintegration/imaging"
)
// "strconv"

func main() {
  // Open test image
  src, err := imaging.Open(os.Args[1])
  if err != nil {
    log.Fatalf("Open failed: %v", err)
  }

  hlen := 8
  // input,_ := strconv.Atoi(os.Args[2])
  // if input {
  // } else {
  //   hlen = input
  // }

  // Hash the image both vertically and horizontally
  hash,err := dhash(src, hlen)
  if err != nil {
    log.Fatalf("Failed hash: %v", err)
  }

  // Print the hex representation of the byte array
  fmt.Println(hex.EncodeToString(hash))
}


// dhash vertical
// dhash horizontal
// dhash from image
// dhash from file

/**
 * Wrapper function that calculates both the horizontal and vertical
 * gradients, then returns them appended as <horizontal><vertical>
 *
 * @method  dhash
 * @param   {Image}  img      The image file to perform the hash on
 * @param   {int}    hashLen  The integer length of the scaled-down image
 * @return  {[]byte}          The calculated byte array
 * @return  {error}           Any errors
 */
func dhash(img image.Image, hashLen int) ([]byte, error) {
  // Grayscale the image. Do this first for performance.
  imgGray := imaging.Grayscale(img)

  // Calculate both horizontal and vertical gradients
  horiz, err1 := horizontalGradient(imgGray, hashLen)
  vert, err2 := verticalGradient(imgGray, hashLen)

  if err1 != nil {
    return nil, err1
  } else if err2 != nil {
    return nil, err2
  }

  // Return the horizontal hash with the vertical one appended
  return append(horiz, vert...), nil
}


/**
 * Wrapper function around the horizontalGradient function,
 * so the Grayscale only ever runs once
 *
 * @method  dhashHorizontal
 * @param   {Image}  img      The image file to perform the hash on
 * @param   {int}    hashLen  The integer length of the scaled-down image
 * @return  {[]byte}          The calculated byte array
 * @return  {error}           Any errors
 */
func dhashHorizontal(img image.Image, hashLen int) ([]byte, error) {
  // Grayscale the image. Do this first for performance.
  imgGray := imaging.Grayscale(img)

  // Calculate the horizontal gradient
  horiz, err1 := horizontalGradient(imgGray, hashLen)
  if err1 != nil { return nil, err1 }

  return horiz, nil
}


/**
 * Wrapper function around the verticalGradient function,
 * so the Grayscale only ever runs once
 *
 * @method dhashVertical
 * @param   {Image}  img      The image file to perform the hash on
 * @param   {int}    hashLen  The integer length of the scaled-down image
 * @return  {[]byte}          The calculated byte array
 * @return  {error}           Any errors
 */
func dhashVertical(img image.Image, hashLen int) ([]byte, error) {
  // Grayscale the image. Do this first for performance.
  imgGray := imaging.Grayscale(img)

  // Calculate the horizontal gradient
  horiz, err1 := verticalGradient(imgGray, hashLen)
  if err1 != nil { return nil, err1 }

  return horiz, nil
}


/**
 * This function calculates the horizontal gradient difference of
 * an image.
 *
 * @method horizontalGradient
 * @param   {Image}  img      The image file to perform the hash on
 * @param   {int}    hashLen  The integer length of the scaled-down image
 * @return  {[]byte}          The calculated byte array
 * @return  {error}           Any errors
 */
func horizontalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen + 1, hashLen

  // Downscale the image by 'hashLen' amount for a horizonal diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := NewAppendBit(hashLen * hashLen)
  if err != nil { return nil, err }

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
          bitArray.appendBit(1) // if it's smaller, append '1'
        } else {
          bitArray.appendBit(0) // else append '0'
        }
      }
      prev = r // Set this current pixel value as the previous one
    }
  }
  return bitArray.getArray(), nil
}


/**
 * This function calculates the vertical gradient difference of
 * an image.
 *
 * @method verticalGradient
 * @param   {Image}  img      The image file to perform the hash on
 * @param   {int}    hashLen  The integer length of the scaled-down image
 * @return  {[]byte}          The calculated byte array
 * @return  {error}           Any errors
 */
func verticalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen, hashLen + 1

  // Downscale the image by 'hashLen' amount for a horizonal diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := NewAppendBit(hashLen * hashLen)
  if err != nil { return nil, err }

  var prev uint32 // Variable to store the previous pixel value

  // Calculate the horizonal gradient difference
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current row, then
      // compare the gradient difference from the previous one
      if y > 0 {
        if prev < r {
          bitArray.appendBit(1) // if it's smaller, append '1'
        } else {
          bitArray.appendBit(0) // else append '0'
        }
      }
      prev = r // Set this current pixel value as the previous one
    }
  }
  return bitArray.getArray(), nil
}


// Struct to simplify appending bits to a byte array,
// from left to right
type AppendBit struct {
  byteArray []byte
  max int
  mask0 byte
  mask1 byte
  arrayIdx int
  bitIdx uint
}

/**
 * Constructor for the AppendBit struct. It must be initialized
 * with a non-zero int that is a multiple of 8.
 * Usage:
 *    bitArray := NewAppendBit(64)
 *
 * @method NewAppendBit
 * @param  {int}        numBits The number of bits, as an int
 * @return {AppendBit}  The AppendBit struct
 * @return {error}      If there is an error
 */
func NewAppendBit(numBits int) (*AppendBit, error) {
  // If numBits is invalid
  if (numBits == 0) || (numBits % 8 != 0) {
    return nil, errors.New("'numBits' must be a non-zero multiple of 8")
  }

  return &AppendBit{
    byteArray: make([]byte, numBits / 8),
    max: numBits / 8,
    mask0: 0x00,
    mask1: 0x01,
    arrayIdx: 0,
    bitIdx: 7,
  }, nil
}


/**
 * Append a bit to the byte array from left to right
 * @method appendBit
 * @param  {int} bit  Append a one with '1', and zero with '0'
 */
func (ab *AppendBit) appendBit(bit int) error {
  if ab.arrayIdx == ab.max {
    return errors.New("Cannot contine to append to a full byteArray")
  }

  // Shift the 'mask' (bit of 1 or 0) by the proper amount to
  // fill the byte up from left to right.
  switch bit {
    case 0:
      ab.byteArray[ab.arrayIdx] |= ab.mask0 << ab.bitIdx
    case 1:
      ab.byteArray[ab.arrayIdx] |= ab.mask1 << ab.bitIdx
  }

  if ab.bitIdx > 0 {
    // Decrement the index into the current byte so the next
    // bit to be set will be on the right.
    ab.bitIdx--
  } else {
    // The last bit in the current byte has been set, so increment
    // the index into the byte array, and reset the bit index.
    ab.arrayIdx++
    ab.bitIdx = 7
  }

  return nil
}


/**
 * Returns the current byte array
 * @method getArray
 * @return {[]byte]}  The byte array
 */
func (ab AppendBit) getArray() []byte {
  return ab.byteArray
}
