package geometry

import "math"

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

func (b Cuboid) Collides(r Ray) bool{
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

func (b Cuboid) MaterialAt(r Ray, depth float64) Material{
  // cpt := r.Point(depth)
  // col := cpt.Sub(b.centrePos).MulV(b.halfExtents.Inv()).Apply(func(x float64) float64 {
  //    return 255*math.Abs(x)
  // })
  // return HalfMirrorMaterial(col)
  //return HalfMirrorMaterial(Vector{255,255,255})
  return SolidColourDiffuseMaterial(Vector{255,0,0})
}

func (b Cuboid) Collision(r Ray) (float64, Object, Material){
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

  return tmin, b, b.MaterialAt(r, tmin)
}

func (b Cuboid) Contains(point Vector) bool {
  min := b.min.ToArray()
  max := b.max.ToArray()
  p := point.ToArray()
  for i := 0; i < 3; i++ {
    if p[i] < min[i] || p[i] > max[i] {
      return false
    }
  }
  return true
}

func (b Cuboid) Normal(r Ray, depth float64) Vector {
  cpt := r.Point(depth) //find the collision point
  ncpt := cpt.Sub(b.centrePos) //then translate it to box coords

  //next have an array of potential normals and find the closest
  normals := [6] Vector{
    Vector{1,0,0},
    Vector{-1,0,0},
    Vector{0,1,0},
    Vector{0,-1,0},
    Vector{0,0,1},
    Vector{0,0,-1},
  }
  i, closest := -1, 1.5
  for j := 0; j < 6; j++ {
    dot := ncpt.Dot(normals[j])
    if dot < closest {
      i, closest = j, dot
    }
  }
  return normals[i]
}

type Sphere struct {
  radius float64
  centre Vector
}

func MakeSphere(c Vector, r float64) Sphere{
  return Sphere{r,c}
}

func (s Sphere) Collides(r Ray) bool{
  o := r.o.Sub(s.centre)
  p := o.Sub(r.d.Mul(o.Dot(r.d)))
  return s.radius*s.radius >= p.SqMag()
}

func (s Sphere) Collision(r Ray) (float64, Object, Material){
  if s.Collides(r) {
    o := r.o.Sub(s.centre)
    //solve quadratic equation for t:
    a := r.d.SqMag()
    b := 2*o.Dot(r.d)
    c := o.SqMag() - s.radius*s.radius

    discr := b*b - 4*a*c
    t1 := (-b + math.Sqrt(discr))/(2*a)
    t2 := (-b - math.Sqrt(discr))/(2*a)
    return math.Min(t1,t2), s, HalfMirrorMaterial(Vector{0,0,255})
  }
  return 1e99, s, SolidColourDiffuseMaterial(Vector{0,0,0})
}

func (s Sphere) Contains(point Vector) bool {
  return s.centre.Sub(point).Mag() <= s.radius
}

func (s Sphere) Normal(r Ray, depth float64) Vector {
  cpt := r.Point(depth) //find the collision point
  return s.centre.Sub(cpt).Mul(1.0/s.radius) //then translate it to sphere coords, and normalise by the sphere radius
}
