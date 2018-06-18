package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"golang.org/x/image/colornames"
)

var world World

type StartScene struct {
	player Player
	world World
}

func (this *StartScene) loadResources() {
	gohome.ResourceMgr.LoadModelsWithSameName = true

	preloadEnemyModels()
	preloadCollectablesResources()
	preloadWeaponResources()
	this.player.PreloadResources()
	gohome.ResourceMgr.PreloadTexture("FloorTexture","Floor1.png")

	gohome.ResourceMgr.LoadPreloadedResources()
	initEnemyModels()

	gohome.ResourceMgr.LoadModelsWithSameName = false
	gohome.RenderMgr.EnableBackBuffer = false
}

func (this *StartScene) initObjects() {
	this.player.Init()
	this.world = &World1{}
	world = this.world
	this.world.Init()
	this.player.SetPosition(this.world.GetStartPosition())
}

func (this *StartScene) Init() {
	gohome.InitDefaultValues()
	gohome.ErrorMgr.ErrorLevel = gohome.ERROR_LEVEL_WARNING

	this.loadResources()
	this.initObjects()

	gohome.LightMgr.SetAmbientLight(colornames.Gray,0)
	gohome.LightMgr.AddDirectionalLight(&gohome.DirectionalLight{
		Direction: [3]float32{1.0,-1.0,0.0},
		DiffuseColor: colornames.Khaki,
		SpecularColor: colornames.Lime,
		CastsShadows: 0,
	},0)

}

func (this *StartScene) Update(delta_time float32) {
	if gohome.InputMgr.JustPressed(gohome.KeyB) {
		gohome.RenderMgr.WireFrameMode = !gohome.RenderMgr.WireFrameMode
	} else if gohome.InputMgr.JustPressed(gohome.KeyF11) {
		if gohome.Framew.WindowIsFullscreen() {
			gohome.Render.SetNativeResolution(1280,720)
		} else {
			gohome.Render.SetNativeResolution(1920,1080)
		}
		gohome.Framew.WindowSetFullscreen(!gohome.Framew.WindowIsFullscreen())
	}

	CheckCollectables(&this.player)
}

func (this *StartScene) Terminate() {
	this.player.Terminate()
	this.world.Terminate()
}
