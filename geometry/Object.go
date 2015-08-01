package geometry

//An interface for anything that a ray can collide with - ie a scene, AABB, plane etc
type Object interface {
  Collides(Ray) bool
  Collision(Ray) (float64, Object, Material)
  Normal(Ray, float64) Vector
  Contains(Vector) bool
}
