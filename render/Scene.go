package render

import (
  "math"
)

//An interface for anything that a ray can collide with - ie a scene, AABB, plane etc
type object interface {
  collides(Ray) bool
  collision(Ray) (float64, Vector)
}

type Scene struct {
}

func (s Scene) fromFile(filename string) {

}

func (s Scene) collides(r Ray) bool{
  return false
}

func (s Scene) collision(r Ray) float64{
  return 0.0
}

func (s Scene) colourAt(r Ray, depth float64) Vector{
  return Vector{0,0,0}
}

type Cuboid struct {
  min, max Vector
  centrePos, halfExtents Vector
}

func MakeCuboid(centrePos, halfExtents Vector) Cuboid {
  var c Cuboid

  //for now, centre it at the origin
  c.min = centrePos.Sub(halfExtents)
  c.max = centrePos.Add(halfExtents)

  c.centrePos = centrePos
  c.halfExtents = halfExtents
  return c
}

func (b Cuboid) collides(r Ray) bool{
  d_inv := r.d.Inv()
  tx1 := (b.min.x - r.o.x)*d_inv.x
  tx2 := (b.max.x - r.o.x)*d_inv.x

  tmin := math.Min(tx1, tx2)
  tmax := math.Max(tx1, tx2)

  ty1 := (b.min.y - r.o.y)*d_inv.y
  ty2 := (b.max.y - r.o.y)*d_inv.y

  tmin = math.Max(tmin, math.Min(ty1, ty2))
  tmax = math.Min(tmax, math.Max(ty1, ty2))

  tz1 := (b.min.z - r.o.z)*d_inv.z
  tz2 := (b.max.z - r.o.z)*d_inv.z

  tmin = math.Max(tmin, math.Min(tz1, tz2))
  tmax = math.Min(tmax, math.Max(tz1, tz2))

  return tmax >= tmin
}

func (b Cuboid) ColourAt(r Ray, depth float64) Vector{
  cpt := r.o.Add(r.d.Mul(depth))
  return cpt.Sub(b.centrePos).MulV(b.halfExtents.Inv()).Apply(func(x float64) float64 {
    return 255*math.Abs(x)
  })
}

func (b Cuboid) collision(r Ray) (float64, Vector){
  d_inv := r.d.Inv()
  tx1 := (b.min.x - r.o.x)*d_inv.x
  tx2 := (b.max.x - r.o.x)*d_inv.x

  tmin := math.Min(tx1, tx2)

  ty1 := (b.min.y - r.o.y)*d_inv.y
  ty2 := (b.max.y - r.o.y)*d_inv.y

  tmin = math.Max(tmin, math.Min(ty1, ty2))

  tz1 := (b.min.z - r.o.z)*d_inv.z
  tz2 := (b.max.z - r.o.z)*d_inv.z

  tmin = math.Max(tmin, math.Min(tz1, tz2))

  return tmin, b.ColourAt(r, tmin)
}

type BinaryTreeNode struct {
  leftChild, rightChild object
}

func MakeTreeNode(leftChild, rightChild object) BinaryTreeNode{
  return BinaryTreeNode{leftChild, rightChild}
}

func (n BinaryTreeNode) collides(r Ray) bool{
  return n.leftChild.collides(r) || n.rightChild.collides(r)
}

func (n BinaryTreeNode) collision(r Ray) (depth float64, colour Vector){
  depth = 1e99
  if n.leftChild.collides(r) {
    depth, colour = n.leftChild.collision(r)
  }
  if n.rightChild.collides(r) {
    rDepth, rColour := n.rightChild.collision(r)
    if depth >= rDepth{
      depth, colour = rDepth, rColour
    }
  }
  return
}
