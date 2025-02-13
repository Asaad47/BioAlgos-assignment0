package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// queue must be at least longer than pattern
func findIfExactMatch(pattern, queue string) bool {
	matched := true
	for i := 0; i < len(pattern); i++ {
		if queue[i] != pattern[i] {
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
		if queue[i] != pattern[i] {
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
// both pattern and queue must start and end with the same characters to be considered a match
func findIfAdditionalChar(pattern, queue string) bool {
	if len(queue) < len(pattern)+1 {
		return false
	}
	if queue[0] != pattern[0] {
		return false
	}

	firstMismatch := -1
	out := true
	for i := 0; i < len(pattern); i++ {

		if firstMismatch == -1 {
			if queue[i] != pattern[i] {
				firstMismatch = i
			}
		} else {
			if queue[i] != pattern[i-1] {
				out = false
				break
			}
		}
	}
	if firstMismatch == -1 {
		return false
	} else {
		if queue[len(pattern)] != pattern[len(pattern)-1] {
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
			if queue[i] != pattern[i] {
				firstMismatch = i
			}
		} else {
			if queue[i-1] != pattern[i] {
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

func findNumMisMatches(seqFileName, patternFileName string, printHeaders bool) int {
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
			pattern += strings.ToLower(strings.TrimSpace(line))
		}
	}

	fmt.Println("Pattern:", pattern)
	fmt.Println("Length of pattern:", len(pattern))
	fmt.Println("Starting sequence file processing: ", seqFileName)
	fmt.Println("-----------------------------------")

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
	start := 0
	queue := ""
	headers := 0
	numSingleMismatches := 0
	numAdditionalCharMatches := 0
	numMissingCharMatches := 0
	foundExactMatch := false

	scanner = bufio.NewScanner(seqFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if printHeaders {
				fmt.Println("Header:", line)
			}
			start = 0
			queue = ""
			chromosome = line
			headers++
		} else {
			queue += strings.ToLower(strings.TrimSpace(line))

			for len(queue) >= len(pattern) {
				if !foundExactMatch && findIfMissingChar(pattern, queue) { // to avoid counting duplicates
					numMissingCharMatches++
					fmt.Printf("Missing character: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-2)
				}
				foundExactMatch = false

				if findIfExactMatch(pattern, queue) {
					numExactMatches++
					fmt.Printf("Exact match: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-1)
					foundExactMatch = true
				}
				if findIfSingleMismatch(pattern, queue) {
					numSingleMismatches++
					fmt.Printf("Single mismatch: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-1)
				}
				if findIfAdditionalChar(pattern, queue) {
					numAdditionalCharMatches++
					fmt.Printf("Additional character: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern))
				}

				queue = queue[1:]
				start++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading sequence file:", err)
	}

	fmt.Println("** Number of exact matches:", numExactMatches)
	fmt.Println("** Number of single mismatches:", numSingleMismatches)
	fmt.Println("** Number of additional character matches:", numAdditionalCharMatches)
	fmt.Println("** Number of missing character matches:", numMissingCharMatches)

	return numExactMatches + numSingleMismatches + numAdditionalCharMatches + numMissingCharMatches
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Run: go run prob2_naive.go <seq_file>")
		fmt.Println("<seq_file> = ", os.Args[1])

		aluFilePath := "DF000000002.fa"
		seqFilePath := os.Args[1]
		fmt.Println("*** Total number of matches: ", findNumMisMatches(seqFilePath, aluFilePath, false))
		return
	}
	fmt.Println("** Total mismatches: ", findNumMisMatches("test/ex_sequence.txt", "test/ex_pattern.txt", true)) // this is for testing

	t2tFilePath := "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"
	grch38FilePath := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna"
	aluFilePath := "DF000000002.fa"

	filePaths := []string{aluFilePath, t2tFilePath, grch38FilePath}
	fileMatchCounts := make(map[string]int)

	for _, seqFilePath := range filePaths {
		numMatches := findNumMisMatches(seqFilePath, aluFilePath, false)
		fileMatchCounts[seqFilePath] = numMatches
		fmt.Println("*** Sequence file:", seqFilePath)
		fmt.Println("*** Total number of matches: ", numMatches)
		fmt.Println("-----------------------------------")
	}

	fmt.Println("File match counts:", fileMatchCounts)
}
