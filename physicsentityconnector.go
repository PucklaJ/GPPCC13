package main

import (
	"github.com/PucklaMotzer09/gohomeengine/src/gohome"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/tbogdala/cubez"
	// m "github.com/tbogdala/cubez/math"
)

type PhysicsCubeEntity struct {
	gohome.Entity3D
	Collider *cubez.CollisionCube
}

type PhysicsSphereEntity struct {
	gohome.Entity3D
	Collider *cubez.CollisionSphere
}

type PhysicsPlaneEntity struct {
	gohome.Entity3D
	Collider *cubez.CollisionPlane
}

type PhysicsEntityConnector struct {
	cubes   []PhysicsCubeEntity
	spheres []PhysicsSphereEntity
	planes  []PhysicsPlaneEntity
}

func (this *PhysicsEntityConnector) NewCube(pos [3]float32, size [3]float32, mass float32) *PhysicsCubeEntity {
	physicsEntity := PhysicsCubeEntity{}
	physicsEntity.Entity3D.InitMesh(gohome.Box("PhysicsCubeEntityCube", size))
	physicsEntity.Collider = PhysicsMgr.NewCube(pos, size, mass)
	physicsEntity.Entity3D.Transform.Position = pos
	gohome.RenderMgr.AddObject(&physicsEntity)
	this.cubes = append(this.cubes, physicsEntity)
	return &this.cubes[len(this.cubes)-1]
}

func (this *PhysicsEntityConnector) NewSphere(pos [3]float32, radius float32, mass float32) *PhysicsSphereEntity {
	physicsEntity := PhysicsSphereEntity{}
	physicsEntity.Entity3D.InitMesh(gohome.Box("PhysicsSphereEntityCube", [3]float32{radius * 2, radius * 2, radius * 2}))
	physicsEntity.Collider = PhysicsMgr.NewSphere(pos, radius, mass)
	physicsEntity.Entity3D.Transform.Position = pos
	gohome.RenderMgr.AddObject(&physicsEntity)
	this.spheres = append(this.spheres, physicsEntity)
	return &this.spheres[len(this.spheres)-1]
}

func (this *PhysicsEntityConnector) NewPlane(size [2]float32, textures float32) *PhysicsPlaneEntity {
	physicsEntity := PhysicsPlaneEntity{}
	physicsEntity.Entity3D.InitMesh(gohome.Plane("PhysicsPlaneEntityPlane", size, textures))
	physicsEntity.Collider = PhysicsMgr.NewPlane()
	gohome.RenderMgr.AddObject(&physicsEntity)
	this.planes = append(this.planes, physicsEntity)
	return &this.planes[len(this.planes)-1]
}

func (this *PhysicsEntityConnector) Update(delta_time float32) {
	for _, cube := range this.cubes {
		pos := cube.Collider.Body.Position
		// pos.Sub(&cube.Collider.HalfSize)
		// pos.Add(&m.Vector3{0.0, cube.Collider.HalfSize[1], 0.0})
		orient := cube.Collider.Body.Orientation
		cube.Entity3D.Transform.Position = [3]float32{float32(pos[0]), float32(pos[1]), float32(pos[2])}
		cube.Entity3D.Transform.Rotation = mgl32.Quat{W: float32(orient[0]), V: [3]float32{float32(orient[1]), float32(orient[2]), float32(orient[3])}}
	}

	for _, sphere := range this.spheres {
		pos := sphere.Collider.Body.Position
		orient := sphere.Collider.Body.Orientation
		sphere.Entity3D.Transform.Position = [3]float32{float32(pos[0]), float32(pos[1]), float32(pos[2])}
		sphere.Entity3D.Transform.Rotation = mgl32.Quat{W: float32(orient[0]), V: [3]float32{float32(orient[1]), float32(orient[2]), float32(orient[3])}}
	}
}

var PhysicsEntityCon PhysicsEntityConnector
