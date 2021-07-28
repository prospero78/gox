package size

import (
	"github.com/prospero78/gox/internal/gox/size/height"
	"github.com/prospero78/gox/internal/gox/size/width"
)

/*
	Файл предоставляет размер экрана.
*/

// TSize -- операции с размером
type TSize struct {
	width  *width.TWidth
	height *height.THeight
}

// New -- возвращает новый *TSize
func New() *TSize {
	return &TSize{
		width:  width.New(),
		height: height.New(),
	}
}

// Weidth -- возвращает объект высоты экрана
func (sf *TSize) Width() *width.TWidth {
	return sf.width
}

// Height -- возвращает объект ширины экрана
func (sf *TSize) Height() *height.THeight {
	return sf.height
}

// Set -- устанавливает размеры экрана
func (sf *TSize) Set(w, h float64) {
	sf.width.Set(w)
	sf.height.Set(h)
}

// Get -- возвращает размеры экрана
func (sf *TSize) Get() (w, h float64) {
	w = sf.width.Get()
	h = sf.height.Get()
	return w, h
}
