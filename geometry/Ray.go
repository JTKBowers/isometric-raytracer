package geometry

type Ray struct {
  o,d Vector
}

func (r Ray) GetDirectionVector() Vector{
  return r.d
}

func (r Ray) GetOriginVector() Vector{
  return r.o
}

func MakeRay(o,d Vector) Ray{
  return Ray{o,d}
}

func (r Ray) Point(depth float64) Vector {
  return r.o.Add(r.d.Mul(depth))
}
