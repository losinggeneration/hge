package hge

import "math"

func InvSqrt(x float32) float32 {
	return float32(1 / math.Sqrt(float64(x)))
}

type Vector struct {
	X, Y float32
}

func NewVector(x, y float32) Vector {
	var v Vector
	v.X = x
	v.Y = y
	return v
}

func (v Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y)
}

func (v Vector) Subtract(v2 Vector) Vector {
	return NewVector(v.X-v2.X, v.Y-v2.Y)
}

func (v Vector) Add(v2 Vector) Vector {
	return NewVector(v.X+v2.X, v.Y+v2.Y)
}

func (v *Vector) SubtractEqual(v2 Vector) *Vector {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector) AddEqual(v2 Vector) *Vector {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v Vector) EQ(v2 Vector) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vector) NEQ(v2 Vector) bool {
	return v.X != v2.X && v.Y != v2.Y
}

func (v Vector) Divide(scalar float32) Vector {
	return NewVector(v.X/scalar, v.Y/scalar)
}

func (v Vector) Multiply(scalar float32) Vector {
	return NewVector(v.X*scalar, v.Y*scalar)
}

func (v *Vector) MultiplyEequal(scalar float32) *Vector {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v Vector) Dot(v2 Vector) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.Dot(v))))
}

func (v Vector) Angle(arg ...interface{}) float32 {
	if len(arg) == 1 {
		if vec, ok := arg[0].(Vector); ok {
			s := vec
			t := vec

			s.Normalize()
			t.Normalize()
			return float32(math.Acos(float64(s.Dot(t))))
		}
	} else {
		return float32(math.Atan2(float64(v.Y), float64(v.X)))
	}

	return 0.0
}

func (v *Vector) Clamp(max float32) {
	if v.Length() > max {
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

func (v *Vector) Rotate(a float32) *Vector {
	var vec Vector

	vec.X = v.X*float32(math.Cos(float64(a))) - v.Y*float32(math.Sin(float64(a)))
	vec.Y = v.X*float32(math.Sin(float64(a))) + v.Y*float32(math.Cos(float64(a)))

	v.X = vec.X
	v.Y = vec.Y

	return v
}

func VectorAngle(v Vector, u Vector) float32 {
	return v.Angle(u)
}

func VectorDot(v Vector, u Vector) float32 {
	return v.Dot(u)
}
