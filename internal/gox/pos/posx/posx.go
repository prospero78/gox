package posx

/*
	Файл предоставляет линейную координату точки по оси.
*/

// TPosX -- операции с положением точки по оси
type TPosX struct {
	val float64
}

// New -- возвращает новый *TPosX
func New() *TPosX {
	return &TPosX{}
}

// Set -- устанавливает значение
func (sf *TPosX) Set(val float64) {
	sf.val = val
}

// Val -- возвращает значение
func (sf *TPosX) Get() float64 {
	return sf.val
}
