package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	a := `qwe\\5`
	newStr, err := unpackStringBuilder(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newStr)
}

func unpackStringBuilder(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	if unicode.IsDigit([]rune(str)[0]) {
		return "", fmt.Errorf("Некорректная строка")
	}
	var builder strings.Builder
	var isEscapeEnabled bool
	var savedChar rune
	casted := []rune(str)
	currentMultiplier := 0
	for i, v := range casted {
		if unicode.IsDigit(v) && !isEscapeEnabled {
			currentMultiplier *= 10
			val, _ := strconv.Atoi(string(v))
			currentMultiplier += val
			continue
		}
		if !isEscapeEnabled {
			if currentMultiplier != 0 {
				for range currentMultiplier {
					builder.WriteRune(savedChar)
				}
				currentMultiplier = 0
			} else {
				if i != 0 {
					builder.WriteRune(savedChar)
				}
			}
		}
		if casted[i] == '\\' && !isEscapeEnabled {
			isEscapeEnabled = true
			continue
		}
		isEscapeEnabled = false
		savedChar = v
	}
	if savedChar == 0 {
		return "", fmt.Errorf("Некорректная строка")
	}
	if currentMultiplier != 0 {
		for range currentMultiplier {
			builder.WriteRune(savedChar)
		}
	} else {
		builder.WriteRune(savedChar)
	}
	return builder.String(), nil
}

/*func unpackString(str string) (string, error) {
	targetLen, holder, ok := getAllMultipliedChars(str)
	if !ok {
		return "", fmt.Errorf("Некорректная строка")
	}
	casted := []rune(str)
	newStr := make([]rune, 0, targetLen)
	currentHolderIndex := 0
	i := 0
	for i < len(casted) {
		if !unicode.IsDigit(casted[i]) {
			newStr = append(newStr, casted[i])
			i++
			continue
		}

		for range holder[currentHolderIndex].num - 1 {
			newStr = append(newStr, holder[currentHolderIndex].multipliedChar)
		}
		i += countDigits(holder[currentHolderIndex].num)
		currentHolderIndex++
	}
	return string(newStr), nil
}*/
/*
func getAllMultipliedChars(str string) (int, []multiplierHolder, bool) {
	newStrBytesCount := 0
	var holder []multiplierHolder
	var charHolder rune
	currentMultiplier := 0

	for _, v := range str {
		if !unicode.IsDigit(v) {
			if currentMultiplier != 0 {
				holder = append(
					holder, multiplierHolder{
						num:            currentMultiplier,
						multipliedChar: charHolder,
					},
				)
				newStrBytesCount += currentMultiplier
				currentMultiplier = 0
			} else {
				newStrBytesCount++
			}
			charHolder = v
			continue
		}
		currentMultiplier *= 10
		val, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		currentMultiplier += val
	}
	if currentMultiplier != 0 {
		holder = append(
			holder, multiplierHolder{
				num:            currentMultiplier,
				multipliedChar: charHolder,
			},
		)
	}
	if charHolder == 0 {
		return 0, nil, false
	}
	return newStrBytesCount, holder, true
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(math.Abs(float64(n)))) + 1
}
*/
