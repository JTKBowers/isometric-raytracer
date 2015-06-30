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
  plane := render.YAxisAlignedPlane{0}
  cube := render.MakeCuboid(render.MakeVector(0,1,0),render.MakeVector(0.5, 0.5, 0.5))
  //cube2 := render.MakeCuboid(render.MakeVector(1,0,0),render.MakeVector(0.5, 0.5, 0.5))
  scene := render.MakeTreeNode(cube, plane)
  img := render.RenderImage(640, 480, scene)
  f, err := os.Create("out.png")
  if err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  png.Encode(f, img)
}
