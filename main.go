package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type wordMap struct {
	data map[[2]string][]string
}

func newWordMap() *wordMap {
	return &wordMap{data: make(map[[2]string][]string)}
}

func (w *wordMap) populate(scanner *bufio.Scanner) {
	var first, second string
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = strings.TrimSpace(word)
			w.insert(first, second, word)
			first, second = second, word
		}
	}
	w.insert(first, second, "")
	w.insert(second, "", "")
}

func (w *wordMap) insert(first, second, word string) {
	key := [2]string{first, second}
	w.data[key] = append(w.data[key], word)
}

func (w *wordMap) get(first, second string) string {
	key := [2]string{first, second}
	choices := w.data[key]
	if len(choices) == 0 {
		return ""
	}
	return choices[rand.Intn(len(choices))]
}

func (w *wordMap) getRandomUpper() (string, string) {
	choices := [][2]string{}
	for key := range w.data {
		if len(key[0]) > 0 && key[0][0] >= 'A' && key[0][0] <= 'Z' {
			choices = append(choices, key)
		}
	}
	if len(choices) == 0 {
		return "", ""
	}
	start := choices[rand.Intn(len(choices))]
	return start[0], start[1]
}

func getArgs() (string, int, bool) {
	path := flag.String("path", "", "path to the text file")
	n := flag.Int("n", 0, "the number of words to output")
	flag.Parse()
	if len(*path) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return *path, *n, *n == 0
}

func isTerminatingWord(word *string) bool {
	return strings.HasSuffix(*word, ".") ||
		strings.HasSuffix(*word, "?") ||
		strings.HasSuffix(*word, "!")
}

func main() {
	path, n, stopAtSentenceEnd := getArgs()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordMap := newWordMap()
	wordMap.populate(scanner)

	first, second := wordMap.getRandomUpper()
	output := []string{first, second}
	for stopAtSentenceEnd && !isTerminatingWord(&second) || len(output) < n {
		word := wordMap.get(first, second)
		if len(word) > 0 {
			output = append(output, word)
		}
		first, second = second, word
	}
	fmt.Println(strings.Join(output, " "))
}
