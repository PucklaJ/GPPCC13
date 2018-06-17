package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/colornames"
)

type Weapon interface {
	Init()
	AttackAnimation()
	IdleAnimation()
	WalkAnimation()
	GetDamage() float32
	GetHitBox(player *Player) gohome.AxisAlignedBoundingBox
	Terminate()
}

func preloadWeaponResources() {
	gohome.ResourceMgr.PreloadLevel("Hammer","Hammer.obj",true)
	gohome.ResourceMgr.PreloadLevel("Sword","Sword.obj",true)
}

type Sword struct {
	gohome.Entity3D

	attackAnimation gohome.Tweenset
	idleAnimation gohome.Tweenset
	walkAnimation gohome.Tweenset
	currentAnimation *gohome.Tweenset

	wantIdleTime float32
	wantWalkTime float32
}

func (this *Sword) Init() {
	this.InitName("Sword_Sword_Cylinder.003")
	this.Transform.Position[2] = -1.0
	this.Transform.Position[1] = -1.0
	this.Transform.Position[0] = 1.0
	this.Transform.Rotation[0] = 90.0
	this.Transform.Rotation[2] = 90.0
	this.NotRelativeToCamera = 0

	gohome.RenderMgr.AddObject(this)

	this.attackAnimation = gohome.Tweenset{
		Tweens:[]gohome.Tween{
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{-0.7,-0.7,-1.0},Time:0.3,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{0.0,0.0,45.0},Time:0.3,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-1.0,-1.0},Time:0.1,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{90.0,0.0,90.0},Time:0.1,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
		},
		Loop: false,
	}

	this.idleAnimation = gohome.Tweenset{
		Tweens:[]gohome.Tween{
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{100.0,20.0,95.0},Time:0.75*2.0,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-0.8,-1.0},Time:0.75*2.0,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{80.0,-20.0,85.0},Time:1.5*2.0,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-1.2,-1.0},Time:1.5*2.0,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{90.0,0.0,90.0},Time:0.75*2.0,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-1.0,-1.0},Time:0.75*2.0,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
		},
		Loop: true,
	}

	this.walkAnimation = gohome.Tweenset{
		Tweens:[]gohome.Tween{
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{110.0,0.0,95.0},Time:0.75*0.5,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-0.8,-1.0},Time:0.75*0.5,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{70.0,0.0,85.0},Time:1.5*0.5,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-1.2,-1.0},Time:1.5*0.5,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
			&gohome.TweenRotation3D{Destination:mgl32.Vec3{90.0,0.0,90.0},Time:0.75*0.5,TweenType:gohome.TWEEN_TYPE_AFTER_PREVIOUS},
			&gohome.TweenPosition3D{Destination:mgl32.Vec3{1.0,-1.0,-1.0},Time:0.75*0.5,TweenType:gohome.TWEEN_TYPE_WITH_PREVIOUS},
		},
		Loop: true,
	}

	this.attackAnimation.SetParent(this)
	this.idleAnimation.SetParent(this)
	this.walkAnimation.SetParent(this)

	gohome.UpdateMgr.AddObject(&this.attackAnimation)
	gohome.UpdateMgr.AddObject(&this.idleAnimation)
	gohome.UpdateMgr.AddObject(&this.walkAnimation)

	this.attackAnimation.Stop()
	this.walkAnimation.Stop()
	this.idleAnimation.Stop()
}
func (this *Sword) AttackAnimation() {
	this.wantIdleTime = 0.0
	this.wantWalkTime = 0.0

	if this.currentAnimation != nil && this.currentAnimation != &this.attackAnimation {
		this.currentAnimation.Stop()
		this.currentAnimation = &this.attackAnimation

		this.attackAnimation.Start()
	}
}
func (this *Sword) IdleAnimation() {
	this.wantIdleTime += gohome.FPSLimit.DeltaTime
	this.wantWalkTime = 0.0

	if (this.currentAnimation != nil && this.currentAnimation == &this.walkAnimation && this.wantIdleTime > 0.25) || (this.currentAnimation == &this.attackAnimation && this.attackAnimation.Done()) || this.currentAnimation == nil {
		if this.currentAnimation != nil {
			this.currentAnimation.Stop()
		}
		this.currentAnimation = &this.idleAnimation

		this.idleAnimation.Start()
	}
}
func (this *Sword) WalkAnimation() {
	this.wantWalkTime += gohome.FPSLimit.DeltaTime
	this.wantIdleTime = 0.0

	if (this.currentAnimation != nil && this.currentAnimation == &this.idleAnimation && this.wantWalkTime > 0.25) || (this.currentAnimation == &this.attackAnimation && this.attackAnimation.Done()) || this.currentAnimation == nil  {
		if this.currentAnimation != nil {
			this.currentAnimation.Stop()
		}
		this.currentAnimation = &this.walkAnimation

		this.walkAnimation.Start()
	}
}

func (this *Sword) GetDamage() float32 {
	return 0.5
}

func (this *Sword) GetHitBox(player *Player) gohome.AxisAlignedBoundingBox {
	var aabb gohome.AxisAlignedBoundingBox

	player.fpc.camera.CalculateViewMatrix()

	aabb.Min = gohome.Mat4MulVec3(player.fpc.camera.GetInverseViewMatrix(),[3]float32{-1.0,-2.0,-3.0})
	aabb.Max = gohome.Mat4MulVec3(player.fpc.camera.GetInverseViewMatrix(),[3]float32{1.0,1.0,0.0})

	var debug gohome.AABBRenderer

	debug.Init(&aabb,nil,colornames.Blue)
	debug.Load()
	gohome.RenderMgr.AddObject(&debug)

	return aabb
}

func (this *Sword) Terminate() {
	this.Entity3D.Terminate()
}
