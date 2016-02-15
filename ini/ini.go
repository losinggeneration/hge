package ini

import (
	"github.com/losinggeneration/hge"
)

type Ini struct {
	Section, Name string
	iniHGE        *hge.HGE
}

func New(section, name string) Ini {
	return Ini{}
}

func (i Ini) SetInt(value int) {
}

func (i Ini) GetInt(def_val int) int {
	return 0
}

func (i Ini) SetFloat(value float64) {
}

func (i Ini) GetFloat(def_val float64) float64 {
	return 0
}

func (i Ini) SetString(value string) {
}

func (i Ini) GetString(def_val string) string {
	return ""
}
