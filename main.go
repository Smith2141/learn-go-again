package main

import (
	"fmt"
	"my-go-app/foo"
)

func main() {
	// f := foo.privateFoo{} // ошибка компиляции
	f := foo.NewPrivateFoo()
	fmt.Println(f.Value) // поле Value экспортируемое, то есть его можно использовать
	// getter demo
	fmt.Println(foo.GetPrivateFooSecret(f))
}
