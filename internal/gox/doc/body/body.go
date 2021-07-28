package body

/*
	Файл предоставляе ттип тела документа браузера
*/

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/size"
)

// TBody -- тело документа браузеа
type TBody struct {
	size *size.TSize
	body js.Value
}

var (
	body *TBody
)

// GetBody -- возвращает (при необходимости создаёт новый) *TBody
func GetBody() *TBody {
	if body != nil {
		return body
	}
	body = &TBody{
		size: size.New(),
		body: js.Global().Get("document").Get("body"),
	}
	body.UpdateSize()
	return body
}

// GetSize -- возвращае тфизические размеры тела документа
func (sf *TBody) GetSize() (w, h float64) {
	sf.UpdateSize()
	return sf.size.Get()
}

// SetSize -- устанавливает физические размеры тела документа
func (sf *TBody) SetSize(w, h float64) {
	sf.body.Set("clientWidth", w)
	sf.body.Set("clientHeight", h)
	sf.size.Set(w, h)
}

// UpdateSize -- обновляет размер тела документа
func (sf *TBody) UpdateSize() {
	w := sf.body.Get("clientWidth").Float()
	h := sf.body.Get("clientHeight").Float()
	sf.size.Set(w, h)
}
