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

func (n BinaryTreeNode) Collision(r Ray) (depth float64, colour Vector){
  depth = 1e99
  if n.leftChild.Collides(r) {
    depth, colour = n.leftChild.Collision(r)
  }
  if n.rightChild.Collides(r) {
    rDepth, rColour := n.rightChild.Collision(r)
    if depth >= rDepth{
      depth, colour = rDepth, rColour
    }
  }
  return
}

func (n BinaryTreeNode) Contains(point Vector) bool {
  return n.leftChild.Contains(point) || n.rightChild.Contains(point)
}

func (n BinaryTreeNode) Normal(r Ray, depth float64) Vector {
  //Laziness and inefficiency! TODO: improve!
  // if the collision point is closer to the left child's collision point than the right's, return its normal.
  var lDepth, rDepth float64
  if n.leftChild.Collides(r) {
    lDepth, _ = n.leftChild.Collision(r)
  }
  if n.rightChild.Collides(r) {
    rDepth, _ = n.rightChild.Collision(r)
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

func MakeIntersection(leftChild, rightChild Object) Intersection{
  return Intersection{BinaryTreeNode{leftChild, rightChild}}
}

func (n Intersection) Collides(r Ray) bool{
  return n.leftChild.Collides(r) && n.rightChild.Collides(r)
}

func (n Intersection) Collision(r Ray) (depth float64, colour Vector){
  depth = 1e99
  if n.leftChild.Collides(r) {
    lDepth, lColour := n.leftChild.Collision(r)
    lpt := r.Point(lDepth)

    if n.rightChild.Collides(r) {
      rDepth, _ := n.rightChild.Collision(r)
      rpt := r.Point(rDepth)

      if lDepth <= rDepth && n.leftChild.Contains(rpt){
        depth, colour = rDepth, lColour // Left child specifies the colour
        return
      }
      if lDepth > rDepth && n.rightChild.Contains(lpt){
        depth, colour = lDepth, lColour // Left child specifies the colour
        return
      }
    }
  }
  return
}

func (n Intersection) Contains(point Vector) bool {
  return n.leftChild.Contains(point) && n.rightChild.Contains(point)
}
