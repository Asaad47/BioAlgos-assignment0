package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// queue must be at least longer than pattern
func findIfExactMatch(pattern, queue string) bool {
	matched := true
	for i := 0; i < len(pattern); i++ {
		if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i])) {
			matched = false
			break
		}
	}
	return matched
}

// finds pattern in sequence with a single character mismatch
func findIfSingleMismatch(pattern, queue string) bool {
	mismatchCount := 0
	for i := 0; i < len(pattern); i++ {
		if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i])) {
			mismatchCount++
			if mismatchCount > 1 {
				break
			}
		}
	}
	if mismatchCount == 1 {
		return true
	}
	return false
}

// finds pattern in sequence with an additional character, so length of sequence is len(pattern) + 1
func findIfAdditionalChar(pattern, queue string) bool {
	if len(queue) < len(pattern)+1 {
		return false
	}
	firstMismatch := -1
	out := true
	for i := 0; i < len(pattern); i++ {

		if firstMismatch == -1 {
			if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i])) {
				firstMismatch = i
			}
		} else {
			if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i-1])) {
				out = false
				break
			}
		}
	}
	if firstMismatch == -1 {
		return false
	} else {
		if unicode.ToLower(rune(queue[len(pattern)])) != unicode.ToLower(rune(pattern[len(pattern)-1])) {
			out = false
		}
		return out
	}
}

// finds pattern in sequence with a missing character, so length of sequence is len(pattern) - 1
func findIfMissingChar(pattern, queue string) bool {
	firstMismatch := -1
	out := true
	for i := 0; i < len(pattern); i++ {
		if firstMismatch == -1 {
			if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i])) {
				firstMismatch = i
			}
		} else {
			if unicode.ToLower(rune(queue[i-1])) != unicode.ToLower(rune(pattern[i])) {
				out = false
				break
			}
		}
	}
	if firstMismatch == -1 {
		return false
	}
	return out
}

func findNumMisMatches(seqFileName, patternFileName string) int {
	// Read the pattern from the pattern file
	pattern := ""
	patternFile, err := os.Open(patternFileName)
	if err != nil {
		fmt.Println("Error opening pattern file:", err)
		return 0
	}
	defer patternFile.Close()

	scanner := bufio.NewScanner(patternFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			fmt.Println("Pattern-- Header:", line)
		} else {
			pattern += strings.TrimSpace(line)
		}
	}

	fmt.Println("Pattern:", pattern)
	fmt.Println("Length of pattern:", len(pattern))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading pattern file:", err)
		return 0
	}

	// Read the sequence file
	seqFile, err := os.Open(seqFileName)
	if err != nil {
		fmt.Println("Error opening sequence file:", err)
		return 0
	}
	defer seqFile.Close()

	numExactMatches := 0
	chromosome := ""
	// validChromosome := false
	start := 0
	queue := ""
	headers := 0
	numSingleMismatches := 0
	numAdditionalCharMatches := 0
	numMissingCharMatches := 0

	scanner = bufio.NewScanner(seqFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			fmt.Println("Header:", line)
			// fmt.Println("-- Number of matches so far:", numMatches)
			start = 0
			queue = ""
			chromosomeParts := strings.Split(line, "chromosome ")
			if len(chromosomeParts) > 1 {
				chromosome = chromosomeParts[1]
				// validChromosome = true
			} else {
				chromosome = chromosomeParts[0]
				// validChromosome = false
			}
			headers++
		} else {
			// if !validChromosome {
			// 	continue
			// }
			// if headers > 1 {
			// 	break
			// }
			queue += strings.TrimSpace(line)
			// fmt.Println("line: ", line)

			for len(queue) >= len(pattern) {
				if findIfExactMatch(pattern, queue) {
					numExactMatches++
					fmt.Printf("Matched: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-1)
				}
				if findIfSingleMismatch(pattern, queue) {
					numSingleMismatches++
					fmt.Printf("Single mismatch: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-1)
				}
				if findIfAdditionalChar(pattern, queue) {
					numAdditionalCharMatches++
					fmt.Printf("Additional character: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)+1)
				}
				if findIfMissingChar(pattern, queue) {
					numMissingCharMatches++
					fmt.Printf("Missing character: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-2)
				}

				queue = queue[1:]
				start++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading sequence file:", err)
	}

	return numExactMatches + numSingleMismatches + numAdditionalCharMatches + numMissingCharMatches
}

func main() {
	fmt.Println("** Total mismatches: ", findNumMisMatches("ex_sequence.txt", "ex_pattern.txt"))
}
