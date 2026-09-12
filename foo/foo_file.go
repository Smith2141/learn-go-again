package foo

type secret string

// privateFoo — не экспортируемый тип
type privateFoo struct {
	Value string
	secret_val secret
}

// NewPrivateFoo — конструктор типа privateFoo
// Функция публичная, то есть может быть вызвана из других пакетов
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data", secret_val: "secret"}
}

// getter
func GetPrivateFooSecret(p privateFoo) secret {
	return p.secret_val
}
