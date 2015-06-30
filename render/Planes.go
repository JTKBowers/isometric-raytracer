package render

type XAxisAlignedPlane struct {
	X float64
}

func (c XAxisAlignedPlane) collides(r Ray) bool{
  return true
}

func (c XAxisAlignedPlane) collision(r Ray) (float64, Vector){
  return (c.X - r.o.x)/r.d.x, Vector{255, 255, 255}
}


type YAxisAlignedPlane struct {
	Y float64
}

func (c YAxisAlignedPlane) collides(r Ray) bool{
  return true
}

func (c YAxisAlignedPlane) collision(r Ray) (float64, Vector){
  return (c.Y - r.o.y)/r.d.y, Vector{255, 255, 255}
}


type ZAxisAlignedPlane struct {
	Z float64
}

func (c ZAxisAlignedPlane) collides(r Ray) bool{
  return true
}

func (c ZAxisAlignedPlane) collision(r Ray) (float64, Vector){
  return (c.Z - r.o.z)/r.d.z, Vector{255, 255, 255}
}
