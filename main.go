package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(string(getFileName()))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	taboo := make(map[string]bool, 5)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		if _, ok := taboo[fileScanner.Text()]; !ok {
			taboo[fileScanner.Text()] = true
		}
	}

	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		wordSlice := strings.Split(scanner.Text(), " ")
		sb.Grow(len(wordSlice))

		for _, s := range wordSlice {
			if s == "exit" {
				fmt.Println("Bye!")
				return
			}

			if _, ok := taboo[strings.ToLower(s)]; !ok {
				sb.WriteString(s + " ")
			} else {
				s = strings.Repeat("*", len(s))
				sb.WriteString(strings.ToLower(s) + " ")
			}
		}

		fmt.Println(sb.String())
		sb.Reset()
	}
}

func getFileName() []byte {
	var fn []byte
	fmt.Scanf("%v", &fn)
	return fn
}
