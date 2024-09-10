package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	outputFileName = "result.txt"
)

func main() {
	column := flag.Int("k", 0, "Sort column")
	numeric := flag.Bool("n", false, "Numeric sort")
	reverse := flag.Bool("r", false, "Reverse sort")
	unique := flag.Bool("u", false, "Remove repeating strings")
	flag.Parse()

	path, _ := filepath.Abs(flag.Arg(0))
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("<Input file> [flags]")
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	stringsSlice := strings.Split(string(data), "\r\n")
	if !*numeric {
		slices.Sort(stringsSlice)
	} else {
		sort.Slice(
			stringsSlice, func(i, j int) bool {
				switch strings.Compare(string(stringsSlice[i][*column]), string(stringsSlice[j][*column])) {
				case -1:
					return true
				case 1:
					return false
				}
				return stringsSlice[i] > stringsSlice[j]
			},
		)
	}

	sort.Slice(
		stringsSlice, func(i, j int) bool {
			if *numeric {
				numI, errI := strconv.ParseFloat(strconv.Itoa(int(stringsSlice[i][*column])), 64)
				numJ, errJ := strconv.ParseFloat(strconv.Itoa(int(stringsSlice[j][*column])), 64)
				if errI == nil && errJ == nil {
					return numI < numJ
				} else {
					return stringsSlice[i][*column] < stringsSlice[j][*column]
				}
			} else {
				return stringsSlice[i][*column] < stringsSlice[j][*column]
			}
		},
	)
	fmt.Println(stringsSlice)

	newFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	writer := bufio.NewWriter(newFile)
	defer writer.Flush()

	if *unique {
		stringsSlice = removeDuplicateStr(stringsSlice)
	}
	fmt.Println(*column)

	if *reverse {
		for i := len(stringsSlice) - 1; i >= 0; i-- {
			_, err = writer.WriteString(stringsSlice[i] + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		for _, v := range stringsSlice {
			_, err = writer.WriteString(v + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, err = os.Open(newFile.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
