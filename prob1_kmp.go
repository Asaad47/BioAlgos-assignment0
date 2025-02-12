package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func precomputeBorderArray(pattern string) []int {
	border := make([]int, len(pattern)+1)
	border[0] = -1
	border[1] = 0
	i := border[1]
	for j := 2; j <= len(pattern); j++ {
		for (i >= 0) && (pattern[i] != pattern[j-1]) {
			i = border[i]
		}
		i++
		border[j] = i
	}
	return border
}

func findExactMatchesKMP(seqFileName, patternFileName string) int {
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

	border := precomputeBorderArray(pattern)
	chromosome := ""
	numMatches := 0
	start := 0
	queue := ""

	scanner = bufio.NewScanner(seqFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// fmt.Println("Header:", line)
			start = 0
			chromosome = line
		} else {
			queue += strings.ToLower(strings.TrimSpace(line))
			i := 0
			j := 0
			for i <= len(queue)-len(pattern) {
				for queue[i+j] == pattern[j] {
					j++
					if j == len(pattern) {
						numMatches++
						fmt.Printf("Match: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start+i-j+1, start+i)
						break
					}
				}
				i += j - border[j]
				j = max(0, border[j])
			}
			queue = queue[i:]
			start += i
		}
	}
	return numMatches
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Run: go run prob1_kmp.go <seq_file>")
		fmt.Println("<seq_file> = ", os.Args[1])

		aluFilePath := "DF000000002.fa"
		seqFilePath := os.Args[1]
		fmt.Println("*** Total number of matches: ", findExactMatchesKMP(seqFilePath, aluFilePath))
		return
	}

	t2tFilePath := "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"
	grch38FilePath := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna"
	aluFilePath := "DF000000002.fa"

	filePaths := []string{aluFilePath, t2tFilePath, grch38FilePath}
	fileMatchCounts := make(map[string]int)

	for _, seqFilePath := range filePaths {
		numMatches := findExactMatchesKMP(seqFilePath, aluFilePath)
		fileMatchCounts[seqFilePath] = numMatches
		fmt.Println("*** Sequence file:", seqFilePath)
		fmt.Println("*** Total number of matches: ", numMatches)
	}

	fmt.Println("File match counts:", fileMatchCounts)
}
