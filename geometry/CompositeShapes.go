package geometry

import "math"

type BinaryTreeNode struct {
  leftChild, rightChild Object
}

func MakeTreeNode(leftChild, rightChild Object) BinaryTreeNode{
  return BinaryTreeNode{leftChild, rightChild}
}

func (n BinaryTreeNode) Collides(r Ray) bool{
  return n.leftChild.Collides(r) || n.rightChild.Collides(r)
}

func (n BinaryTreeNode) Collision(r Ray) (depth float64, obj Object, mat Material){
  depth = 1e99
  if n.leftChild.Collides(r) {
    depth, obj, mat = n.leftChild.Collision(r)
  }
  if n.rightChild.Collides(r) {
    rDepth, rObj, rColour := n.rightChild.Collision(r)
    if depth >= rDepth{
      depth, obj, mat = rDepth, rObj, rColour
    }
  }
  return
}

func (n BinaryTreeNode) Contains(point Vector) bool {
  return n.leftChild.Contains(point) || n.rightChild.Contains(point)
}

func (n BinaryTreeNode) Normal(r Ray, depth float64) Vector {
  //Laziness and inefficiency! TODO: improve!
  //However, it should hardly ever get called.
  // if the collision point is closer to the left child's collision point than the right's, return its normal.
  var lDepth, rDepth float64
  if n.leftChild.Collides(r) {
    lDepth, _, _ = n.leftChild.Collision(r)
  }
  if n.rightChild.Collides(r) {
    rDepth, _, _ = n.rightChild.Collision(r)
  }
  if math.Abs(lDepth - depth) <= math.Abs(rDepth - depth){
    return n.leftChild.Normal(r, depth)
  }
  return n.rightChild.Normal(r, depth)
}

// A composite shape that only renders when both its component shapes do.
// Can be used to create eg a sphere section
type Intersection struct {
  BinaryTreeNode
}

func MakeIntersection(object, clipShape Object) Intersection{
  return Intersection{BinaryTreeNode{object, clipShape}}
}

func (n Intersection) Collides(r Ray) bool{
  return n.leftChild.Collides(r) && n.rightChild.Collides(r)
}

func (n Intersection) Collision(r Ray) (depth float64, obj Object, mat Material){
  depth = 1e99
  if n.leftChild.Collides(r) {
    lDepth, lObj, lMaterial := n.leftChild.Collision(r)
    lpt := r.Point(lDepth)

    if n.rightChild.Collides(r) {
      obj = lObj
      rDepth, rObj, _ := n.rightChild.Collision(r)
      rpt := r.Point(rDepth)

      //if clipping occurs - ie the object's virtual surface is closer than the clipping surface
      if lDepth <= rDepth && n.leftChild.Contains(rpt){
        depth, mat = rDepth, lMaterial // Left child specifies the material
        obj = rObj
        return
      }
      //if no clipping occurs
      if lDepth > rDepth && n.rightChild.Contains(lpt){
        depth, mat = lDepth, lMaterial // Left child specifies the material
        return
      }
    }
  }
  return
}

func (n Intersection) Contains(point Vector) bool {
  return n.leftChild.Contains(point) && n.rightChild.Contains(point)
}
