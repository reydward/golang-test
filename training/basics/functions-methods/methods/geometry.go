package methods

// Rectangle Definition de una structure llamada Rectángulo
type Rectangle struct {
	Width, Height int
}

// Area Método de la estructura Rectangle para calcular el área
func (r Rectangle) Area() int {
	return r.Width * r.Height
}
