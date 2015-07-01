package geometry

type XAxisAlignedPlane struct {
	X float64
}

func (c XAxisAlignedPlane) Collides(r Ray) bool{
  return r.d.x != 0 || r.o.x == c.X
}

func (c XAxisAlignedPlane) Collision(r Ray) (float64, Vector){
  return (c.X - r.o.x)/r.d.x, Vector{255, 255, 255}
}

func (c XAxisAlignedPlane) Contains(point Vector) bool {
  return point.x == c.X
}
func (c XAxisAlignedPlane) Normal(r Ray, depth float64) Vector {
  return Vector{1,0,0}
}

type YAxisAlignedPlane struct {
	Y float64
}

func (c YAxisAlignedPlane) Collides(r Ray) bool{
  // A ray will only miss if it is parallel and not on the plane
  return r.d.y != 0 || r.o.y == c.Y
}

func (c YAxisAlignedPlane) Collision(r Ray) (float64, Vector){
  return (c.Y - r.o.y)/r.d.y, Vector{255, 255, 255}
}

func (c YAxisAlignedPlane) Contains(point Vector) bool {
  return point.y == c.Y
}
func (c YAxisAlignedPlane) Normal(r Ray, depth float64) Vector {
  return Vector{0,1,0}
}


type ZAxisAlignedPlane struct {
	Z float64
}

func (c ZAxisAlignedPlane) Collides(r Ray) bool{
  return r.d.z != 0 || r.o.z == c.Z
}

func (c ZAxisAlignedPlane) Collision(r Ray) (float64, Vector){
  return (c.Z - r.o.z)/r.d.z, Vector{255, 255, 255}
}

func (c ZAxisAlignedPlane) Contains(point Vector) bool {
  return point.z == c.Z
}
func (c ZAxisAlignedPlane) Normal(r Ray, depth float64) Vector {
  return Vector{0,1,0}
}
