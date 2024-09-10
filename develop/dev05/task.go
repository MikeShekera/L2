package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "print +N lines after match")
	before := flag.Int("B", 0, "print +N lines before match")
	context := flag.Bool("C", false, "(A+B) print +-N lines before match")
	count := flag.Bool("c", false, "Num of lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert Result Show")
	fixed := flag.Bool("F", false, "Full Match")
	lineNum := flag.Bool("n", false, "Print line number")
	flag.Parse()

	searchString := flag.Arg(1)
	path, _ := filepath.Abs(flag.Arg(0))
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("<Input file> <Pattern> [flags]")
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	stringData := string(data)

	if *ignoreCase {
		stringData = strings.ToLower(stringData)
		searchString = strings.ToLower(searchString)
	}

	stringsSlice := strings.Split(stringData, "\r\n")
	counter := 0
	for i, v := range stringsSlice {
		if !*fixed {
			if *invert != strings.Contains(v, searchString) {
				if *count {
					counter++
					continue
				}
				if *lineNum {
					fmt.Printf("%d ", i+1)
				}
				if *after != 0 || *before != 0 {
					if *context {
						printAround(stringsSlice, i-(*after+*before), i+(*after+*before))
					} else {
						printAround(stringsSlice, i-*before, i+*after)
					}
					continue
				}
				fmt.Println(v)
			}
		} else {
			if v == searchString {
				if *count {
					counter++
					continue
				}
				if *lineNum {
					fmt.Printf("%d ", i+1)
				}
				if *after != 0 || *before != 0 {
					if *context {
						printAround(stringsSlice, i-(*after+*before), i+(*after+*before))
					} else {
						printAround(stringsSlice, i-*before, i+*after)
					}
					continue
				}
				fmt.Println(v)
			}
		}
	}
	if *count {
		fmt.Println(counter)
	}
}

func printAround(sl []string, firstElem, lastElem int) {
	if firstElem < 0 {
		firstElem = 0
	}
	if lastElem >= len(sl) {
		lastElem = len(sl) - 1
	}
	for firstElem <= lastElem {
		fmt.Println(sl[firstElem])
		firstElem++
	}
	fmt.Println("\n")
}
