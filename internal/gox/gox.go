package gox

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/canvas"
	"github.com/prospero78/gox/internal/gox/doc"
	"github.com/prospero78/gox/internal/gox/doc/body"
	"github.com/prospero78/gox/internal/gox/mouse"
	"github.com/prospero78/gox/internal/gox/size"
)

/*
	Графический сервер для показа результат рендеринга в окне браузера
*/

// TGox -- графический браузерный сервер
type TGox struct {
	size   *size.TSize     // Размеры экрана
	mouse  *mouse.TMouse   // Объект мыши
	doc    *doc.TDoc       // ОБъект документа браузера
	body   *body.TBody     // Тело документа
	canvas *canvas.TCanvas // Хост для рисования
	ctx    js.Value        // Контекст рисования
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
		size:   size.New(),
		doc:    doc.GetDoc(),
		mouse:  mouse.GetMouse(),
		body:   body.GetBody(),
		canvas: canvas.GetCanvas(),
	}
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

// Doc -- возвращает объект размера документа
func (sf *TGox) Doc() *doc.TDoc {
	return sf.doc
}

// Body -- возвращает объект размера документа
func (sf *TGox) Body() *body.TBody {
	return sf.body
}

// Canvas -- возвращает объект хоста
func (sf *TGox) Canvas() *canvas.TCanvas {
	return sf.canvas
}
