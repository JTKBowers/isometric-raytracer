package render

type Ray struct {
  o,d Vector
}

func (r Ray) GetDirectionVector() Vector{
  return r.d
}

func (r Ray) GetOriginVector() Vector{
  return r.o
}
