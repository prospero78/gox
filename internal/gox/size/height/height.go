package height

/*
	Предоставляет высоту экрана
*/

import (
	"log"
)

// THeight -- операции с высотой экрана
type THeight struct {
	val float64
}

// New -- возвращает новую высоту экрана*THeight
func New() *THeight {
	return &THeight{}
}

// Get -- возвращает хранимую высоту экрана
func (sf *THeight) Get() float64 {
	return sf.val
}

// Set -- устанавливает высоту экрана
func (sf *THeight) Set(val float64) {
	if val < 0 {
		log.Printf("THeight.Set(): val(%v)<0", val)
		return
	}
	sf.val = val
}
