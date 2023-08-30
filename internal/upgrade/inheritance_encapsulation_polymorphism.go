package upgrade

import "fmt"

// Базовый класс "Фигура"
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Класс "Прямоугольник", наследующий "Фигуру"
type rectangle struct {
	width  float64
	height float64
}

// Реализация методов "Area" и "Perimeter" для "Прямоугольника"
func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (r rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Класс "Круг", наследующий "Фигуру"
type Circle struct {
	radius float64
}

// Реализация методов "Area" и "Perimeter" для "Круга"
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func IEP() {
	// Создание объектов "Прямоугольник" и "Круг"
	r := rectangle{width: 4, height: 5}
	c := Circle{radius: 3}

	// Использование полиморфизма при вызове методов "Area" и "Perimeter"
	shapes := []Shape{r, c}

	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}
