package main

import "fmt"

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

func NewItem(opts ...func(*Item)) *Item {
	// инициализируем типовыми значениями
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}
	// применяем опции в том порядке, в котором они были заявлены
	for _, opt := range opts {
		opt(i)
	}
	return i
}

func Option1(option1 string) func(*Item) {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}
func Option2(option2 int) func(*Item) {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

func main() {
	// с параметрами по умолчанию
	item1 := NewItem()
	// с применением одной опции
	item2 := NewItem(Option2(70))
	// или двух
	item3 := NewItem(Option1("unusual"), Option2(99))
	// опции можно заявлять в разном порядке
	item4 := NewItem(Option2(88), Option1("rare"))

	fmt.Println(item1.Parameter1)
	fmt.Println(item1.Parameter2)
	fmt.Println(item1.NoOption)
	fmt.Println("**********")

	fmt.Println(item2.Parameter1)
	fmt.Println(item2.Parameter2)
	fmt.Println(item2.NoOption)
	fmt.Println("**********")

	fmt.Println(item3.Parameter1)
	fmt.Println(item3.Parameter2)
	fmt.Println(item3.NoOption)
	fmt.Println("**********")

	fmt.Println(item4.Parameter1)
	fmt.Println(item4.Parameter2)
	fmt.Println(item4.NoOption)
}
