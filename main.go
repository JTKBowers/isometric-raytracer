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
  cube1 := geometry.MakeCuboid(geometry.MakeVector(-2,0,0),geometry.MakeVector(0.5, 0.5, 0.5))
  sphere := geometry.MakeSphere(geometry.MakeVector(2,0,0), 2.0)
  sphere2 := geometry.MakeSphere(geometry.MakeVector(-2,2,0), 0.5)
  //cube2 := geometry.MakeCuboid(geometry.MakeVector(1,0,0),geometry.MakeVector(0.5, 0.5, 0.5))
  n := geometry.MakeTreeNode(sphere2, cube1)
  scene := geometry.MakeTreeNode(plane, n)
  scene = geometry.MakeTreeNode(sphere, scene)
  //cube := geometry.MakeCuboid(geometry.MakeVector(-2,0,0),geometry.MakeVector(0.5, 0.5, 0.5))
  //scene := geometry.MakeTreeNode(sphere, cube)

  img := render.RenderImage(1366, 768, scene)
  f, err := os.Create("out.png")
  if err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  png.Encode(f, img)
}
