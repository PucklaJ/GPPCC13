package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	MEAT_HEALTH_GAIN float32 = PLAYER_MAX_HEALTH/5.0
)

func preloadCollectablesResources() {
	gohome.ResourceMgr.PreloadLevel("Meat","Meat.obj",true)
	gohome.ResourceMgr.PreloadLevel("Coin","Coin.obj",true)
}

type Collectable interface {
	Init(position mgl32.Vec3)
	OnCollect(player *Player)
	IsColliding(player *Player) bool
	Terminate()
}

var collectables []Collectable

func AddCollectable(collectable Collectable, position mgl32.Vec3) {
	collectables = append(collectables,collectable)
	collectable.Init(position)
}

func RemoveCollectable(collectable Collectable) {
	for i:=0;i<len(collectables);i++ {
		if collectables[i] == collectable {
			collectable.Terminate()
			if i+1 <len(collectables) {
				collectables = append(collectables[:i],collectables[i+1:]...)
			} else {
				collectables = collectables[:i]
			}
			return
		}
	}
}

func CheckCollectables(player *Player) {
	for i:=0;i<len(collectables);i++ {
		if collectables[i].IsColliding(player) {
			collectables[i].OnCollect(player)
			RemoveCollectable(collectables[i])
			if len(collectables) == 0 {
				return
			}
			i--
		}
	}
}

type Meat struct {
	entity gohome.Entity3D
}

func (this *Meat) Init(position mgl32.Vec3) {
	this.entity.InitName("Meat")
	this.entity.Transform.Position = position
	gohome.RenderMgr.AddObject(&this.entity)
}

func (this *Meat) OnCollect(player *Player) {
	player.AddHealth(MEAT_HEALTH_GAIN)
}

func (this *Meat) IsColliding(player *Player) bool {
	return this.entity.Model3D.AABB.Intersects(this.entity.Transform.Position,player.AABB,player.GetPosition())
}

func (this *Meat) Terminate() {
	gohome.RenderMgr.RemoveObject(&this.entity)
}

type Coin struct {
	entity gohome.Entity3D
}

func (this *Coin) Init(position mgl32.Vec3) {
	this.entity.InitName("Coin")
	this.entity.Transform.Position = position
	this.entity.Transform.Rotation[0] = 90.0

	gohome.UpdateMgr.AddObject(this)
	gohome.RenderMgr.AddObject(&this.entity)
}

func (this *Coin) Update(delta_time float32) {
	this.entity.Transform.Rotation[2] += 30.0 * delta_time
}

func (this *Coin) OnCollect(player *Player) {
}

func (this *Coin) IsColliding(player *Player) bool {
	return this.entity.Model3D.AABB.Intersects(this.entity.Transform.Position,player.AABB,player.GetPosition())
}

func (this *Coin) Terminate() {
	gohome.UpdateMgr.RemoveObject(this)
	gohome.RenderMgr.RemoveObject(&this.entity)
}




