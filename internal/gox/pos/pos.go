package pos

import (
	"github.com/prospero78/gox/internal/gox/pos/posx"
	"github.com/prospero78/gox/internal/gox/pos/posy"
)

/*
	Файл предоставляет тип для хранения координат точки
*/

// TPos -- позиция точки
type TPos struct {
	posx *posx.TPosX
	posy *posy.TPosY
}

// New -- возвращает новый *TPos
func New() *TPos {
	return &TPos{
		posx: posx.New(),
		posy: posy.New(),
	}
}

// X - -возвращает объект Х точки
func (sf *TPos) X() *posx.TPosX {
	return sf.posx
}

// Y -- возвращает объект Y точки
func (sf *TPos) Y() *posy.TPosY {
	return sf.posy
}

// Set -- устанавливает положение точки
func (sf *TPos) Set(x, y float64) {
	sf.posx.Set(x)
	sf.posy.Set(y)
}

// Get -- возвращает положение точки
func (sf *TPos) Get() (x, y float64) {
	sf.posx.Set(x)
	sf.posy.Set(y)
	return x, y
}
