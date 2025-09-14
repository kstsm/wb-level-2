package main

import (
	"fmt"
	"sort"
	"strings"
)

// findAnagrams ищет все множества анаграмм в срезе слов
func findAnagrams(words []string) map[string][]string {
	// Сначала приводим все слова к нижнему регистру
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}

	// Создаем map для группировки слов по "ключу анаграммы"
	anagramGroups := make(map[string][]string)

	// Проходим по всем словам
	for _, w := range words {
		// Сортируем буквы в слове, чтобы получить ключ
		key := sortString(w)
		// Добавляем слово в соответствующую группу
		anagramGroups[key] = append(anagramGroups[key], w)
	}

	// Создаем итоговую мапу для результата
	result := make(map[string][]string)

	// Проходим по каждой группе слов
	for _, group := range anagramGroups {
		// Убираем дубликаты, чтобы одно и то же слово не повторялось
		unique := uniqueStrings(group)

		// Если группа содержит больше одного слова, то это уже настоящие анаграммы
		if len(unique) > 1 {
			// Сортируем слова внутри группы по алфавиту для красоты
			sort.Strings(unique)
			// В качестве ключа используем первое слово из группы
			result[unique[0]] = unique
		}
	}

	return result
}

// sortString сортирует буквы в строке и возвращает новую строку
func sortString(s string) string {
	// Преобразуем строку в срез рун, чтобы корректно работать с unicode
	runes := []rune(s)
	// Сортируем срез рун по возрастанию
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Преобразуем обратно в строку и возвращаем
	return string(runes)
}

// uniqueStrings убирает дубликаты из среза строк
func uniqueStrings(slice []string) []string {
	// Мапа для отслеживания уникальных слов
	seen := make(map[string]struct{})
	var res []string

	for _, s := range slice {
		// Если слово еще не встречалось, добавляем его в результат
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			res = append(res, s)
		}
	}

	return res
}

func main() {
	// Выходные данные
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	// Вызываем функцию поиска анаграмм
	result := findAnagrams(input)

	// Выводим результат
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
}
