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

func NewDistortionMesh(cols, rows int) {
}

func (dm *DistortionMesh) Cmp(dm2 DistortionMesh) DistortionMesh {
	return *dm
}

func (dm *DistortionMesh) Render(x, y float64) {
}

func (dm *DistortionMesh) Clear(a ...interface{}) {
	//DWORD col=0xFFFFFFFF, float z=0.5f);
}

func (dm *DistortionMesh) SetTexture(tex Texture) {
}

func (dm *DistortionMesh) SetTextureRect(x, y, w, h float64) {
}

func (dm *DistortionMesh) SetBlendMode(blend int) {
}

func (dm *DistortionMesh) SetZ(col, row int, z float64) {
}

func (dm *DistortionMesh) SetColor(col, row int, color Dword) {
}

func (dm *DistortionMesh) SetDisplacement(col, row int, dx, dy float64, ref int) {
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
