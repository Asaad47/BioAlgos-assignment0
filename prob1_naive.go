package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findNumExactMatches(seqFileName, patternFileName string) int {
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

	numMatches := 0
	chromosome := ""
	start := 0
	queue := ""
	headers := 0

	scanner = bufio.NewScanner(seqFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			// fmt.Println("Header:", line)
			start = 0
			queue = ""
			chromosome = line
			headers++
		} else {
			queue += strings.ToLower(strings.TrimSpace(line))

			for len(queue) >= len(pattern) {
				matched := true
				for i := 0; i < len(pattern); i++ {
					if queue[i] != pattern[i] {
						matched = false
						break
					}
				}
				if matched {
					numMatches++
					fmt.Printf("Matched: (chromosome, start, end) = (%s, %d, %d)\n", chromosome, start, start+len(pattern)-1)
				}
				queue = queue[1:]
				start++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading sequence file:", err)
	}

	return numMatches
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Run: go run prob1_naive.go <seq_file>")
		fmt.Println("<seq_file> = ", os.Args[1])

		aluFilePath := "DF000000002.fa"
		seqFilePath := os.Args[1]
		fmt.Println("*** Total number of matches: ", findNumExactMatches(seqFilePath, aluFilePath))
		return
	}

	t2tFilePath := "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"
	// t2tFilePath2 := "ncbi_dataset_T2T/ncbi_dataset/data/GCF_009914755.1/GCF_009914755.1_T2T-CHM13v2.0_genomic.fna"
	grch38FilePath := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna"
	// grch38FilePath2 := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCF_000001405.40/GCF_000001405.40_GRCh38.p14_genomic.fna"
	aluFilePath := "DF000000002.fa"

	// fmt.Println("*** Total number of matches: ", findNumExactMatches(grch38FilePath, aluFilePath))

	filePaths := []string{aluFilePath, t2tFilePath, grch38FilePath}
	fileMatchCounts := make(map[string]int)

	for _, seqFilePath := range filePaths {
		numMatches := findNumExactMatches(seqFilePath, aluFilePath)
		fileMatchCounts[seqFilePath] = numMatches
		fmt.Println("*** Sequence file:", seqFilePath)
		fmt.Println("*** Total number of matches: ", numMatches)
	}

	fmt.Println("File match counts:", fileMatchCounts)
}
