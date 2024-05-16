package main

import (
	"math"
)

// Vec2 represents a 2D vector.
type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Multiply(scalar float64) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

func (v Vec2) Divide(scalar float64) Vec2 {
	return Vec2{X: v.X / scalar, Y: v.Y / scalar}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) DotProduct(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Normalize() Vec2 {
	return v.Divide(v.Length())
}

// Vec3 represents a 3D vector.
type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) ToVec4() Vec4 {
	return Vec4{v.X, v.Y, v.Z, 1}
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Multiply(scalar float64) Vec3 {
	return Vec3{v.X * scalar, v.Y * scalar, v.Z * scalar}
}

func (v Vec3) Divide(scalar float64) Vec3 {
	return Vec3{v.X / scalar, v.Y / scalar, v.Z / scalar}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) CrossProduct(other Vec3) Vec3 {
	x := v.Y*other.Z - v.Z*other.Y
	y := v.Z*other.X - v.X*other.Z
	z := v.X*other.Y - v.Y*other.X
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) DotProduct(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Normalize() Vec3 {
	return v.Divide(v.Length())
}

func (v Vec3) RotateX(angle float64) Vec3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vec3{
		X: v.X,
		Y: v.Y*cos - v.Z*sin,
		Z: v.Y*sin + v.Z*cos,
	}
}

func (v Vec3) RotateY(angle float64) Vec3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vec3{
		X: v.X*cos + v.Z*sin,
		Y: v.Y,
		Z: -v.X*sin + v.Z*cos,
	}
}

func (v Vec3) RotateZ(angle float64) Vec3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vec3{
		X: v.X*cos - v.Y*sin,
		Y: v.X*sin + v.Y*cos,
		Z: v.Z,
	}
}

type Vec4 struct {
	X, Y, Z, W float64
}

func (v Vec4) ToVec3() Vec3 {
	return Vec3{v.X, v.Y, v.Z}
}

func (v Vec4) Add(other Vec4) Vec4 {
	return Vec4{v.X + other.X, v.Y + other.Y, v.Z + other.Z, v.W + other.W}
}

func (v Vec4) Sub(other Vec4) Vec4 {
	return Vec4{v.X - other.X, v.Y - other.Y, v.Z - other.Z, v.W - other.W}
}

func (v Vec4) Divide(scalar float64) Vec4 {
	return Vec4{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar, W: v.W / scalar}
}

func (v Vec4) Multiply(scalar float64) Vec4 {
	return Vec4{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar, W: v.W * scalar}
}

func (v Vec4) DotProduct(other Vec4) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z + v.W*other.W
}

func (v Vec4) Conjugate() Vec4 {
	return Vec4{X: -v.X, Y: -v.Y, Z: -v.Z, W: v.W}
}

func (v Vec4) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

func (v Vec4) Normalize() Vec4 {
	return v.Divide(v.Length())
}
