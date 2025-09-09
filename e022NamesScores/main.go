/*

Using names.txt, a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical
value for each name, multiply this value by its alphabetical position in the
list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN,
which is worth $3 + 15 + 12 + 9 + 14 = 53$, is the $938$th name in the list.
So, COLIN would obtain a score of $938 \times 53 = 49714$.

What is the total of all the name scores in the file?

Answer:  871198282

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode/utf8"
)

// ScanCSV is a splitfunction for a scanner. It splits on the comma rune
func ScanCSV(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	// Scan until comma, marking end of commas.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func readNamesToArray(filename string) []string {
	var names []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(ScanCSV)

	for scanner.Scan() {
		n := scanner.Text()
		names = append(names, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(names); i++ {
		if names[i][0] == '"' {
			names[i] = names[i][1:]
		}
		if names[i][len(names[i])-1] == '"' {
			names[i] = names[i][:len(names[i])-1]
		}
	}
	return names
}

func runeValue(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 96
	}
	return int(r) - 64
}

func wordValue(word string) int {
	value := 0
	for _, r := range word {
		value += runeValue(r)
	}
	return value
}

func main() {
	names := readNamesToArray("names.txt")
	sort.Strings(names)

	sum := 0
	for k, n := range names {
		sum += wordValue(n) * (k + 1)

	}
	fmt.Println(sum)

}
