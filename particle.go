package hge

import (
	"math"
)

const (
	hge_MAX_PARTICLES = 500
	hge_MAX_PSYSTEMS  = 100
)

type particle struct {
	location Vector
	velocity Vector

	gravity, radial_accel, tangential_accel float64
	spin, spin_delta                        float64
	size, size_delta                        float64
	color, color_delta                      ColorRGB

	age, terminal_age float64
}

type ParticleSystemInfo struct {
	Sprite                     *Sprite // texture + blend mode
	Emission                   int     // particles per sec
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

	hge *HGE

	update_speed, residue, age, emission_residue float64
	prev_location, location                      Vector

	tx, ty float64

	particles_alive     int
	bounding_box        Rect
	update_bounding_box bool
	particles           []particle
}

func NewParticleSystem(filename string, sprite *Sprite, a ...interface{}) *ParticleSystem {
	fps := 0.0

	if len(a) == 1 {
		if f, ok := a[0].(float64); ok {
			fps = f
		}
	}

	ps := new(ParticleSystem)

	ps.hge = Create(VERSION)

	psi, size := ps.hge.Resource_Load(filename)

	if psi == nil || size == 0 {
		return nil
	}

	ps.location.X, ps.prev_location.X = 0.0, 0.0
	ps.location.Y, ps.prev_location.Y = 0.0, 0.0
	ps.tx, ps.ty = 0, 0

	ps.emission_residue = 0.0
	ps.particles_alive = 0
	ps.age = -2.0
	if fps != 0.0 {
		ps.update_speed = 1.0 / fps
	} else {
		ps.update_speed = 0.0
	}
	ps.residue = 0.0

	ps.bounding_box.Clear()
	ps.update_bounding_box = false

	return ps
}

func NewParticleSystemWithInfo(psi *ParticleSystemInfo, a ...interface{}) *ParticleSystem {
	fps := 0.0

	if len(a) == 1 {
		if f, ok := a[0].(float64); ok {
			fps = f
		}
	}

	ps := new(ParticleSystem)

	ps.hge = Create(VERSION)

	ps.location.X, ps.prev_location.X = 0.0, 0.0
	ps.location.Y, ps.prev_location.Y = 0.0, 0.0
	ps.tx, ps.ty = 0, 0

	ps.emission_residue = 0.0
	ps.particles_alive = 0
	ps.age = -2.0
	if fps != 0.0 {
		ps.update_speed = 1.0 / fps
	} else {
		ps.update_speed = 0.0
	}
	ps.residue = 0.0

	return ps
}

func (ps *ParticleSystem) Render() {
	col := ps.Info.Sprite.GetColor()

	for i := 0; i < ps.particles_alive; i++ {
		par := ps.particles[i]
		ps.Info.Sprite.SetColor(par.color.GetHWColor())
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
		ps.particles_alive = 0
		ps.bounding_box.Clear()
	}
}

func (ps *ParticleSystem) Update(delta_time float64) {
	if ps.update_speed == 0.0 {
		ps.update(delta_time)
	} else {
		ps.residue += delta_time
		if ps.residue >= ps.update_speed {
			ps.update(ps.update_speed)
			for ps.residue >= ps.update_speed {
				ps.residue -= ps.update_speed
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

		for i := 0; i < ps.particles_alive; i++ {
			ps.particles[i].location.X += dx
			ps.particles[i].location.Y += dy
		}

		ps.prev_location.X = ps.prev_location.X + dx
		ps.prev_location.Y = ps.prev_location.Y + dy
	} else {
		if ps.age == -2.0 {
			ps.prev_location.X, ps.prev_location.Y = x, y
		} else {
			ps.prev_location.X, ps.prev_location.Y = ps.location.X, ps.location.Y
		}
	}

	ps.location.X, ps.location.Y = x, y
}

func (ps *ParticleSystem) Transpose(x, y float64) {
	ps.tx, ps.ty = x, y
}

func (ps *ParticleSystem) TrackBoundingBox(track bool) {
	ps.update_bounding_box = track
}

func (ps ParticleSystem) GetParticlesAlive() int {
	return ps.particles_alive
}

func (ps ParticleSystem) GetAge() float64 {
	return ps.age
}

func (ps ParticleSystem) GetPosition() (x, y float64) {
	return ps.location.X, ps.location.Y
}

func (ps ParticleSystem) GetTransposition() (x, y float64) {
	return ps.tx, ps.ty
}

func (ps ParticleSystem) GetBoundingBox(rect *Rect) *Rect {
	rect.SetRect(ps.bounding_box)
	return rect
}

func (ps *ParticleSystem) update(delta_time float64) {
	if ps.age >= 0 {
		ps.age += delta_time
		if ps.age >= ps.Info.Lifetime {
			ps.age = -2.0
		}
	}

	// update all alive particles
	if ps.update_bounding_box {
		ps.bounding_box.Clear()
	}

	for i := 0; i < ps.particles_alive; i++ {
		par := ps.particles[i]
		par.age += delta_time
		if par.age >= par.terminal_age {
			ps.particles_alive--

			// memcpy(par, &particles[nParticlesAlive], sizeof(hgeParticle));
			i--
			continue
		}

		accel := par.location.Subtract(ps.location)
		accel.Normalize()
		accel2 := accel
		accel.MultiplyEqual(par.radial_accel)

		// vecAccel2.Rotate(M_PI_2);
		// the following is faster
		ang := accel2.X
		accel2.X = -accel2.Y
		accel2.Y = ang

		accel2.MultiplyEqual(par.tangential_accel)
		par.velocity.AddEqual((accel.Add(accel2)).Multiply(delta_time))
		par.velocity.Y += par.gravity * delta_time

		par.location.AddEqual(par.velocity)

		par.spin += par.spin_delta * delta_time
		par.size += par.size_delta * delta_time
		// par.color += par.color_delta*delta_time

		if ps.update_bounding_box {
			ps.bounding_box.Encapsulate(par.location.X, par.location.Y)
		}
	}

	// generate new particles
	if ps.age != -2.0 {
		particles_needed := ps.Info.Emission*int(delta_time) + int(ps.emission_residue)
		particles_created := particles_needed
		ps.emission_residue = float64(particles_needed - particles_created)

		par := ps.particles[ps.particles_alive]

		for i := 0; i < particles_created; i++ {
			if ps.particles_alive >= hge_MAX_PARTICLES {
				break
			}

			par.age = 0.0
			par.terminal_age = ps.hge.Random_Float(ps.Info.LifeMin, ps.Info.LifeMax)

			//par.location = ps.prev_location + (ps.location-ps.prev_location)*ps.hge.Random_Float(0.0, 1.0)
			par.location.X += ps.hge.Random_Float(-2.0, 2.0)
			par.location.Y += ps.hge.Random_Float(-2.0, 2.0)

			ang := ps.Info.Direction - Pi_2 + ps.hge.Random_Float(0, ps.Info.Spread) - ps.Info.Spread/2.0
			//if ps.Info.Relative {
			//	ang += (ps.prev_location - ps.location).Angle() + Pi_2
			//}
			par.velocity.X = math.Cos(ang)
			par.velocity.Y = math.Sin(ang)
			//par.velocity *= ps.hge.Random_Float(ps.Info.SpeedMin, ps.Info.SpeedMax)

			par.gravity = ps.hge.Random_Float(ps.Info.GravityMin, ps.Info.GravityMax)
			par.radial_accel = ps.hge.Random_Float(ps.Info.RadialAccelMin, ps.Info.RadialAccelMax)
			par.tangential_accel = ps.hge.Random_Float(ps.Info.TangentialAccelMin, ps.Info.TangentialAccelMax)

			par.size = ps.hge.Random_Float(ps.Info.SizeStart, ps.Info.SizeStart+(ps.Info.SizeEnd-ps.Info.SizeStart)*ps.Info.SizeVar)
			par.size_delta = (ps.Info.SizeEnd - par.size) / par.terminal_age

			par.spin = ps.hge.Random_Float(ps.Info.SpinStart, ps.Info.SpinStart+(ps.Info.SpinEnd-ps.Info.SpinStart)*ps.Info.SpinVar)
			par.spin_delta = (ps.Info.SpinEnd - par.spin) / par.terminal_age

			par.color.R = ps.hge.Random_Float(ps.Info.ColorStart.R, ps.Info.ColorStart.R+(ps.Info.ColorEnd.R-ps.Info.ColorStart.R)*ps.Info.ColorVar)
			par.color.G = ps.hge.Random_Float(ps.Info.ColorStart.G, ps.Info.ColorStart.G+(ps.Info.ColorEnd.G-ps.Info.ColorStart.G)*ps.Info.ColorVar)
			par.color.B = ps.hge.Random_Float(ps.Info.ColorStart.B, ps.Info.ColorStart.B+(ps.Info.ColorEnd.B-ps.Info.ColorStart.B)*ps.Info.ColorVar)
			par.color.A = ps.hge.Random_Float(ps.Info.ColorStart.A, ps.Info.ColorStart.A+(ps.Info.ColorEnd.A-ps.Info.ColorStart.A)*ps.Info.AlphaVar)

			par.color_delta.R = (ps.Info.ColorEnd.R - par.color.R) / par.terminal_age
			par.color_delta.G = (ps.Info.ColorEnd.G - par.color.G) / par.terminal_age
			par.color_delta.B = (ps.Info.ColorEnd.B - par.color.B) / par.terminal_age
			par.color_delta.A = (ps.Info.ColorEnd.A - par.color.A) / par.terminal_age

			if ps.update_bounding_box {
				ps.bounding_box.Encapsulate(par.location.X, par.location.Y)
			}

			ps.particles_alive++
			par = ps.particles[ps.particles_alive]
		}
	}
	ps.prev_location = ps.location

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

	return pm
}

func (pm *ParticleManager) Update(dt float64) {
	for i := 0; i < pm.ps; i++ {
		pm.list[i].Update(dt)
		if pm.list[i].GetAge() == -2.0 && pm.list[i].GetParticlesAlive() == 0 {
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

func (pm *ParticleManager) SpawPS(psi *ParticleSystemInfo, x, y float64) *ParticleSystem {
	if pm.ps == hge_MAX_PSYSTEMS {
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

func (pm ParticleManager) GetTransposition() (dx, dy float64) {
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
