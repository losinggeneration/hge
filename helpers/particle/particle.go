package particle

import (
	"C"
	. "github.com/losinggeneration/hge-go/helpers/color"
	. "github.com/losinggeneration/hge-go/helpers/rect"
	. "github.com/losinggeneration/hge-go/helpers/sprite"
	. "github.com/losinggeneration/hge-go/helpers/vector"
	"github.com/losinggeneration/hge-go/hge"
	"math"
	"reflect"
	"unsafe"
)

const (
	hgeMAX_PARTICLES = 500
	hgeMAX_PSYSTEMS  = 100
)

type cast unsafe.Pointer

type particle struct {
	location Vector
	velocity Vector

	gravity, radialAccel, tangentialAccel float64
	spin, spinDelta                       float64
	size, sizeDelta                       float64
	color, colorDelta                     ColorRGB

	age, terminalAge float64
}

type ParticleSystemInfo struct {
	Sprite                         // texture + blend mode
	Emission                   int // particles per sec
	Lifetime, LifeMin, LifeMax float64
	Direction, Spread          float64

	Relative bool

	SpeedMin, SpeedMax                     float64
	GravityMin, GravityMax                 float64
	RadialAccelMin, RadialAccelMax         float64
	TangentialAccelMin, TangentialAccelMax float64
	SizeStart, SizeEnd, SizeVar            float64
	SpinStart, SpinEnd, SpinVar            float64

	ColorStart, ColorEnd ColorRGB

	ColorVar, AlphaVar float64
}

type ParticleSystem struct {
	Info ParticleSystemInfo

	updateSpeed, residue, age, emissionResidue float64
	prevLocation, location                     Vector

	tx, ty float64

	particlesAlive    int
	boundingBox       Rect
	updateBoundingBox bool
	particles         []particle
}

func NewParticleSystem(filename string, sprite Sprite, a ...interface{}) *ParticleSystem {
	ps := new(ParticleSystem)
	if len(a) == 1 {
		if fps, ok := a[0].(float64); ok {
			if fps != 0.0 {
				ps.updateSpeed = 1.0 / fps
			}
		}
	}

	ptr := hge.LoadBytes(filename)

	if ptr == nil {
		hge.Log("Particle file (%s) seems to be empty.", filename)
		return nil
	}

	// skip the first four bytes
	i := uintptr(4)

	// Ok, First we reflect the ParticleSystemInfo struct
	s := reflect.ValueOf(&ps.Info).Elem()

	// Then we loop through each element, skipping sprite for obvious reasons
	for j := 1; j < s.NumField(); j++ {
		// Then we get the field of the struct
		f := s.Field(j)

		// Here we examine the type
		switch f.Type().String() {
		case "float64":
			// Then we set the structure's field to the value at the current
			// byte(s)
			// We cast the value pointed to by the ptr[i] with unsafe.Pointer
			f.SetFloat(float64(*(*float32)(cast(&ptr[i]))))
			// Next we skip ahead based on the size of the data read
			i += unsafe.Sizeof(float32(0.0))

		case "int":
			f.SetInt(int64(*(*int32)(cast(&ptr[i]))))
			i += unsafe.Sizeof(int32(0))

		case "bool":
			f.SetBool(*(*bool)(cast(&ptr[i])))
			i += unsafe.Sizeof(bool(false))
			i += 3 // padding

		case "color.ColorRGB":
			for k := 0; k < 4; k++ {
				f.Field(k).SetFloat(float64(*(*float32)(cast(&ptr[i]))))
				i += unsafe.Sizeof(float32(0.0))
			}
		}
	}

	ps.Info.Sprite = sprite
	ps.age = -2.0

	ps.particles = make([]particle, hgeMAX_PARTICLES+1)

	return ps
}

func NewParticleSystemWithInfo(psi ParticleSystemInfo, a ...interface{}) *ParticleSystem {
	ps := new(ParticleSystem)
	if len(a) == 1 {
		if fps, ok := a[0].(float64); ok {
			if fps != 0.0 {
				ps.updateSpeed = 1.0 / fps
			}
		}
	}

	ps.Info = psi
	ps.age = -2.0

	ps.particles = make([]particle, hgeMAX_PARTICLES)

	return ps
}

func (ps *ParticleSystem) Render() {
	col := ps.Info.Sprite.Color()

	for i := 0; i < ps.particlesAlive; i++ {
		par := ps.particles[i]
		ps.Info.Sprite.SetColor(par.color.HWColor())
		ps.Info.Sprite.RenderEx(par.location.X+ps.tx, par.location.Y+ps.ty, par.spin*par.age, par.size)
	}

	ps.Info.Sprite.SetColor(col)
}

func (ps *ParticleSystem) FireAt(x, y float64) {
	ps.Stop()
	ps.MoveTo(x, y)
	ps.Fire()
}

func (ps *ParticleSystem) Fire() {
	if ps.Info.Lifetime == -1.0 {
		ps.age = -1.0
	} else {
		ps.age = 0.0
	}
	ps.residue = 0.0
}

func (ps *ParticleSystem) Stop(a ...interface{}) {
	kill_particles := false

	if len(a) == 1 {
		if kp, ok := a[0].(bool); ok {
			kill_particles = kp
		}
	}

	ps.age = -2.0
	if kill_particles {
		ps.particlesAlive = 0
		ps.boundingBox.Clear()
	}
}

func (ps *ParticleSystem) Update(deltaTime float64) {
	if ps.updateSpeed == 0.0 {
		ps.update(deltaTime)
	} else {
		ps.residue += deltaTime
		if ps.residue >= ps.updateSpeed {
			ps.update(ps.updateSpeed)
			for ps.residue >= ps.updateSpeed {
				ps.residue -= ps.updateSpeed
			}
		}
	}
}

func (ps *ParticleSystem) MoveTo(x, y float64, a ...interface{}) {
	move_particles := false

	if len(a) == 1 {
		if mp, ok := a[0].(bool); ok {
			move_particles = mp
		}
	}

	if move_particles {
		dx := x - ps.location.X
		dy := y - ps.location.Y

		for i := 0; i < ps.particlesAlive; i++ {
			ps.particles[i].location.X += dx
			ps.particles[i].location.Y += dy
		}

		ps.prevLocation.X = ps.prevLocation.X + dx
		ps.prevLocation.Y = ps.prevLocation.Y + dy
	} else {
		if ps.age == -2.0 {
			ps.prevLocation.X, ps.prevLocation.Y = x, y
		} else {
			ps.prevLocation.X, ps.prevLocation.Y = ps.location.X, ps.location.Y
		}
	}

	ps.location.X, ps.location.Y = x, y
}

func (ps *ParticleSystem) Transpose(x, y float64) {
	ps.tx, ps.ty = x, y
}

func (ps *ParticleSystem) TrackBoundingBox(track bool) {
	ps.updateBoundingBox = track
}

func (ps ParticleSystem) ParticlesAlive() int {
	return ps.particlesAlive
}

func (ps ParticleSystem) Age() float64 {
	return ps.age
}

func (ps ParticleSystem) Position() (x, y float64) {
	return ps.location.X, ps.location.Y
}

func (ps ParticleSystem) Transposition() (x, y float64) {
	return ps.tx, ps.ty
}

func (ps ParticleSystem) BoundingBox(rect *Rect) *Rect {
	rect.SetRect(ps.boundingBox)
	return rect
}

func (ps *ParticleSystem) update(deltaTime float64) {
	if ps.age >= 0 {
		ps.age += deltaTime
		if ps.age >= ps.Info.Lifetime {
			ps.age = -2.0
		}
	}

	// update all alive particles
	if ps.updateBoundingBox {
		ps.boundingBox.Clear()
	}

	for i := 0; i < ps.particlesAlive; i++ {
		par := &ps.particles[i]
		par.age += deltaTime
		if par.age >= par.terminalAge {
			ps.particlesAlive--

			par = &ps.particles[ps.particlesAlive]
			i--
			continue
		}

		accel := par.location.Sub(ps.location)
		accel.Normalize()
		accel2 := accel
		accel.MulEqual(par.radialAccel)

		// accel2.Rotate(Pi_2);
		// the following is faster
		ang := accel2.X
		accel2.X = -accel2.Y
		accel2.Y = ang

		accel2.MulEqual(par.tangentialAccel)
		par.velocity.AddEqual((accel.Add(accel2)).Mul(deltaTime))
		par.velocity.Y += par.gravity * deltaTime

		par.location.AddEqual(par.velocity)

		par.spin += par.spinDelta * deltaTime
		par.size += par.sizeDelta * deltaTime
		par.color.AddEqual(par.colorDelta.MulScalar(deltaTime))

		if ps.updateBoundingBox {
			ps.boundingBox.Encapsulate(par.location.X, par.location.Y)
		}
	}

	// generate new particles
	if ps.age != -2.0 {
		particles_needed := float64(ps.Info.Emission)*deltaTime + ps.emissionResidue
		particles_created := int(particles_needed)
		ps.emissionResidue = particles_needed - float64(particles_created)

		par := &ps.particles[ps.particlesAlive]

		for i := 0; i < particles_created; i++ {
			if ps.particlesAlive >= hgeMAX_PARTICLES {
				break
			}

			par.age = 0.0
			par.terminalAge = hge.RandomFloat(ps.Info.LifeMin, ps.Info.LifeMax)

			par.location = ps.prevLocation.Add(ps.location.Sub(ps.prevLocation).Mul(hge.RandomFloat(0.0, 1.0)))
			par.location.X += hge.RandomFloat(-2.0, 2.0)
			par.location.Y += hge.RandomFloat(-2.0, 2.0)

			ang := ps.Info.Direction - hge.Pi_2 + hge.RandomFloat(0, ps.Info.Spread) - ps.Info.Spread/2.0
			if ps.Info.Relative {
				ang += ps.prevLocation.Sub(ps.location).Angle() + hge.Pi_2
			}
			par.velocity.X = math.Cos(ang)
			par.velocity.Y = math.Sin(ang)
			par.velocity.MulEqual(hge.RandomFloat(ps.Info.SpeedMin, ps.Info.SpeedMax))

			par.gravity = hge.RandomFloat(ps.Info.GravityMin, ps.Info.GravityMax)
			par.radialAccel = hge.RandomFloat(ps.Info.RadialAccelMin, ps.Info.RadialAccelMax)
			par.tangentialAccel = hge.RandomFloat(ps.Info.TangentialAccelMin, ps.Info.TangentialAccelMax)

			par.size = hge.RandomFloat(ps.Info.SizeStart, ps.Info.SizeStart+(ps.Info.SizeEnd-ps.Info.SizeStart)*ps.Info.SizeVar)
			par.sizeDelta = (ps.Info.SizeEnd - par.size) / par.terminalAge

			par.spin = hge.RandomFloat(ps.Info.SpinStart, ps.Info.SpinStart+(ps.Info.SpinEnd-ps.Info.SpinStart)*ps.Info.SpinVar)
			par.spinDelta = (ps.Info.SpinEnd - par.spin) / par.terminalAge

			par.color.R = hge.RandomFloat(ps.Info.ColorStart.R, ps.Info.ColorStart.R+(ps.Info.ColorEnd.R-ps.Info.ColorStart.R)*ps.Info.ColorVar)
			par.color.G = hge.RandomFloat(ps.Info.ColorStart.G, ps.Info.ColorStart.G+(ps.Info.ColorEnd.G-ps.Info.ColorStart.G)*ps.Info.ColorVar)
			par.color.B = hge.RandomFloat(ps.Info.ColorStart.B, ps.Info.ColorStart.B+(ps.Info.ColorEnd.B-ps.Info.ColorStart.B)*ps.Info.ColorVar)
			par.color.A = hge.RandomFloat(ps.Info.ColorStart.A, ps.Info.ColorStart.A+(ps.Info.ColorEnd.A-ps.Info.ColorStart.A)*ps.Info.AlphaVar)

			par.colorDelta.R = (ps.Info.ColorEnd.R - par.color.R) / par.terminalAge
			par.colorDelta.G = (ps.Info.ColorEnd.G - par.color.G) / par.terminalAge
			par.colorDelta.B = (ps.Info.ColorEnd.B - par.color.B) / par.terminalAge
			par.colorDelta.A = (ps.Info.ColorEnd.A - par.color.A) / par.terminalAge

			if ps.updateBoundingBox {
				ps.boundingBox.Encapsulate(par.location.X, par.location.Y)
			}

			ps.particlesAlive++
			par = &ps.particles[ps.particlesAlive]
		}
	}

	ps.prevLocation = ps.location

}

type ParticleManager struct {
	fps  float64
	ps   int
	x, y float64
	list []*ParticleSystem
}

func NewParticleManager(a ...interface{}) ParticleManager {
	fps := 0.0

	if len(a) == 1 {
		if f, ok := a[0].(float64); ok {
			fps = f
		}
	}

	var pm ParticleManager

	pm.ps = 0
	pm.fps = fps
	pm.x, pm.y = 0.0, 0.0

	pm.list = make([]*ParticleSystem, hgeMAX_PSYSTEMS+1)

	return pm
}

func (pm *ParticleManager) Update(dt float64) {
	for i := 0; i < pm.ps; i++ {
		pm.list[i].Update(dt)
		if pm.list[i].Age() == -2.0 && pm.list[i].ParticlesAlive() == 0 {
			pm.list[i] = nil
			pm.list[i] = pm.list[pm.ps-1]
			pm.ps--
			i--
		}
	}
}

func (pm *ParticleManager) Render() {
	for i := 0; i < pm.ps; i++ {
		pm.list[i].Render()
	}
}

func (pm *ParticleManager) SpawPS(psi ParticleSystemInfo, x, y float64) *ParticleSystem {
	if pm.ps == hgeMAX_PSYSTEMS {
		return nil
	}

	pm.list[pm.ps] = NewParticleSystemWithInfo(psi, pm.fps)
	pm.list[pm.ps].FireAt(x, y)
	pm.list[pm.ps].Transpose(pm.x, pm.y)
	pm.ps++

	return pm.list[pm.ps-1]

}

func (pm ParticleManager) IsPSAlive(ps *ParticleSystem, x, y float64) bool {
	for i := 0; i < pm.ps; i++ {
		if pm.list[i] == ps {
			return true
		}
	}

	return false
}

func (pm *ParticleManager) Transpose(x, y float64) {
	for i := 0; i < pm.ps; i++ {
		pm.list[i].Transpose(x, y)
	}

	pm.x, pm.y = x, y
}

func (pm ParticleManager) Transposition() (dx, dy float64) {
	return pm.x, pm.y
}

func (pm *ParticleManager) KillPS(ps *ParticleSystem) {
	for i := 0; i < pm.ps; i++ {
		if pm.list[i] == ps {
			pm.list[i] = nil
			pm.list[i] = pm.list[pm.ps-1]
			pm.ps--
			return
		}
	}
}

func (pm *ParticleManager) KillAll() {
	for i := 0; i < pm.ps; i++ {
		pm.list[i] = nil
	}
	pm.ps = 0
}
