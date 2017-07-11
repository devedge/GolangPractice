package main

import (
  "fmt"
  "errors"
  "encoding/hex"
)

func main()  {
  byteres := make([]byte, 16)
  var mask byte = 0x01
  byteres[0] = 0x80
  byteres[0] |= mask << 2
  byteres[2] = mask
  fmt.Println(byteres)
  fmt.Println(hex.EncodeToString(byteres))

  bitArray,_ := NewAppendBit(64)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(0)
  bitArray.appendBit(0)
  bitArray.appendBit(1)
  bitArray.appendBit(1)
  bitArray.appendBit(1)

  fmt.Println(hex.EncodeToString(bitArray.getArray()))
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
