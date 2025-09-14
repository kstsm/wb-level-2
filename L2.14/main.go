package main

import (
	"fmt"
	"time"
)

// or объединяет несколько done-каналов в один.
// Как только закрывается хотя бы один из входных каналов
// возвращаемый канал тоже закрывается.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	// Если нет каналов, то возвращаем nil
	case 0:
		return nil
	// Если только один канал, то возвращаем его напрямую
	case 1:
		return channels[0]
	}

	// Канал для результата
	orDone := make(chan interface{})

	go func() {
		// Закрываем его при выходе из горутины
		defer close(orDone)

		switch len(channels) {
		// Если всего два канала, то ждём любой из них через select
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		// Если каналов больше, то используем рекурсию
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

// sig возвращает канал, который закроется через заданное время
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	// Запускаем несколько каналов с разными тайм-аутами
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	// Результат появится примерно через 1 секунду,
	// т.к. самый быстрый канал sig(1*time.Second) закроется первым
	fmt.Printf("done after %v\n", time.Since(start))
}
