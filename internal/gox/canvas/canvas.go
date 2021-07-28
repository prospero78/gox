package canvas

/*
	Глобальный объект холста для рсования
*/

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/doc"
	"github.com/prospero78/gox/internal/gox/size"
)

// TCanvas -- глобальный хост для рисования
type TCanvas struct {
	canv js.Value
	doc  *doc.TDoc
	size *size.TSize
	ctx  js.Value // Контекст рисования
}

var (
	canv *TCanvas
)

// GetCanvas -- возвращает (при необходимости создаёт) *TCanvas
func GetCanvas() *TCanvas {
	if canv != nil {
		return canv
	}
	canv = &TCanvas{
		doc:  doc.GetDoc(),
		size: size.New(),
	}
	canv.canv = canv.doc.GetElementByID("mycanvas")
	canv.UpdateSize()
	canv.ctx = canv.canv.Call("getContext", "2d")
	return canv
}

// UpdateSize -- обновляет размер холста
func (sf *TCanvas) UpdateSize() {
	w := sf.canv.Get("width").Float()
	h := sf.canv.Get("height").Float()
	sf.size.Set(w, h)
}

// Size -- возвращает объект размера холста
func (sf *TCanvas) Size() *size.TSize {
	return sf.size
}

// SizeSet -- устанавливает размера холста
func (sf *TCanvas) SizeSet(w, h float64) {
	sf.canv.Set("width", w)
	sf.canv.Set("height", h)
	sf.size.Set(w, h)
}

// Ctx -- возвращает контекст рисования холста
func (sf *TCanvas) Ctx() js.Value {
	return sf.ctx
}
