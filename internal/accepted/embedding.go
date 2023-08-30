package accepted

import "fmt"

// встраивание
// Родительская структура
type Parent struct {
	parentProperty string
}

func (p *Parent) ParentMethod() {
	fmt.Println("Вызван ParentMethod()")
}

// Дочерняя структура, встраивающая родительскую структуру
type Child struct {
	Parent
	childProperty string
}

func (c *Child) ChildMethod() {
	fmt.Println("Вызван ChildMethod()")
}

func Embedding() {
	// Создаем экземпляр дочерней структуры
	child := &Child{
		Parent: Parent{
			parentProperty: "Значение родительского свойства",
		},
		childProperty: "Значение дочернего свойства",
	}

	// Обращаемся к свойствам и методам родительской и дочерней структур
	fmt.Println("Родительское свойство:", child.parentProperty)
	fmt.Println("Дочернее свойство:", child.childProperty)

	child.ParentMethod()
	child.ChildMethod()
	//shadowed доделать

}
