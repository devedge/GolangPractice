package main

import (
  "fmt"
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

  bitArray,_ := NewBitAppendArray(64)
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

  fmt.Println(hex.EncodeToString(bitArray.returnArray()))

  // byteStruct := BitAppendArray{}
  // byteStruct.init(64)
  // fmt.Print(byteStruct.lengthStruct())
}

type BitAppendArray struct {
  arr []byte
  mask0 byte
  mask1 byte
  arrIdx int
  bitIdx uint
}

// Constructor for the BitAppendArray struct
//
// usage:
// bitArray := NewBitAppendArray(64)
func NewBitAppendArray(numBits int) (*BitAppendArray, error) {
  // return error here if numBits is not good

  return &BitAppendArray{
    arr: make([]byte, numBits / 8),
    mask0: 0x00,
    mask1: 0x01,
    arrIdx: 0,
    bitIdx: 7,
  }, nil
}

// Append a bit to the byte array from left to right
func (ds *BitAppendArray) appendBit(bit int) {
  switch bit {
    case 0:
      ds.arr[ds.arrIdx] |= ds.mask0 << ds.bitIdx
    case 1:
      ds.arr[ds.arrIdx] |= ds.mask1 << ds.bitIdx
  }

  if ds.bitIdx > 0 {
    ds.bitIdx--
  } else {
    // the last bit in the current byte has been set, so
    ds.arrIdx++
    ds.bitIdx = 7
    // do something once ds.arrIdx has reached it's max
  }
}

func (ds *BitAppendArray) returnArray() []byte {
  return ds.arr
}
