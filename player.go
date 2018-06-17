package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	PLAYER_HEIGHT float32 = 5.0
	PLAYER_WIDTH float32 = 4.0
	PLAYER_DEPTH float32 = 2.0
	PLAYER_MAX_HEALTH float32 = 20.0
)

type Player struct {
	lifeBar LifeBar
	weapon Weapon
	fpc FirstPersonController
	health float32
	AABB gohome.AxisAlignedBoundingBox
}

func (this *Player) PreloadResources() {
	gohome.ResourceMgr.PreloadLevel("Sword","Sword.obj",true)
	this.lifeBar.PreloadResources()
	this.weapon = &Sword{}
}

func (this *Player) Init() {
	this.lifeBar.Init()
	this.weapon.Init()
	this.fpc.Init()
	this.SetHealth(PLAYER_MAX_HEALTH/2.0)

	gohome.UpdateMgr.AddObject(this)

	this.AABB.Min = [3]float32{-PLAYER_WIDTH/2.0,0.0,-PLAYER_DEPTH/2.0}
	this.AABB.Max = [3]float32{PLAYER_WIDTH/2.0,PLAYER_HEIGHT,PLAYER_DEPTH/2.0}
}

func (this *Player) Update(delta_time float32) {
	if gohome.InputMgr.JustPressed(gohome.MouseButtonLeft) {
		this.weapon.AttackAnimation()
		enemies := world.GetEnemies()
		for _,enemy := range enemies {
			if enemy.GetHurtBox().Intersects(mgl32.Vec3{0.0,0.0,0.0},this.weapon.GetHitBox(this),mgl32.Vec3{0.0,0.0,0.0}) {
				enemy.Damage(this.weapon.GetDamage())
			}
		}

	} else if this.fpc.Velocity.Len() != 0.0 {
		this.weapon.WalkAnimation()
	} else {
		this.weapon.IdleAnimation()
	}

}

func (this *Player) GetPosition() mgl32.Vec3 {
	return this.fpc.GetPosition().Sub([3]float32{0.0,PLAYER_HEIGHT,0.0})
}

func (this *Player) SetPosition(position mgl32.Vec3) {
	this.fpc.SetPosition(position)
}

func (this *Player) AddHealth(value float32) {
	this.SetHealth(this.health+value)
}

func (this *Player) ReduceHealth(value float32) {
	this.AddHealth(-value)
}

func (this *Player) SetHealth(value float32) {
	this.health = mgl32.Clamp(value,0.0,PLAYER_MAX_HEALTH)
	this.lifeBar.SetLife(this.health/PLAYER_MAX_HEALTH)
}

func (this *Player) Terminate() {
	this.lifeBar.Terminate()
	this.weapon.Terminate()
}