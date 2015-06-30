package render

import (
	"image"
	"image/color"
	"math"
	//"fmt"
  "math/rand"
)

const sinθ float64 = 0.4472135954999579 //1.0/math.Sqrt(5)
const sinθ2 float64 = 1.0/5.0 //sinθ*sinθ
func MakeViewRay(u,v,k float64) Ray {
	//var d1 float64 = math.Sqrt(1/(2+4*sinθ2))
	// = math.Sqrt(5/14)
	const d1 float64 = 0.5976143046671968
	//var d2 float64 = -2*sinθ*d1
	const d2 float64 = -2*sinθ*d1
	// = -math.Sqrt(2/7)
	d := Vector{d1,d2,d1}


	bu := Vector{1.0/math.Sqrt(2), 0, -1.0/math.Sqrt(2)}
	var bv1 float64 = math.Sqrt(1.0/7.0)//math.Sqrt(1/(2 + 1/sinθ2))
	var bv2 float64 = math.Sqrt(5.0/7.0)//bv1/sinθ
	bv := Vector{bv1, bv2, bv1}

	//fmt.Println(bu.Dot(d), bv.Dot(d), bu.Dot(bv))

	o := (bu.Mul(u)) .Add( bv.Mul(v) ) .Add(d.Mul(k))

	return Ray{o,d}
}
func RenderImage(width, height uint32, scene object) *image.RGBA {
	distance := 1.0
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
  b := img.Bounds()
  for y := b.Min.Y; y <= b.Max.Y; y++ {
   for x := b.Min.X; x <= b.Max.X; x++ {
		x_range := b.Max.X - b.Min.X
		y_range := b.Max.Y - b.Min.Y

		fxr := float64(x_range)
		fyr := float64(y_range)

		const subsamples int = 4

		colour := Vector{0,0,0}
		for i:= 0; i < subsamples; i++ {
			u  := (fxr/fyr)*(float64(x-x_range/2) + 0.5*rand.Float64() - 0.25)/fxr
			v  := -(float64(y-y_range/2) + 0.5*rand.Float64() - 0.25)/fyr

			r := MakeViewRay(u, v, distance)
			//r := MakeViewRay(math.Abs(distance)*u, math.Abs(distance)*v, distance)

			if scene.collides(r) {
				depth := scene.collisionDepth(r)

				cpt := r.o.Add(r.d.Mul(depth))
				//fmt.Printf("%f %f %f\n",cpt.x, cpt.y, cpt.z)
				//colour = colour.Add(Vector{255,255,255})
				colour = colour.Add(cpt.Mul(1275))
	    	//img.Set(x, y, color.RGBA{uint8(math.Abs(100*cpt.x)), uint8(math.Abs(100*cpt.y)), uint8(math.Abs(100*cpt.z)), 255})
	    	///img.Set(x, y, color.RGBA{255-uint8(depth), 255-uint8(depth), 255-uint8(depth), 255})
				//img.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				//img.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
		colour = colour.Apply(func(x float64) float64 {
			return math.Abs(x)/float64(subsamples)
		})
		if colour.y != 0{
			//fmt.Println(colour.y)
		}
		img.Set(x, y, color.RGBA{uint8(colour.x), uint8(colour.y), uint8(colour.z), 255})
   }
  }
	return img
}
