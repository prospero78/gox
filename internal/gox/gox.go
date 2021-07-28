package gox

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/mouse"
	"github.com/prospero78/gox/internal/gox/size"
)

/*
	Графический сервер для показа результат рендеринга в окне браузера
*/

// TGox -- графический браузерный сервер
type TGox struct {
	size   *size.TSize   // Размеры экрана
	mouse  *mouse.TMouse // Объект мыши
	doc    js.Value
	canvas js.Value
	ctx    js.Value // Контекст рисования
}

var (
	gox *TGox
)

// GetGox -- возвращает (если надо создаёт новый) *TGox
func GetGox() *TGox {
	if gox != nil {
		return gox
	}
	gox = &TGox{
		size:  size.New(),
		doc:   js.Global().Get("document"),
		mouse: mouse.GetMouse(),
	}
	gox.canvas = gox.doc.Call("getElementById", "mycanvas")
	w := gox.doc.Get("body").Get("clientWidth").Float()
	h := gox.doc.Get("body").Get("clientHeight").Float()
	gox.canvas.Set("width", w)
	gox.canvas.Set("height", h)
	gox.ctx = gox.canvas.Call("getContext", "2d")
	// gox.size.Set(w, h)
	return gox
}

// Size -- возвращает объект размер экрана
func (sf *TGox) Size() *size.TSize {
	return sf.size
}

// Mouse -- возвращает объект мыши
func (sf *TGox) Mouse() *mouse.TMouse {
	return sf.mouse
}

// Close -- обязательный вызов в конце работы
func (sf *TGox) Close() {
	sf.mouse.Close()
}
