package doc

/*
	Глобальный объект документа браузера
*/

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/size"
)

// TDoc -- операции с документом браузера
type TDoc struct {
	size *size.TSize
	doc  js.Value
}

var (
	doc *TDoc
)

// GetDoc -- возвращает (при необходимости создаёт) *TDoc
func GetDoc() *TDoc {
	if doc != nil {
		return doc
	}
	doc = &TDoc{
		size: size.New(),
		doc:  js.Global().Get("document"),
	}
	return doc
}

// GetElementByID -- возвращает элемент документа по его имени
func (sf *TDoc) GetElementByID(id string) js.Value {
	return sf.doc.Call("getElementById", id)
}

// Bind -- привязывает событие к обработчику JavaScript
func (sf *TDoc) Bind(event string, fnBackEvent js.Func) {
	sf.doc.Call("addEventListener", event, fnBackEvent)
}

// GetSize -- возвращае тфизические размеры тела документа
func (sf *TDoc) GetSize() (w, h float64) {
	sf.UpdateSize()
	return sf.size.Get()
}

// UpdateSize -- обновляет размер тела документа
func (sf *TDoc) UpdateSize() {
	w := sf.doc.Get("innerWidth").Float()
	h := sf.doc.Get("innerHeight").Float()
	sf.size.Set(w, h)
}
