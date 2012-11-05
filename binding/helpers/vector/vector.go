package vector

import "math"

func InvSqrt(x float64) float64 {
	return 1 / math.Sqrt(x)
}

type Vector struct {
	X, Y float64
}

func New(x, y float64) Vector {
	return Vector{x, y}
}

func (v Vector) Neg() Vector {
	return New(-v.X, -v.Y)
}

func (v Vector) Sub(v2 Vector) Vector {
	return New(v.X-v2.X, v.Y-v2.Y)
}

func (v Vector) Add(v2 Vector) Vector {
	return New(v.X+v2.X, v.Y+v2.Y)
}

func (v *Vector) SubEqual(v2 Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) AddEqual(v2 Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vector) Eq(v2 Vector) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vector) Div(scalar float64) Vector {
	return New(v.X/scalar, v.Y/scalar)
}

func (v Vector) Mul(scalar float64) Vector {
	return New(v.X*scalar, v.Y*scalar)
}

func (v *Vector) MulEqual(scalar float64) *Vector {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Angle(arg ...interface{}) float64 {
	if len(arg) == 1 {
		if vec, ok := arg[0].(Vector); ok {
			v.Normalize()
			vec.Normalize()

			return math.Acos(v.Dot(vec))
		}
	} else {
		return math.Atan2(v.Y, v.X)
	}

	return 0.0
}

func (v *Vector) Clamp(max float64) {
	if v.Len() > max {
		v.Normalize()
		v.X *= max
		v.Y *= max
	}
}

func (v *Vector) Normalize() *Vector {
	rc := InvSqrt(v.Dot(*v))
	v.X *= rc
	v.Y *= rc

	return v
}

func (v *Vector) Rotate(a float64) *Vector {
	var vec Vector

	vec.X = v.X*math.Cos(a) - v.Y*math.Sin(a)
	vec.Y = v.X*math.Sin(a) + v.Y*math.Cos(a)

	v.X, v.Y = vec.X, vec.Y

	return v
}

func VectorAngle(v Vector, u Vector) float64 {
	return v.Angle(u)
}

func VectorDot(v Vector, u Vector) float64 {
	return v.Dot(u)
}
