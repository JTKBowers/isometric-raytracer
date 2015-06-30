package main

import (
 "os"
 "image"
 "image/png"
 "io"
 "fmt"
 "isometric-renderer/render"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
 img, _, err := image.Decode(r)
 if err != nil {
  return err
 }
 return png.Encode(w, img)
}

func main(){
  img := render.RenderImage(640, 480,render.MakeCuboid(render.MakeVector(0.1,0.1,0.1),render.MakeVector(0.1, 0.1, 0.1)))
  f, err := os.Create("out.png")
  if err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  png.Encode(f, img)
}
