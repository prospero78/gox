package posy

/*
	Файл предоставляет линейную координату точки по оси.
*/

// TPosY -- операции с положением точки по оси
type TPosY struct {
	val float64
}

// New -- возвращает новый *TPosY
func New() *TPosY {
	return &TPosY{}
}

// Set -- устанавливает значение
func (sf *TPosY) Set(val float64) {
	sf.val = val
}

// Val -- возвращает значение
func (sf *TPosY) Get() float64 {
	return sf.val
}
