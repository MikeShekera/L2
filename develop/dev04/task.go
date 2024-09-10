package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	arr := []string{"Тяпка", "Пятак", "пятак", "Пятка", "а", "ав", "ва", "листок", "СЛиТок", "столИК"}
	fmt.Println(findAnagrams(arr))
}

func findAnagrams(arr []string) map[string][]string {
	m := make(map[string][]string)
	m2 := make(map[string][]string)
	for _, v := range arr {
		if utf8.RuneCountInString(v) < 2 {
			continue
		}
		var temp []rune
		v = strings.ToLower(v)
		temp = []rune(v)
		sort.Slice(
			temp, func(i, j int) bool {
				return temp[i] < temp[j]
			},
		)
		if slV, ok := m[string(temp)]; ok {
			if !contains(slV, v) {
				slV = append(slV, v)
				m[string(temp)] = slV
			}
		} else {
			m[string(temp)] = []string{v}
		}
	}
	for _, v := range m {
		m2[v[0]] = v[1:]
	}
	return m2
}

func contains(sl []string, searchingString string) bool {
	for _, innerValue := range sl {
		if innerValue == searchingString {
			return true
		}
	}
	return false
}
