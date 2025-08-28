package main

//Что выведет программа?
//Объяснить вывод программы.

// Программа выведет error

// Результат функция test() возвращает тип *customError, но сам return возвращает nil.
// В main присваиваем это значение переменной err типа error, а в Go error это интерфейс.
// Поскольку интерфейс хранит в себе два поля: тип и значение, то получается,
// что тип у нас будет *customError, а значение nil.
// Поэтому условие err != nil выполняется, и программа выведет "error".

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
