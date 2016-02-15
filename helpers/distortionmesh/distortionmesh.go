package distortionmesh

import (
	hge "github.com/losinggeneration/hge-go"
	"github.com/losinggeneration/hge-go/gfx"
)

const (
	DISP_NODE    = 0
	DISP_TOPLEFT = 1
	DISP_CENTER  = 2
)

type DistortionMesh struct {
	dispArray             []gfx.Vertex
	rows, cols            int
	cellw, cellh          float64
	tx, ty, width, height float64
	gfx.Quad
}

func New(cols, rows int) DistortionMesh {
	var dm DistortionMesh

	dm.rows = rows
	dm.cols = cols
	dm.Blend = gfx.BLEND_COLORMUL | gfx.BLEND_ALPHABLEND | gfx.BLEND_ZWRITE
	dm.dispArray = make([]gfx.Vertex, rows*cols)

	for i := 0; i < rows*cols; i++ {
		dm.dispArray[i].Z = 0.5
		dm.dispArray[i].Color = 0xFFFFFFFF
	}

	return dm
}

func (dm *DistortionMesh) Render(x, y float64) {
	x32, y32 := float32(x), float32(y)

	for j := 0; j < dm.rows-1; j++ {
		for i := 0; i < dm.cols-1; i++ {
			idx := j*dm.cols + i

			dm.V[0].TX = dm.dispArray[idx].TX
			dm.V[0].TY = dm.dispArray[idx].TY
			dm.V[0].X = x32 + dm.dispArray[idx].X
			dm.V[0].Y = y32 + dm.dispArray[idx].Y
			dm.V[0].Z = dm.dispArray[idx].Z
			dm.V[0].Color = dm.dispArray[idx].Color

			dm.V[1].TX = dm.dispArray[idx+1].TX
			dm.V[1].TY = dm.dispArray[idx+1].TY
			dm.V[1].X = x32 + dm.dispArray[idx+1].X
			dm.V[1].Y = y32 + dm.dispArray[idx+1].Y
			dm.V[1].Z = dm.dispArray[idx+1].Z
			dm.V[1].Color = dm.dispArray[idx+1].Color

			dm.V[2].TX = dm.dispArray[idx+dm.cols+1].TX
			dm.V[2].TY = dm.dispArray[idx+dm.cols+1].TY
			dm.V[2].X = x32 + dm.dispArray[idx+dm.cols+1].X
			dm.V[2].Y = y32 + dm.dispArray[idx+dm.cols+1].Y
			dm.V[2].Z = dm.dispArray[idx+dm.cols+1].Z
			dm.V[2].Color = dm.dispArray[idx+dm.cols+1].Color

			dm.V[3].TX = dm.dispArray[idx+dm.cols].TX
			dm.V[3].TY = dm.dispArray[idx+dm.cols].TY
			dm.V[3].X = x32 + dm.dispArray[idx+dm.cols].X
			dm.V[3].Y = y32 + dm.dispArray[idx+dm.cols].Y
			dm.V[3].Z = dm.dispArray[idx+dm.cols].Z
			dm.V[3].Color = dm.dispArray[idx+dm.cols].Color

			dm.Quad.Render()
		}
	}
}

//DWORD col=0xFFFFFFFF, float z=0.5f);
func (dm *DistortionMesh) Clear(a ...interface{}) {
	col := hge.Dword(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		switch a[i].(type) {
		case float64:
			z = a[i].(float64)
		case float32:
			z = float64(a[i].(float32))
		case hge.Dword:
			col = a[i].(hge.Dword)
		case uint:
			col = hge.Dword(a[i].(uint))
		}
	}

	cols := float64(dm.cols)

	for j := 0.0; j < float64(dm.rows); j++ {
		for i := 0.0; i < cols; i++ {
			dm.dispArray[int(j*cols+i)].X = float32(i * dm.cellw)
			dm.dispArray[int(j*cols+i)].Y = float32(j * dm.cellh)
			dm.dispArray[int(j*cols+i)].Color = col
			dm.dispArray[int(j*cols+i)].Z = float32(z)
		}

	}
}

func (dm *DistortionMesh) SetTexture(tex *gfx.Texture) {
	dm.Quad.Texture = tex
}

func (dm *DistortionMesh) SetTextureRect(x, y, w, h float64) {
	var tw, th float64

	dm.tx, dm.ty = x, y
	dm.width, dm.height = w, h

	if dm.Quad.Texture != nil {
		tw = float64(dm.Quad.Texture.Width())
		th = float64(dm.Quad.Texture.Height())
	} else {
		tw = w
		th = h
	}

	dm.cellw = w / float64(dm.cols-1)
	dm.cellh = h / float64(dm.rows-1)

	cols := float64(dm.cols)

	for j := 0.0; j < float64(dm.rows); j++ {
		for i := 0.0; i < float64(dm.cols); i++ {
			dm.dispArray[int(j*cols+i)].TX = float32((x + i*dm.cellw) / tw)
			dm.dispArray[int(j*cols+i)].TY = float32((y + j*dm.cellh) / th)

			dm.dispArray[int(j*cols+i)].X = float32(i * dm.cellw)
			dm.dispArray[int(j*cols+i)].Y = float32(j * dm.cellh)
		}
	}
}

func (dm *DistortionMesh) SetBlendMode(blend int) {
	dm.Blend = blend
}

func (dm *DistortionMesh) SetZ(col, row int, z float64) {
	if row < dm.rows && col < dm.cols {
		dm.dispArray[row*dm.cols+col].Z = float32(z)
	}
}

func (dm *DistortionMesh) SetColor(col, row int, color hge.Dword) {
	if row < dm.rows && col < dm.cols {
		dm.dispArray[row*dm.cols+col].Color = color
	}
}

func (dm *DistortionMesh) SetDisplacement(col, row int, dx, dy float64, ref int) {
	if row < dm.rows && col < dm.cols {
		switch ref {
		case DISP_NODE:
			dx += float64(col) * dm.cellw
			dy += float64(row) * dm.cellh
		case DISP_CENTER:
			dx += dm.cellw * float64(dm.cols-1) / 2
			dy += dm.cellh * float64(dm.rows-1) / 2
		case DISP_TOPLEFT:
		}

		dm.dispArray[row*dm.cols+col].X = float32(dx)
		dm.dispArray[row*dm.cols+col].Y = float32(dy)
	}
}

func (dm DistortionMesh) Texture() *gfx.Texture {
	return dm.Quad.Texture
}

func (dm DistortionMesh) TextureRect() (x, y, w, h float64) {
	return dm.tx, dm.ty, dm.width, dm.height
}

func (dm DistortionMesh) BlendMode() int {
	return dm.Blend
}

func (dm DistortionMesh) Z(col, row int) float64 {
	if row < dm.rows && col < dm.cols {
		return float64(dm.dispArray[row*dm.cols+col].Z)
	}
	return 0.0
}

func (dm DistortionMesh) Color(col, row int) hge.Dword {
	if row < dm.rows && col < dm.cols {
		return dm.dispArray[row*dm.cols+col].Color
	}

	return 0
}

func (dm DistortionMesh) Displacement(col, row, ref int) (dx, dy float64) {
	if row < dm.rows && col < dm.cols {
		switch ref {
		case DISP_NODE:
			dx = float64(dm.dispArray[row*dm.cols+col].X) - float64(col)*dm.cellw
			dy = float64(dm.dispArray[row*dm.cols+col].Y) - float64(row)*dm.cellh

		case DISP_CENTER:
			dx = float64(dm.dispArray[row*dm.cols+col].X) - dm.cellw*float64(dm.cols-1)/2
			dy = float64(dm.dispArray[row*dm.cols+col].X) - dm.cellh*float64(dm.rows-1)/2

		case DISP_TOPLEFT:
			dx = float64(dm.dispArray[row*dm.cols+col].X)
			dy = float64(dm.dispArray[row*dm.cols+col].Y)
		}
	}

	return dx, dy
}

func (dm DistortionMesh) Rows() int {
	return dm.rows
}

func (dm DistortionMesh) Cols() int {
	return dm.cols
}
