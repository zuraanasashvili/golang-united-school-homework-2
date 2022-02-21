package square

// Define custom int type to hold sides number and update CalcSquare
// signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.
// Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type custType int

const SidesTriangle custType = 3
const SidesSquare custType = 4
const SidesCircle custType = 0

func CalcSquare(sideLen float64, sidesNum custType) float64 {
	if sidesNum == SidesTriangle {
		// TODO: needs to be worked on
		// operator ^ not defined on float64 (in our case sideLen)
		return (math.Sqrt(3) * sideLen * sideLen) / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} else {
		return 0
	}
}
