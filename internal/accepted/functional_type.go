package accepted

import "fmt"

// функциональный тип
// Определяем интерфейс с методом
type Greeter interface {
	Greet()
}

// Определяем тип, который реализует метод интерфейса
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Привет, меня зовут %s!\n", p.Name)
}

func (p Person) Greet2() func(name string) {
	return func(name string) {
		fmt.Printf("Привет, меня зовут %s!\n", name)
	}
}

func FunctionalType() {
	var g Greeter
	// Создаем экземпляр типа, реализующего метод интерфейса
	p := Person{Name: "Иван"}

	// Присваиваем функциональное значение переменной интерфейса
	g = p

	g.Greet()
}
