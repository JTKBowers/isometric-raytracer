package geometry

import "math"

type Vector struct {
  x,y,z float64
}

func MakeVector(x,y,z float64) Vector {
  return Vector{x,y,z}
}

func (v Vector) Add(o Vector) Vector{
  return Vector{v.x+o.x, v.y+o.y, v.z+o.z}
}

func (v Vector) Sub(o Vector) Vector{
  return Vector{v.x-o.x, v.y-o.y, v.z-o.z}
}

func (v Vector) Dot(o Vector) float64{
  return v.x*o.x + v.y*o.y + v.z*o.z
}

func (v Vector) Inv() Vector{
  return Vector{1/v.x, 1/v.y, 1/v.z}
}

func (v Vector) Mul(s float64) Vector{
  return Vector{s*v.x, s*v.y, s*v.z}
}

func (v Vector) MulV(o Vector) Vector{
  return Vector{v.x*o.x, v.y*o.y, v.z*o.z}
}

func (v Vector) Mag() float64{
  return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) SqMag() float64{
  return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector) Apply(f func(float64) float64) Vector {
  return Vector{f(v.x), f(v.y), f(v.z)}
}

func (v Vector) ToArray() [3]float64 {
  return [3]float64 {v.x,v.y,v.z}
}

func (v Vector) Normalise() Vector {
  mag := v.Mag()
  if mag == 0 {
    return v
  }
  return v.Mul(1.0/mag)
}
