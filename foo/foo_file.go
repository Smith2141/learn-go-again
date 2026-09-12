package foo

// privateFoo — не экспортируемый тип
type privateFoo struct {
	Value string
	secret_val string
}

// NewPrivateFoo — конструктор типа privateFoo
// Функция публичная, то есть может быть вызвана из других пакетов
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data", secret_val: "secret"}
}

// getter
func GetPrivateFooSecret(p privateFoo) string {
	return p.secret_val
}
