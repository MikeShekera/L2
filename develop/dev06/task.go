package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intSlice []int

func (i *intSlice) String() string {
	return fmt.Sprint(*i)
}

func (i *intSlice) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		num, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return err
		}
		*i = append(*i, num)
	}
	return nil
}

func main() {
	var columnsHolder intSlice
	flag.Var(&columnsHolder, "f", "num of column")
	delimiter := flag.String("d", "", "new delimeter")
	separated := flag.Bool("s", false, "new output only separated")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}
	inputString := builder.String()
	stringsSl := strings.Split(strings.Trim(inputString, "\n"), "\n")
	for _, v := range stringsSl {
		if *separated && !strings.Contains(v, *delimiter) {
			continue
		}
		tempSl := strings.Split(v, *delimiter)
		for _, column := range columnsHolder {
			if column < 0 || column >= len(tempSl) {
				fmt.Printf("Invalid column index: %d \n", column)
				continue
			}
			fmt.Println(tempSl[column])
		}
	}
}
