package oglabstr

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/chewxy/math32"
)

type Direction byte

const (
	Forward  Direction = iota
	Backward
	Right
	Left
	Up
	Down
)

type Camera struct {
	Position         mgl32.Vec3
	Front, Up, Right mgl32.Vec3
	WorldUp          mgl32.Vec3

	Yaw, Pitch float32

	MovementSpeed float32
	Sensitivity   float32
	Zoom          float32
}

func NewCamera(position mgl32.Vec3) *Camera {
	c := Camera{}

	c.Position = position
	c.Front = mgl32.Vec3{0.0, 0.0, -1.0}
	c.WorldUp = mgl32.Vec3{0.0, 1.0, 0.0}

	c.Yaw = -90.0
	c.Pitch = 0.0
	c.MovementSpeed = 2.5
	c.Sensitivity = 0.1
	c.Zoom = 45.0

	c.updateVectors()

	return &c
}

func (c *Camera) GetViewMat() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}

func (c *Camera) KeyboardUpdate(direction Direction, deltaTime float32) {
	velocity := c.MovementSpeed * deltaTime

	switch direction {
	case Forward:
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	case Backward:
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	case Right:
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	case Left:
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	case Up:
		c.Position = c.Position.Add(c.Up.Mul(velocity))
	case Down:
		c.Position = c.Position.Sub(c.Up.Mul(velocity))
	}
}

func (c *Camera) MouseUpdate(xoffset, yoffset float32) {
	xoffset *= c.Sensitivity
	yoffset *= c.Sensitivity

	c.Yaw += xoffset
	c.Pitch += yoffset

	if c.Pitch > 89.0 {
		c.Pitch = 89.0
	}
	if c.Pitch < -89.0 {
		c.Pitch = -89.0
	}

	c.updateVectors()
}

func (c *Camera) ScrollUpdate(yoffset float32) {
	if c.Zoom >= 1.0 && c.Zoom <= 45.0 {
		c.Zoom -= yoffset
	}
	if c.Zoom <= 1.0 {
		c.Zoom = 1.0
	}
	if c.Zoom >= 45.0 {
		c.Zoom = 45.0
	}
}

func (c *Camera) updateVectors() {
	c.Front = mgl32.Vec3{
		math32.Cos(mgl32.DegToRad(c.Yaw)) * math32.Cos(mgl32.DegToRad(c.Pitch)),
		math32.Sin(mgl32.DegToRad(c.Pitch)),
		math32.Sin(mgl32.DegToRad(c.Yaw)) * math32.Cos(mgl32.DegToRad(c.Pitch)),
	}.Normalize()

	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}
