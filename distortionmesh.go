package hge

const (
	DISP_NODE    = 0
	DISP_TOPLEFT = 1
	DISP_CENTER  = 2
)

type DistortionMesh struct {
	hge                   *HGE
	dispArray             []Vertex
	rows, cols            int
	cellw, cellh          float64
	tx, ty, width, height float64
	quad                  Quad
}

func NewDistortionMesh(cols, rows int) DistortionMesh {
	var dm DistortionMesh

	dm.hge = Create(VERSION)

	dm.rows = rows
	dm.cols = cols
	dm.cellw, dm.cellh = 0, 0
	dm.quad.Tex = 0
	dm.quad.Blend = BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_ZWRITE
	dm.dispArray = make([]Vertex, rows*cols)

	for i := 0; i < rows*cols; i++ {
		dm.dispArray[i].X = 0.0
		dm.dispArray[i].Y = 0.0
		dm.dispArray[i].TX = 0.0
		dm.dispArray[i].TY = 0.0

		dm.dispArray[i].Z = 0.5
		dm.dispArray[i].Col = 0xFFFFFFFF
	}

	return dm
}

func (dm *DistortionMesh) Render(x, y float64) {
	x32, y32 := float32(x), float32(y)

	for j := 0; j < dm.rows-1; j++ {
		for i := 0; i < dm.cols-1; i++ {
			idx := j*dm.cols + i

			dm.quad.V[0].TX = dm.dispArray[idx].TX
			dm.quad.V[0].TY = dm.dispArray[idx].TY
			dm.quad.V[0].X = x32 + dm.dispArray[idx].X
			dm.quad.V[0].Y = y32 + dm.dispArray[idx].Y
			dm.quad.V[0].Z = dm.dispArray[idx].Z
			dm.quad.V[0].Col = dm.dispArray[idx].Col

			dm.quad.V[1].TX = dm.dispArray[idx+1].TX
			dm.quad.V[1].TY = dm.dispArray[idx+1].TY
			dm.quad.V[1].X = x32 + dm.dispArray[idx+1].X
			dm.quad.V[1].Y = y32 + dm.dispArray[idx+1].Y
			dm.quad.V[1].Z = dm.dispArray[idx+1].Z
			dm.quad.V[1].Col = dm.dispArray[idx+1].Col

			dm.quad.V[2].TX = dm.dispArray[idx+dm.cols+1].TX
			dm.quad.V[2].TY = dm.dispArray[idx+dm.cols+1].TY
			dm.quad.V[2].X = x32 + dm.dispArray[idx+dm.cols+1].X
			dm.quad.V[2].Y = y32 + dm.dispArray[idx+dm.cols+1].Y
			dm.quad.V[2].Z = dm.dispArray[idx+dm.cols+1].Z
			dm.quad.V[2].Col = dm.dispArray[idx+dm.cols+1].Col

			dm.quad.V[3].TX = dm.dispArray[idx+dm.cols].TX
			dm.quad.V[3].TY = dm.dispArray[idx+dm.cols].TY
			dm.quad.V[3].X = x32 + dm.dispArray[idx+dm.cols].X
			dm.quad.V[3].Y = y32 + dm.dispArray[idx+dm.cols].Y
			dm.quad.V[3].Z = dm.dispArray[idx+dm.cols].Z
			dm.quad.V[3].Col = dm.dispArray[idx+dm.cols].Col

			dm.hge.Gfx_RenderQuad(&dm.quad)
		}
	}
}

//DWORD col=0xFFFFFFFF, float z=0.5f);
func (dm *DistortionMesh) Clear(a ...interface{}) {
	col := Dword(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		switch a[i].(type) {
		case float64:
			z = a[i].(float64)
		case float32:
			z = float64(a[i].(float32))
		case Dword:
			col = a[i].(Dword)
		case uint:
			col = Dword(a[i].(uint))
		}
	}

	cols := float64(dm.cols)

	for j := 0.0; j < float64(dm.rows); j++ {
		for i := 0.0; i < cols; i++ {
			dm.dispArray[int(j*cols+i)].X = float32(i * dm.cellw)
			dm.dispArray[int(j*cols+i)].Y = float32(j * dm.cellh)
			dm.dispArray[int(j*cols+i)].Col = col
			dm.dispArray[int(j*cols+i)].Z = float32(z)
		}

	}
}

func (dm *DistortionMesh) SetTexture(tex Texture) {
	dm.quad.Tex = tex
}

func (dm *DistortionMesh) SetTextureRect(x, y, w, h float64) {
	var tw, th float64

	dm.tx, dm.ty = x, y
	dm.width, dm.height = w, h

	if dm.quad.Tex != 0 {
		tw = float64(dm.hge.Texture_GetWidth(dm.quad.Tex))
		th = float64(dm.hge.Texture_GetHeight(dm.quad.Tex))
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
	dm.quad.Blend = blend
}

func (dm *DistortionMesh) SetZ(col, row int, z float64) {
	if row < dm.rows && col < dm.cols {
		dm.dispArray[row*dm.cols+col].Z = float32(z)
	}
}

func (dm *DistortionMesh) SetColor(col, row int, color Dword) {
	if row < dm.rows && col < dm.cols {
		dm.dispArray[row*dm.cols+col].Col = color
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

func (dm DistortionMesh) GetTexture() Texture {
	return dm.quad.Tex
}

func (dm DistortionMesh) GetTextureRect() (x, y, w, h float64) {
	return dm.tx, dm.ty, dm.width, dm.height
}

func (dm DistortionMesh) GetBlendMode() int {
	return dm.quad.Blend
}

func (dm DistortionMesh) GetZ(col, row int) float64 {
	if row < dm.rows && col < dm.cols {
		return float64(dm.dispArray[row*dm.cols+col].Z)
	}
	return 0.0
}

func (dm DistortionMesh) GetColor(col, row int) Dword {
	if row < dm.rows && col < dm.cols {
		return dm.dispArray[row*dm.cols+col].Col
	}

	return 0
}

func (dm DistortionMesh) GetDisplacement(col, row, ref int) (dx, dy float64) {
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

func (dm DistortionMesh) GetRows() int {
	return dm.rows
}

func (dm DistortionMesh) GetCols() int {
	return dm.cols
}
