package imagehash

import (
  "image"
  "github.com/disintegration/imaging"
)

// A wrapper function around the Open function from 'imaging',
// so the user of 'imagehash' doesn't need to import 2
// packages to use it.
func OpenImg(fp string) (image.Image, error) {
  file,err := imaging.Open(fp)
  return file,err
}
