package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
)

const FLOOR_SIZE float32 = 1000.0
const FLOOR_TEXTURES_PER_FLOOR float32 = 0.25

type World interface {
	Init()
	Terminate()
	GetStartPosition() mgl32.Vec3
	RemoveEnemy(enemy *Enemy)
	GetEnemies() []*Enemy
}

type World1 struct {
	enemy  Enemy
	sword  gohome.Entity3D
	hammer gohome.Entity3D
	floor  *PhysicsPlaneEntity
}

func (this *World1) Init() {
	this.enemy.Init([3]float32{0.0, ENEMY_HEIGHT, -5.0})

	this.sword.InitName("Sword_Sword_Cylinder.003")
	this.sword.Transform.Position[2] = -5.0
	this.sword.Transform.Position[0] = 3.0
	this.sword.Transform.Position[1] = 5.0
	this.sword.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})

	this.hammer.InitName("Hammer")
	this.hammer.Transform.Position[2] = -5.0
	this.hammer.Transform.Position[0] = 6.0
	this.hammer.Transform.Position[1] = 5.0
	this.hammer.Transform.Rotation = mgl32.QuatRotate(mgl32.DegToRad(90.0), mgl32.Vec3{1.0, 0.0, 0.0})

	this.floor = PhysicsEntityCon.NewPlane([2]float32{FLOOR_SIZE, FLOOR_SIZE}, FLOOR_TEXTURES_PER_FLOOR*FLOOR_SIZE)
	floorMaterial := this.floor.Model3D.GetMeshIndex(0).GetMaterial()
	floorMaterial.SetTextures("FloorTexture", "", "")
	floorMaterial.Shinyness = 0.3
	gohome.ResourceMgr.GetTexture("FloorTexture").SetWrapping(gohome.WRAPPING_REPEAT)

	AddCollectable(&Meat{}, [3]float32{9.0, 5.0, -5.0})
	AddCollectable(&Coin{}, [3]float32{12.0, 5.0, -5.0})

	this.addObjectsToGame()

}

func (this *World1) addObjectsToGame() {
	gohome.RenderMgr.AddObject(&this.sword)
	gohome.RenderMgr.AddObject(&this.hammer)
	PhysicsEntityCon.NewCube([3]float32{0.0, 100.0, -3.0}, [3]float32{10.0, 10.0, 2.0}, 8.0)
	PhysicsEntityCon.NewCube([3]float32{0.0, 200.0, -3.0}, [3]float32{10.0, 10.0, 2.0}, 8.0)

	PhysicsEntityCon.NewCube([3]float32{5.0, 5.0, -5.0}, [3]float32{30.0, 10.0, 30.0}, 0.0)

	gohome.UpdateMgr.AddObject(this)
}

func (this *World1) removeObjectsFromGame() {
	gohome.RenderMgr.RemoveObject(&this.sword)
	gohome.RenderMgr.RemoveObject(&this.hammer)
	gohome.RenderMgr.RemoveObject(this.floor)

	gohome.UpdateMgr.RemoveObject(this)
}

func (this *World1) Update(delta_time float32) {
	if gohome.InputMgr.JustPressed(gohome.KeyJ) {
		PhysicsEntityCon.NewCube([3]float32{0.0, 100.0, -3.0}, [3]float32{10.0, 10.0, 2.0}, 8.0)
		PhysicsEntityCon.NewCube([3]float32{0.0, 200.0, -3.0}, [3]float32{10.0, 10.0, 2.0}, 8.0)
	}
}

func (this *World1) GetEnemies() []*Enemy {
	return []*Enemy{&this.enemy}
}

func (this *World1) RemoveEnemy(enemy *Enemy) {
	gohome.RenderMgr.RemoveObject(enemy)
	gohome.UpdateMgr.RemoveObject(enemy)
	enemy.Terminate()
}

func (this *World1) Terminate() {
	this.removeObjectsFromGame()
	this.enemy.Terminate()
	this.sword.Terminate()
	this.hammer.Terminate()
}

func (this *World1) GetStartPosition() mgl32.Vec3 {
	return [3]float32{0.0, 0.0, 0.0}
}
