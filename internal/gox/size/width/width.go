package width

/*
	Предоставляет ширину экрана
*/

import (
	"log"
)

// TWidth - -операци ис шириной экрана
type TWidth struct {
	val float64
}

// New -- возвращает новую ширину экрана*TWidth
func New() *TWidth {
	return &TWidth{}
}

// Get -- возвращает хранимую ширину экрана
func (sf *TWidth) Get() float64 {
	return sf.val
}

// Set -- устанавливает ширину экрана
func (sf *TWidth) Set(val float64) {
	if val < 0 {
		log.Printf("TWidth.Set(): val(%v)<0", val)
		return
	}
	sf.val = val
}
