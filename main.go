package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var options []rune

func errLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseOptions() {
	lenArgs := len(os.Args)
	if lenArgs > 2 {
		for i:=0; i< lenArgs-1; i++ {
			if !strings.Contains(os.Args[i], "-") || strings.Count(os.Args[i], "-") > 1 {
				help()
				break
			}
		}

		for i:=0; i < lenArgs-1; i++ {
			options = append([]rune(strings.Trim(os.Args[i], "-")))
		}
	} else if lenArgs == 2 {
		countDash := strings.Count(os.Args[0], "-")
		if countDash == 1 {
			options = append([]rune(strings.Trim(os.Args[0], "-")))
		} else if countDash == 2 {
			fmt.Println("Handling double dashed arguments")
		} else {
			help()
		}
	} else {
		help()
	}
}
		

		

func getOpts(idx int) rune {
	if idx > len(options)-1 {
		return 0
	}
	return options[idx]

}

func countTotalBytes() int{
	return 0
}

func countTotalLines() int{
	return 0
}

func countTotalWords() int{
	return 0
}
func countTotalChars() int{
	return 0
}


func help() {
	fmt.Println("Help message")
}

func main() {
	var idx = 0
	parseOptions()

	var printLines, printWords, printChars, printBytes bool
	var totalLines, totalWords, totalChars, totalBytes int
	
	// fmt.Println(printLines, printChars, printWords, printBytes)
	for getOpts(idx) != 0 {
		switch getOpts(idx) {
		case 'c':
			printBytes = true
		case 'l':
			printLines = true
		case 'w':
			printLines = true
		case 'm':
			printChars = true
		default:
			help()
		}

		idx++
	}


	if printLines {
		totalLines = countTotalLines()
	}
	if printChars {
		totalChars = countTotalChars()
	}
	if printBytes {
		totalBytes = countTotalBytes()
	}
	if printWords {
		totalWords = countTotalWords()
	}

	fmt.Println(totalLines, totalChars, totalBytes, totalWords)
}
