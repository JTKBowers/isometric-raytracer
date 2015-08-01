package geometry

type Material struct{
  BaseColour Vector
  CTransmission, CReflectance, CBase float64
  RefractiveIndex float64
}

func SolidColourDiffuseMaterial(colour Vector) Material{
  return Material{colour, 0, 0, 1, 0 }
}

func HalfMirrorMaterial(colour Vector) Material{
  return Material{colour, 0, 0.5, 0.5, 0 }
}
