package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scaner := bufio.NewScanner(inputFile)

	wordToDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var sumOfAll int64
	var countLines int
	for scaner.Scan() {
		countLines++

		inputLine := scaner.Text()
		digits := make([]string, 0)
		digitWordBuilder := strings.Builder{}
		for _, char := range inputLine {
			if !unicode.IsDigit(char) {
				digitWordBuilder.WriteRune(char)

				for word, digit := range wordToDigits {
					if !strings.Contains(digitWordBuilder.String(), word) {
						continue
					}

					digits = append(digits, digit)
					digitWordBuilder.Reset()
					digitWordBuilder.WriteRune(char)
				}

				continue
			}

			digits = append(digits, string(char))
			digitWordBuilder.Reset()
		}

		digitBuilder := strings.Builder{}
		if len(digits) == 0 {
			continue
		}

		if len(digits) == 1 {
			firstNumber := digits[0]
			digitBuilder.WriteString(firstNumber)
			digitBuilder.WriteString(firstNumber)
		}

		if len(digits) > 1 {
			firstNumber := digits[0]
			digitBuilder.WriteString(firstNumber)

			lastNumber := digits[len(digits)-1]
			digitBuilder.WriteString(lastNumber)
		}

		numberStr := digitBuilder.String()
		numberInt, err := strconv.ParseInt(numberStr, 0, 64)
		if err != nil {
			log.Fatal("got string that is not a number", numberStr, err)
		}

		sumOfAll += numberInt
	}

	fmt.Println("Total: ", sumOfAll)
}
