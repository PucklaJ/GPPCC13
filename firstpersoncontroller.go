package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	CAM_SENSITIVITY float32 = 0.5
	MOVE_SPEED float32 = 10.0
)

type FirstPersonController struct {
	camera gohome.Camera3D
	Velocity mgl32.Vec3
}

func (this *FirstPersonController) Init() {
	this.camera.Init()
	gohome.RenderMgr.SetCamera3D(&this.camera,0)
	gohome.UpdateMgr.AddObject(this)
}

func (this *FirstPersonController) Update(delta_time float32) {
	if gohome.InputMgr.JustPressed(gohome.KeyM) {
		gohome.Framew.CursorDisable()
	} else if gohome.InputMgr.JustPressed(gohome.KeyEscape) {
		gohome.Framew.CurserShow()
	}

	pitch := float32(gohome.InputMgr.Mouse.DPos[1]) * CAM_SENSITIVITY
	yaw := float32(gohome.InputMgr.Mouse.DPos[0]) * CAM_SENSITIVITY

	this.camera.AddRotation([2]float32{-pitch, -yaw})

	var forward mgl32.Vec3
	var left mgl32.Vec3
	var right mgl32.Vec3

	forward = [3]float32{
		this.camera.LookDirection.X(),
		0.0,
		this.camera.LookDirection.Z(),
	}

	forward = forward.Normalize()

	right = gohome.Mat4MulVec3(mgl32.Rotate3DY(mgl32.DegToRad(-90.0)).Mat4(),forward)
	left = right.Mul(-1.0)

	if gohome.InputMgr.IsPressed(gohome.KeyW) {
		this.Velocity = forward.Mul(MOVE_SPEED)
	} else if gohome.InputMgr.IsPressed(gohome.KeyA) {
		this.Velocity = left.Mul(MOVE_SPEED)
	} else if gohome.InputMgr.IsPressed(gohome.KeyD) {
		this.Velocity = right.Mul(MOVE_SPEED)
	} else if gohome.InputMgr.IsPressed(gohome.KeyS) {
		this.Velocity = forward.Mul(-MOVE_SPEED)
	} else {
		this.Velocity = mgl32.Vec3{0.0,0.0}
	}

	this.camera.Position = this.camera.Position.Add(this.Velocity.Mul(delta_time))

	this.camera.Position[1] = PLAYER_HEIGHT
}

func (this *FirstPersonController) GetPosition() mgl32.Vec3 {
	return this.camera.Position
}

func (this *FirstPersonController) SetPosition(position mgl32.Vec3) {
	this.camera.Position = position
}
