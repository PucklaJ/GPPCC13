package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	LIFE_BAR_REGEN_SPEED float32 = 0.5
)

type LifeBar struct {
	background gohome.Sprite2D
	foreground gohome.Sprite2D

	desiredLife float32
	currentLife float32
}

func (this *LifeBar) PreloadResources() {
	gohome.ResourceMgr.PreloadTexture("LifeBar","lifebar.png")
}

func (this *LifeBar) Init() {
	if gohome.ResourceMgr.GetTexture("LifeBar") != nil {
		gohome.ResourceMgr.GetTexture("LifeBar").SetFiltering(gohome.FILTERING_NEAREST)
	}

	this.background.Init("LifeBar")
	this.foreground.Init("LifeBar")

	this.background.TextureRegion.Max[0] = 270
	this.background.TextureRegion.Max[1] = 35

	this.foreground.TextureRegion.Min[1] = 36
	this.foreground.TextureRegion.Max[0] = 228
	this.foreground.TextureRegion.Max[1] = 51

	this.background.Transform.Size = [2]float32{270.0,35.0}
	this.foreground.Transform.Size = [2]float32{228.0,15.0}

	this.background.Transform.Scale = [2]float32{2.0,2.0}
	this.foreground.Transform.Scale = [2]float32{2.0,2.0}

	this.foreground.Transform.Position = this.background.Transform.Position.Add(mgl32.Vec2{21.0,12.0}.Mul(this.background.Transform.Scale[0])).Sub([2]float32{1.0,0.0})

	this.background.Depth = 0
	this.foreground.Depth = 1

	gohome.RenderMgr.AddObject(&this.background)
	gohome.RenderMgr.AddObject(&this.foreground)
	gohome.UpdateMgr.AddObject(this)
}

func (this *LifeBar) Update(delta_time float32) {
	if this.desiredLife < this.currentLife {
		this.currentLife = mgl32.Clamp(this.currentLife-LIFE_BAR_REGEN_SPEED*delta_time,this.desiredLife,1.0)
	} else {
		this.currentLife = mgl32.Clamp(this.currentLife+LIFE_BAR_REGEN_SPEED*delta_time,0.0,this.desiredLife)
	}

	this.foreground.TextureRegion.Max[0] = 0.0 + 228.0*this.currentLife
	this.foreground.Transform.Size[0] = 228.0*this.currentLife+2.0/this.foreground.Transform.Scale[0]
}

func (this *LifeBar) SetLife(value float32) {
	this.desiredLife = value
}

func (this *LifeBar) Terminate() {
	this.background.Terminate()
	this.foreground.Terminate()
}
