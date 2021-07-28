package mouse

/*
	Файл предоставляет тип для работы с мышью
*/

import (
	"syscall/js"

	"github.com/prospero78/gox/internal/gox/pos"
)

// TMouse -- операции с мышью
type TMouse struct {
	pos      *pos.TPos
	callMove js.FuncOf
}

var (
	mouse *TMouse
)

// GetMouse -- возвращает (при нобходимости создаёт новый)*TMouse
func GetMouse() *TMouse {
	if mouse != nil {
		return mouse
	}
	mouse = &TMouse{
		pos: pos.New(),
	}
	mouse.callMove = js.FuncOf(mouse.mouseMove)
	return mouse
}

// Обработчик перемещения мыши
func (sf *TMouse) mouseMove(this js.Value, args []js.Value) interface{} {
	e := args[0]
	mX := e.Get("clientX").Float()
	mY := e.Get("clientY").Float()
	sf.pos.Set(mX, mY)
	return nil
}

// Pos -- возвращает объект координат мыши
func (sf *TMouse) Pos() *pos.TPos {
	return sf.pos
}

// PosGet -- возвращает хранимые координаты мыши
func (sf *TMouse) PosGet() (x, y float64) {
	return sf.pos.Get()
}

// PosSet -- принудительно устанавливает положение мыши на экране
func (sf *TMouse) PosSet(x, y float64) {
	sf.pos.Set(x, y)
}

// Close -- обязательный вызов в конце работы
func (sf *TMouse) Close() {
	sf.callMove.Release()
}
