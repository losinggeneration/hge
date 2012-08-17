package main

// smaller sun & moon, underwater
// moon shape, hide stars behind the moon

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/color"
	dist "github.com/losinggeneration/hge-go/helpers/distortionmesh"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	HGE "github.com/losinggeneration/hge-go/hge"
	"math"
	"time"
)

// Pointer to the HGE interface (helper classes require this to work)
var (
	fnt *font.Font
)

// Simulation constants
const (
	SCREEN_WIDTH    = 800
	SCREEN_HEIGHT   = 600
	NUM_STARS       = 100
	SEA_SUBDIVISION = 16

	SKY_HEIGHT    = (SCREEN_HEIGHT * 0.6)
	STARS_HEIGHT  = (SKY_HEIGHT * 0.9)
	ORBITS_RADIUS = (SCREEN_WIDTH * 0.43)
)

var (
	skyTopColors = []HGE.Dword{0xFF15092A, 0xFF6C6480, 0xFF89B9D0}
	skyBtmColors = []HGE.Dword{0xFF303E57, 0xFFAC7963, 0xFFCAD7DB}
	seaTopColors = []HGE.Dword{0xFF3D546B, 0xFF927E76, 0xFF86A2AD}
	seaBtmColors = []HGE.Dword{0xFF1E394C, 0xFF2F4E64, 0xFF2F4E64}
	seq          = []int{0, 0, 1, 2, 2, 2, 1, 0, 0}
)

// Simulation resource handles
var (
	texObjects                          HGE.Texture
	sky, sun, moon, glow, seaglow, star sprite.Sprite
	sea                                 dist.DistortionMesh
	colWhite                            color.ColorRGB
)

// Simulation state variables
var (
	timet float64 // 0-24 hrs
	speed float64 // hrs per second

	seq_id      int
	seq_residue float64

	starX [NUM_STARS]float64       // x
	starY [NUM_STARS]float64       // y
	starS [NUM_STARS]float64       // scale
	starA [NUM_STARS]float64       // alpha
	seaP  [SEA_SUBDIVISION]float64 // phase shift array

	colSkyTop, colSkyBtm, colSeaTop, colSeaBtm color.ColorRGB

	colSun, colSunGlow, colMoon, colMoonGlow, colSeaGlow color.ColorRGB

	sunX, sunY, sunS, sunGlowS     float64
	moonX, moonY, moonS, moonGlowS float64
	seaGlowX, seaGlowSX, seaGlowSY float64
)

///////////////////////// Implementation ///////////////////////////
func frame() int {
	// Process keys
	switch HGE.GetKey() {
	case HGE.K_0:
		speed = 0.0
	case HGE.K_1:
		speed = 0.1
	case HGE.K_2:
		speed = 0.2
	case HGE.K_3:
		speed = 0.4
	case HGE.K_4:
		speed = 0.8
	case HGE.K_5:
		speed = 1.6
	case HGE.K_6:
		speed = 3.2
	case HGE.K_7:
		speed = 6.4
	case HGE.K_8:
		speed = 12.8
	case HGE.K_9:
		speed = 25.6
	case HGE.K_ESCAPE:
		return 1
	}

	// Update scene
	UpdateSimulation()

	return 0
}

func render() int {
	// 	int hrs, mins, secs;
	// 	float tmp;

	// Calculate display time
	hrs := int(math.Floor(timet))
	tmp := (timet - float64(hrs)) * 60.0
	mins := int(math.Floor(tmp))
	secs := int(math.Floor((tmp - float64(mins)) * 60.0))

	// Render scene
	HGE.GfxBeginScene()
	RenderSimulation()
	fnt.Printf(7, 7, font.TEXT_LEFT, "Keys 1-9 to adjust simulation speed, 0 - real time\nFPS: %d", HGE.GetFPS())
	fnt.Printf(SCREEN_WIDTH-50, 7, font.TEXT_LEFT, "%02d:%02d:%02d", hrs, mins, secs)
	HGE.GfxEndScene()

	return 0
}

func main() {
	defer HGE.Free()

	// Set desired system states and initialize HGE
	HGE.SetState(HGE.LOGFILE, "tutorial08.log")
	HGE.SetState(HGE.FRAMEFUNC, frame)
	HGE.SetState(HGE.RENDERFUNC, render)
	HGE.SetState(HGE.TITLE, "HGE Tutorial 08 - The Big Calm")
	HGE.SetState(HGE.USESOUND, false)
	HGE.SetState(HGE.WINDOWED, true)
	HGE.SetState(HGE.SCREENWIDTH, SCREEN_WIDTH)
	HGE.SetState(HGE.SCREENHEIGHT, SCREEN_HEIGHT)
	HGE.SetState(HGE.SCREENBPP, 32)

	if err := HGE.Initiate(); err == nil {
		defer HGE.Shutdown()

		fnt = font.NewFont("font1.fnt")

		if !InitSimulation() {
			// If one of the data files is not found, display an error message and shutdown
			fmt.Println("Error: Can't load resources. See log for details.\n")
			return
		}

		HGE.Start()

		DoneSimulation()
	}

	return
}

func GetTime() float64 {
	t := time.Now()
	tmp := float64(t.Second())
	tmp = float64(t.Minute()) + tmp/60.0
	tmp = float64(t.Hour()) + tmp/60.0

	return tmp
}

func InitSimulation() bool {
	// Load texture
	texObjects = HGE.LoadTexture("objects.png")
	if texObjects == 0 {
		return false
	}

	// Create sprites
	sky = sprite.NewSprite(0, 0, 0, SCREEN_WIDTH, SKY_HEIGHT)
	sea = dist.NewDistortionMesh(SEA_SUBDIVISION, SEA_SUBDIVISION)
	sea.SetTextureRect(0, 0, SCREEN_WIDTH, SCREEN_HEIGHT-SKY_HEIGHT)

	sun = sprite.NewSprite(texObjects, 81, 0, 114, 114)
	sun.SetHotSpot(57, 57)
	moon = sprite.NewSprite(texObjects, 0, 0, 81, 81)
	moon.SetHotSpot(40, 40)
	star = sprite.NewSprite(texObjects, 72, 81, 9, 9)
	star.SetHotSpot(5, 5)

	glow = sprite.NewSprite(texObjects, 128, 128, 128, 128)
	glow.SetHotSpot(64, 64)
	glow.SetBlendMode(HGE.BLEND_COLORADD | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE)
	seaglow = sprite.NewSprite(texObjects, 128, 224, 128, 32)
	seaglow.SetHotSpot(64, 0)
	seaglow.SetBlendMode(HGE.BLEND_COLORADD | HGE.BLEND_ALPHAADD | HGE.BLEND_NOZWRITE)

	// Initialize simulation state
	colWhite.SetHWColor(0xFFFFFFFF)
	timet = GetTime()
	speed = 0.0

	for i := 0; i < NUM_STARS; i++ { // star positions
		starX[i] = HGE.RandomFloat(0, SCREEN_WIDTH)
		starY[i] = HGE.RandomFloat(0, STARS_HEIGHT)
		starS[i] = HGE.RandomFloat(0.1, 0.7)
	}

	for i := 0; i < SEA_SUBDIVISION; i++ { // sea waves phase shifts
		seaP[i] = float64(i) + HGE.RandomFloat(-15.0, 15.0)
	}

	// Systems are ready now!
	return true
}

func DoneSimulation() {
	// Free texture
	texObjects.Free()
}

func UpdateSimulation() {
	cellw := SCREEN_WIDTH / (SEA_SUBDIVISION - 1)

	var col1, col2 color.ColorRGB

	// Update time of day
	if speed == 0.0 {
		timet = GetTime()
	} else {
		timet += HGE.NewTimer().Delta() * speed
		if timet >= 24.0 {
			timet -= 24.0
		}
	}

	seq_id = int(timet / 3)
	seq_residue = timet/3 - float64(seq_id)
	zenith := -(timet/12.0*HGE.Pi - HGE.Pi_2)

	// Interpolate sea and sky colors
	col1.SetHWColor(skyTopColors[seq[seq_id]])
	col2.SetHWColor(skyTopColors[seq[seq_id+1]])
	colSkyTop = col2.MulScalar(seq_residue).Add(col1.MulScalar(1.0 - seq_residue))

	col1.SetHWColor(skyBtmColors[seq[seq_id]])
	col2.SetHWColor(skyBtmColors[seq[seq_id+1]])
	colSkyBtm = col2.MulScalar(seq_residue).Add(col1.MulScalar(1.0 - seq_residue))

	col1.SetHWColor(seaTopColors[seq[seq_id]])
	col2.SetHWColor(seaTopColors[seq[seq_id+1]])
	colSeaTop = col2.MulScalar(seq_residue).Add(col1.MulScalar(1.0 - seq_residue))

	col1.SetHWColor(seaBtmColors[seq[seq_id]])
	col2.SetHWColor(seaBtmColors[seq[seq_id+1]])
	colSeaBtm = col2.MulScalar(seq_residue).Add(col1.MulScalar(1.0 - seq_residue))

	var a float64
	// Update stars
	if seq_id >= 6 || seq_id < 2 {
		for i := 0; i < NUM_STARS; i++ {
			a = 1.0 - starY[i]/STARS_HEIGHT
			a *= HGE.RandomFloat(0.6, 1.0)
			if seq_id >= 6 {
				a *= math.Sin((timet - 18.0) / 6.0 * HGE.Pi_2)
			} else {
				a *= math.Sin((1.0 - timet/6.0) * HGE.Pi_2)
			}
			starA[i] = a
		}
	}

	// Calculate sun position, scale and colors
	if seq_id == 2 {
		a = math.Sin(seq_residue * HGE.Pi_2)
	} else if seq_id == 5 {
		a = math.Cos(seq_residue * HGE.Pi_2)
	} else if seq_id > 2 && seq_id < 5 {
		a = 1.0
	} else {
		a = 0.0
	}

	colSun.SetHWColor(0xFFEAE1BE)
	colSun = colSun.MulScalar(1 - a).Add(colWhite.MulScalar(a))

	a = (math.Cos(timet/6.0*HGE.Pi) + 1.0) / 2.0
	if seq_id >= 2 && seq_id <= 6 {
		colSunGlow = colWhite.MulScalar(a)
		colSunGlow.A = 1.0
	} else {
		colSunGlow.SetHWColor(0xFF000000)
	}

	sunX = SCREEN_WIDTH*0.5 + math.Cos(zenith)*ORBITS_RADIUS
	sunY = SKY_HEIGHT*1.2 + math.Sin(zenith)*ORBITS_RADIUS
	sunS = 1.0 - 0.3*math.Sin((timet-6.0)/12.0*HGE.Pi)
	sunGlowS = 3.0*(1.0-a) + 3.0

	// Calculate moon position, scale and colors
	if seq_id >= 6 {
		a = math.Sin((timet - 18.0) / 6.0 * HGE.Pi_2)
	} else {
		a = math.Sin((1.0 - timet/6.0) * HGE.Pi_2)
	}
	colMoon.SetHWColor(0x20FFFFFF)
	colMoon = colMoon.MulScalar(1 - a).Add(colWhite.MulScalar(a))

	colMoonGlow = colWhite
	colMoonGlow.A = 0.5 * a

	moonX = SCREEN_WIDTH*0.5 + math.Cos(zenith-HGE.Pi)*ORBITS_RADIUS
	moonY = SKY_HEIGHT*1.2 + math.Sin(zenith-HGE.Pi)*ORBITS_RADIUS
	moonS = 1.0 - 0.3*math.Sin((timet+6.0)/12.0*HGE.Pi)
	moonGlowS = a*0.4 + 0.5

	// Calculate sea glow
	if timet > 19.0 || timet < 4.5 { // moon
		a = 0.2 // intensity
		if timet > 19.0 && timet < 20.0 {
			a *= (timet - 19.0)
		} else if timet > 3.5 && timet < 4.5 {
			a *= 1.0 - (timet - 3.5)
		}

		colSeaGlow = colMoonGlow
		colSeaGlow.A = a
		seaGlowX = moonX
		seaGlowSX = moonGlowS * 3.0
		seaGlowSY = moonGlowS * 2.0
	} else if timet > 6.5 && timet < 19.0 { // sun
		a = 0.3 // intensity
		if timet < 7.5 {
			a *= (timet - 6.5)
		} else if timet > 18.0 {
			a *= 1.0 - (timet - 18.0)
		}

		colSeaGlow = colSunGlow
		colSeaGlow.A = a
		seaGlowX = sunX
		seaGlowSX = sunGlowS
		seaGlowSY = sunGlowS * 0.6
	} else {
		colSeaGlow.A = 0.0
	}

	var dwCol1 HGE.Dword
	// Move waves and update sea color
	for i := 1; i < SEA_SUBDIVISION-1; i++ {
		a = float64(i) / (SEA_SUBDIVISION - 1)
		col1 = colSeaTop.MulScalar(1 - a).Add(colSeaBtm.MulScalar(a))
		dwCol1 = col1.HWColor()
		fTime := 2.0 * HGE.NewTimer().Time()
		a *= 20

		for j := 0; j < SEA_SUBDIVISION; j++ {
			sea.SetColor(j, i, dwCol1)

			dy := a * math.Sin(seaP[i]+(float64(j)/(SEA_SUBDIVISION-1)-0.5)*HGE.Pi*16.0-fTime)
			sea.SetDisplacement(j, i, 0.0, dy, dist.DISP_NODE)
		}
	}

	dwCol1 = colSeaTop.HWColor()
	dwCol2 := colSeaBtm.HWColor()

	for j := 0; j < SEA_SUBDIVISION; j++ {
		sea.SetColor(j, 0, dwCol1)
		sea.SetColor(j, SEA_SUBDIVISION-1, dwCol2)
	}

	var posX float64
	// Calculate light path
	if timet > 19.0 || timet < 5.0 { // moon
		a = 0.12 // intensity
		if timet > 19.0 && timet < 20.0 {
			a *= (timet - 19.0)
		} else if timet > 4.0 && timet < 5.0 {
			a *= 1.0 - (timet - 4.0)
		}
		posX = moonX
	} else if timet > 7.0 && timet < 17.0 { // sun
		a = 0.14 // intensity
		if timet < 8.0 {
			a *= (timet - 7.0)
		} else if timet > 16.0 {
			a *= 1.0 - (timet - 16.0)
		}
		posX = sunX
	} else {
		a = 0.0
	}

	if a != 0.0 {
		k := int(math.Floor(posX / float64(cellw)))
		s1 := (1.0 - (posX-float64(k*cellw))/float64(cellw))
		s2 := (1.0 - (float64((k+1)*cellw)-posX)/float64(cellw))

		if s1 > 0.7 {
			s1 = 0.7
		}
		if s2 > 0.7 {
			s2 = 0.7
		}

		s1 *= a
		s2 *= a

		for i := 0; i < SEA_SUBDIVISION; i += 2 {
			a = math.Sin(float64(i) / (SEA_SUBDIVISION - 1) * HGE.Pi_2)

			col1.SetHWColor(sea.Color(k, i))
			col1.AddEqual(colSun.MulScalar(s1 * (1 - a)))
			col1.Clamp()
			sea.SetColor(k, i, col1.HWColor())

			col1.SetHWColor(sea.Color(k+1, i))
			col1.AddEqual(colSun.MulScalar(s2 * (1 - a)))
			col1.Clamp()
			sea.SetColor(k+1, i, col1.HWColor())
		}
	}
}

func RenderSimulation() {
	// Render sky
	sky.SetColor(colSkyTop.HWColor(), 0)
	sky.SetColor(colSkyTop.HWColor(), 1)
	sky.SetColor(colSkyBtm.HWColor(), 2)
	sky.SetColor(colSkyBtm.HWColor(), 3)
	sky.Render(0, 0)

	// Render stars
	if seq_id >= 6 || seq_id < 2 {
		for i := 0; i < NUM_STARS; i++ {
			star.SetColor((HGE.Dword(starA[i]*255.0) << 24) | 0xFFFFFF)
			star.RenderEx(starX[i], starY[i], 0.0, starS[i])
		}
	}

	// Render sun
	glow.SetColor(colSunGlow.HWColor())
	glow.RenderEx(sunX, sunY, 0.0, sunGlowS)
	sun.SetColor(colSun.HWColor())
	sun.RenderEx(sunX, sunY, 0.0, sunS)

	// Render moon
	glow.SetColor(colMoonGlow.HWColor())
	glow.RenderEx(moonX, moonY, 0.0, moonGlowS)
	moon.SetColor(colMoon.HWColor())
	moon.RenderEx(moonX, moonY, 0.0, moonS)

	// Render sea
	sea.Render(0, SKY_HEIGHT)
	seaglow.SetColor(colSeaGlow.HWColor())
	seaglow.RenderEx(seaGlowX, SKY_HEIGHT, 0.0, seaGlowSX, seaGlowSY)
}
