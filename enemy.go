package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tbogdala/cubez"
	"strconv"
)

const (
	ENEMY_HEIGHT float32 = 5.2 * 0.75
)

var enemyModels [40]*gohome.Model3D

func preloadEnemyModels() {
	for i := 1; i <= 20; i++ {
		var zeros string = ""
		var numZeros int
		if i < 10 {
			numZeros = 5
		} else {
			numZeros = 4
		}
		for j := 0; j < numZeros; j++ {
			zeros += "0"
		}
		gohome.ResourceMgr.PreloadLevel("SirKnochen"+strconv.Itoa(i), "SirKnochen/SirKnochen_"+zeros+strconv.Itoa(i)+".obj", true)
	}
}

func initEnemyModels() {
	for i := 1; i <= 20; i++ {
		enemyModels[i-1] = gohome.ResourceMgr.GetLevel("SirKnochen" + strconv.Itoa(i)).GetModel("SirKnochen_Cube.001")
		enemyModels[39-(i-1)] = gohome.ResourceMgr.GetLevel("SirKnochen" + strconv.Itoa(i)).GetModel("SirKnochen_Cube.001")
	}
}

type Enemy struct {
	gohome.Entity3D

	currentModel  int
	modelTime     float32
	CollisionCube *cubez.CollisionCube
}

func (this *Enemy) Init(position mgl32.Vec3) {
	this.InitModel(gohome.ResourceMgr.GetLevel("SirKnochen9").GetModel("SirKnochen_Cube.001"))
	this.Transform.Position = position
	this.Transform.Scale = mgl32.Vec3{1.0, 1.0, 1.0}.Mul(0.75)

	gohome.UpdateMgr.AddObject(this)
	gohome.RenderMgr.AddObject(this)

	// this.currentModel = 19-1%20
	this.currentModel = 0
	this.modelTime = 0.0

	aabb := this.Model3D.AABB
	width := aabb.Max.X() - aabb.Min.X()
	height := aabb.Max.Y() - aabb.Min.Y()
	depth := aabb.Max.Z() - aabb.Min.Z()
	this.CollisionCube = PhysicsMgr.NewCube(position, [3]float32{width, height, depth}, 8.0)
}

func (this *Enemy) Update(delta_time float32) {
	this.modelTime += delta_time

	if this.modelTime >= 0.025 {
		this.modelTime = 0.0
		// this.currentModel = 19-(this.currentModel+1)%20
		this.currentModel++
		if this.currentModel == 40 {
			this.currentModel = 0
		}
	}

	this.Model3D = enemyModels[this.currentModel]
	this.Transform.Position[0] = float32(this.CollisionCube.Body.Position[0])
	this.Transform.Position[1] = float32(this.CollisionCube.Body.Position[1])
	this.Transform.Position[2] = float32(this.CollisionCube.Body.Position[2])
}

func (this *Enemy) GetHurtBox() gohome.AxisAlignedBoundingBox {
	aabb := this.Model3D.AABB

	this.Transform.CalculateTransformMatrix(&gohome.RenderMgr, -1)

	aabb.Min = gohome.Mat4MulVec3(this.Transform.GetTransformMatrix(), aabb.Min)
	aabb.Max = gohome.Mat4MulVec3(this.Transform.GetTransformMatrix(), aabb.Max)

	return aabb
}

func (this *Enemy) Damage(damage float32) {
	world.RemoveEnemy(this)
}

func (this *Enemy) Terminate() {
	gohome.UpdateMgr.RemoveObject(this)
	gohome.RenderMgr.RemoveObject(this)
}
