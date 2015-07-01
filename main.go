package main

import (
 "os"
 "image"
 "image/png"
 "io"
 "fmt"
 "isometric-renderer/render"
 "isometric-renderer/geometry"
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
  plane := geometry.YAxisAlignedPlane{0}
  cube := geometry.MakeCuboid(geometry.MakeVector(0,1,0),geometry.MakeVector(0.5, 0.5, 0.5))
  sphere := geometry.MakeSphere(geometry.MakeVector(0,0,0), 1.0)
  //cube2 := geometry.MakeCuboid(geometry.MakeVector(1,0,0),geometry.MakeVector(0.5, 0.5, 0.5))
  n := geometry.MakeIntersection(sphere, cube)
  scene := geometry.MakeTreeNode(n, plane)
  img := render.RenderImage(640, 480, scene)
  f, err := os.Create("out.png")
  if err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  png.Encode(f, img)
}
