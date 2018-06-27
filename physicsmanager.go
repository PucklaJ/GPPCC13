package main

import (
	"github.com/tbogdala/cubez"
	m "github.com/tbogdala/cubez/math"
)

type PhysicsManager struct {
	cubes   []*cubez.CollisionCube
	spheres []*cubez.CollisionSphere
	planes  []*cubez.CollisionPlane
}

func (this *PhysicsManager) NewCube(pos [3]float32, size [3]float32, mass float32) *cubez.CollisionCube {
	var cubeMass m.Real = m.Real(mass)
	var cubeInertia m.Matrix3
	cubeCollider := cubez.NewCollisionCube(nil, m.Vector3{m.Real(size[0] / 2.0), m.Real(size[1] / 2.0), m.Real(size[2] / 2.0)})
	cubeCollider.Body.Position = m.Vector3{m.Real(pos[0]), m.Real(pos[1]), m.Real(pos[2])}
	if cubeMass > 0.0 {
		cubeCollider.Body.SetMass(cubeMass)
	} else {
		cubeCollider.Body.SetInfiniteMass()
	}
	cubeInertia.SetBlockInertiaTensor(&cubeCollider.HalfSize, cubeMass)
	cubeCollider.Body.SetInertiaTensor(&cubeInertia)
	cubeCollider.Body.CalculateDerivedData()
	cubeCollider.CalculateDerivedData()
	this.cubes = append(this.cubes, cubeCollider)
	return cubeCollider
}

func (this *PhysicsManager) NewSphere(pos [3]float32, radius float32, mass float32) *cubez.CollisionSphere {
	var sphereMass m.Real = m.Real(mass)
	var sphereInertia m.Matrix3
	sphereCollider := cubez.NewCollisionSphere(nil, m.Real(radius))
	sphereCollider.Body.Position = m.Vector3{m.Real(pos[0]), m.Real(pos[1]), m.Real(pos[2])}
	if sphereMass > 0.0 {
		sphereCollider.Body.SetMass(sphereMass)
	} else {
		sphereCollider.Body.SetInfiniteMass()
	}
	var coeff m.Real = 0.4 * sphereMass * m.Real(radius) * m.Real(radius)
	sphereInertia.SetInertiaTensorCoeffs(coeff, coeff, coeff, 0.0, 0.0, 0.0)
	sphereCollider.Body.SetInertiaTensor(&sphereInertia)
	sphereCollider.Body.CalculateDerivedData()
	sphereCollider.CalculateDerivedData()
	this.spheres = append(this.spheres, sphereCollider)
	return sphereCollider
}

func (this *PhysicsManager) NewPlane() *cubez.CollisionPlane {
	plane := cubez.NewCollisionPlane(m.Vector3{0.0, 1.0, 0.0}, 0.0)
	this.planes = append(this.planes, plane)
	return plane
}

func (this *PhysicsManager) updateObjects(delta_time float32) {
	for _, cube := range this.cubes {
		if cube.Body.HasFiniteMass() {
			cube.Body.Integrate(m.Real(delta_time))
			cube.CalculateDerivedData()
		}
	}

	for _, sphere := range this.spheres {
		sphere.Body.Integrate(m.Real(delta_time))
		sphere.CalculateDerivedData()
	}
}

func (this *PhysicsManager) checkCollisions(delta_time float32) {
	foundCollisions := false
	foundCollisions1 := false
	var contacts []*cubez.Contact

	for _, cube := range this.cubes {
		for _, otherCube := range this.cubes {
			if cube != otherCube && cube.Body.HasFiniteMass() {
				if foundCollisions1, contacts = cubez.CheckForCollisions(cube, otherCube, contacts); foundCollisions1 {
					foundCollisions = true
				}
			}
		}
		for _, sphere := range this.spheres {
			if foundCollisions1, contacts = cubez.CheckForCollisions(cube, sphere, contacts); foundCollisions1 {
				foundCollisions = true
			}
		}
		for _, plane := range this.planes {
			if foundCollisions1, contacts = cubez.CheckForCollisions(cube, plane, contacts); foundCollisions1 {
				foundCollisions = true
			}
		}
	}

	for _, sphere := range this.spheres {
		for _, otherSphere := range this.spheres {
			if sphere != otherSphere {
				if foundCollisions1, contacts = cubez.CheckForCollisions(sphere, otherSphere, contacts); foundCollisions1 {
					foundCollisions = true
				}
			}
		}
		for _, plane := range this.planes {
			if foundCollisions1, contacts = cubez.CheckForCollisions(sphere, plane, contacts); foundCollisions1 {
				foundCollisions = true
			}
		}
	}
	if foundCollisions {
		cubez.ResolveContacts(len(contacts)*8, contacts, m.Real(delta_time))
	}
}

func (this *PhysicsManager) Update(delta_time float32) {
	this.updateObjects(delta_time)
	this.checkCollisions(delta_time)
}

var PhysicsMgr PhysicsManager
