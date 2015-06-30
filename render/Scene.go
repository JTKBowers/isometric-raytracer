package render

import "math"

//An interface for anything that a ray can collide with - ie a scene, AABB, plane etc
type object interface {
  collides(Ray) bool
  collisionDepth(Ray) float64
}

type Scene struct {
}

func (s Scene) fromFile(filename string) {

}

func (s Scene) collides(r Ray) bool{
  return false
}

func (s Scene) collisionDepth(r Ray) float64{
  return 0.0
}

type Cuboid struct {
  min, max Vector
}

func MakeCuboid(centrePos, halfExtents Vector) Cuboid {
  var c Cuboid

  //for now, centre it at the origin
  c.min = centrePos.Sub(halfExtents)
  c.max = centrePos.Add(halfExtents)
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

func (b Cuboid) collisionDepth(r Ray) float64{
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

  return tmin
}

type XAxisAlignedPlane struct {
	X float64
}

func (c XAxisAlignedPlane) collides(r Ray) bool{
  return true
}

func (c XAxisAlignedPlane) collisionDepth(r Ray) float64{
  return (c.X - r.o.x)/r.d.x
}
