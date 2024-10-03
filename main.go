package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type wordMap struct {
	data map[[2]string][]string
}

func newWordMap() *wordMap {
	return &wordMap{data: make(map[[2]string][]string)}
}

func (w *wordMap) insert(first, second, word string) {
	key := [2]string{first, second}
	w.data[key] = append(w.data[key], word)
}

func (w *wordMap) get(first, second string) string {
	key := [2]string{first, second}
	return w.data[key][rand.Intn(len(w.data[key]))]
}

func (w *wordMap) getFirstUpper() (string, string) {
	first, second := "", ""
	for key := range w.data {
		if len(key[0]) > 0 && key[0][0] >= 'A' && key[0][0] <= 'Z' {
			first, second = key[0], key[1]
		}
	}
	return first, second
}

func getArgs() (string, int) {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <path> <n>\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("invalid arg '%s' for <n>: %v\n", os.Args[2], err)
		os.Exit(1)
	}
	return os.Args[1], n
}

func main() {
	path, n := getArgs()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	wordMap := newWordMap()

	var first, second string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			if len(first) > 0 && len(second) > 0 {
				wordMap.insert(first, second, word)
			}
			first, second = second, word
		}
	}

	first, second = wordMap.getFirstUpper()
	output := []string{first, second}
	for range n {
		word := wordMap.get(first, second)
		output = append(output, word)
		first, second = second, word
	}
	fmt.Println(strings.Join(output, " "))
}
