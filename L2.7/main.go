package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Что выведет программа?
//Объяснить работу конвейера с использованием select.

// Программа выведет все числа от 1 до 8, но в разном порядке

//Программа создает два канала: в одном отправляются числа 1, 3, 5, 7, а во втором 2, 4, 6, 8.
//Каждое число приходит с небольшой случайной задержкой. Функция merge с помощью select объединяет оба канала в один,
//выбирая то значение, которое поступило раньше. Если какой-то канал закрылся, merge в select присваивает ему nil
//и перестает из него читать, чтобы больше не ждать данных. Когда оба канала завершаются, общий канал в merge также закрывается.
//В итоге программа выводит все числа от 1 до 8, но их порядок будет случайным при каждом запуске.

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	for v := range c {
		fmt.Print(v)
	}
}
