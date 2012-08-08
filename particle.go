package hge

const (
	hge_MAX_PARTICLES = 500
	hge_MAX_PSYSTEMS  = 100
)

type particle struct {
	location Vector
	velocity Vector

	gravity, radial_accel, tangential_accel float32
	spin, spin_delta                        float32
	size, size_delta                        float32
	color, color_delta                      Color

	age, terminal_age float32
}

type ParticleSystemInfo struct {
	Sprite                     *Sprite // texture + blend mode
	Emission                   int     // particles per sec
	Lifetime, LiveMin, LifeMax float32
	Direction, Spread          float32

	Relation bool

	SpeedMin, SpeedMax                     float32
	GravityMin, GravityMax                 float32
	RadialAccelMin, RadialAccelMax         float32
	TangentialAccelMin, TangentialAccelMax float32
	SizeStart, SizeEnd, SizeVar            float32
	SpinStart, SpinEnd, SpinVar            float32

	ColorStart, ColorEnd Color

	ColorVar, AlphaVar float32
}

type ParticleSystem struct {
	Info ParticleSystemInfo

	hge *HGE

	update_speed, residue, age, emission_residue float32
	prev_location, location                      Vector

	tx, ty float32

	particles_alive     int
	bounding_box        Rect
	update_bounding_box bool
	particles           []particle
}

func NewParticleSystem(filename string, sprite *Sprite, a ...interface{}) ParticleSystem {
	fps := float32(0.0)

	if len(a) == 1 {
		if f, ok := a[0].(float32); ok {
			fps = f
		}
	}

	var ps ParticleSystem

	ps.location.X = 0.0
	ps.prev_location.X = 0.0
	ps.location.Y = 0.0
	ps.prev_location.Y = 0.0
	ps.tx = 0
	ps.ty = 0

	ps.emission_residue = 0.0
	ps.particles_alive = 0
	ps.age = -2.0
	if fps != 0.0 {
		ps.update_speed = 1.0 / fps
	} else {
		ps.update_speed = 0.0
	}

	ps.bounding_box.Clear()
	ps.update_bounding_box = false

	return ps
}

func NewParticleSystemWithInfo(psi *ParticleSystemInfo, a ...interface{}) ParticleSystem {
	fps := float32(0.0)

	if len(a) == 1 {
		if f, ok := a[0].(float32); ok {
			fps = f
		}
	}

	var ps ParticleSystem

	ps.location.X = 0.0
	ps.prev_location.X = 0.0
	ps.location.Y = 0.0
	ps.prev_location.Y = 0.0
	ps.tx = 0
	ps.ty = 0

	ps.emission_residue = 0.0
	ps.particles_alive = 0
	ps.age = -2.0
	if fps != 0.0 {
		ps.update_speed = 1.0 / fps
	} else {
		ps.update_speed = 0.0
	}

	ps.bounding_box.Clear()
	ps.update_bounding_box = false

	return ps
}

func (ps *ParticleSystem) Equal(ps1 *ParticleSystem) *ParticleSystem {
	return ps
}

func (ps *ParticleSystem) Render() {
}

func (ps *ParticleSystem) FireAt(x, y float32) {
}

func (ps *ParticleSystem) Fire() {
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

func (ps *ParticleSystem) Update(dalta_time float32) {
}

func (ps *ParticleSystem) MoveTo(x, y float32, a ...interface{}) {
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
			ps.prev_location.X = x
			ps.prev_location.Y = y
		} else {
			ps.prev_location.X = ps.location.X
			ps.prev_location.Y = ps.location.Y
		}
	}

	ps.location.X = x
	ps.location.Y = y

}

func (ps *ParticleSystem) Transpose(x, y float32) {
	ps.tx = x
	ps.ty = y
}

func (ps *ParticleSystem) TrackBoundingBox(track bool) {
	ps.update_bounding_box = track
}

func (ps ParticleSystem) GetParticlesAlive() int {
	return ps.particles_alive
}

func (ps ParticleSystem) GetAge() float32 {
	return ps.age
}

func (ps ParticleSystem) GetPosition() (x, y float32) {
	return ps.location.X, ps.location.Y
}

func (ps ParticleSystem) GetTransposition() (x, y float32) {
	return ps.tx, ps.ty
}

func (ps ParticleSystem) GetBoundingBox(rect *Rect) *Rect {
	rect.SetRect(ps.bounding_box)
	return rect
}

func (ps *ParticleSystem) update(delta_time float32) {
}

type ParticleManager struct {
	fps  float32
	ps   int
	x, y float32
	list []ParticleSystem
}

func NewParticleManager(a ...interface{}) ParticleManager {
	fps := float32(0.0)

	if len(a) == 1 {
		if f, ok := a[0].(float32); ok {
			fps = f
		}
	}

	var pm ParticleManager

	pm.ps = 0
	pm.fps = fps
	pm.x = 0.0
	pm.y = 0.0

	return pm
}

func (pm *ParticleManager) Update(dt float32) {
}

func (pm *ParticleManager) Render() {
}

func (pm *ParticleManager) SpawPS(psi *ParticleSystemInfo, x, y float32) *ParticleSystem {
	if pm.ps == hge_MAX_PSYSTEMS {
		return nil
	}

	pm.list[pm.ps] = NewParticleSystemWithInfo(psi, pm.fps)
	pm.list[pm.ps].FireAt(x, y)
	pm.list[pm.ps].Transpose(pm.x, pm.y)
	pm.ps++

	return &pm.list[pm.ps-1]

}

func (pm ParticleManager) IsPSAlive(ps *ParticleSystem, x, y float32) bool {
	for i := 0; i < pm.ps; i++ {
		//if pm.list[i] == *ps {
		//	return true
		//}
	}

	return false
}

func (pm *ParticleManager) Transpose(x, y float32) {
}

func (pm ParticleManager) GetTransposition() (dx, dy float32) {
	return pm.x, pm.y
}

func (pm *ParticleManager) KillPS(ps *ParticleSystem) {
}

func (pm *ParticleManager) KillAll() {
}
