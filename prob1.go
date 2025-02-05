package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
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
			pattern += strings.TrimSpace(line)
		}
	}

	// fmt.Println("Pattern:", pattern)
	// fmt.Println("Length of pattern:", len(pattern))

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
	// validChromosome := false
	start := 0
	queue := ""
	headers := 0

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
				matched := true
				for i := 0; i < len(pattern); i++ {
					if unicode.ToLower(rune(queue[i])) != unicode.ToLower(rune(pattern[i])) {
						matched = false
						break
					}
					// if i > 10 {
					// 	fmt.Println("reached i: ", i)
					// }
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
	t2tFilePath := "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"
	t2tFilePath2 := "ncbi_dataset_T2T/ncbi_dataset/data/GCF_009914755.1/GCF_009914755.1_T2T-CHM13v2.0_genomic.fna"
	grch38FilePath := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna"
	grch38FilePath2 := "ncbi_dataset_GRCh38/ncbi_dataset/data/GCF_000001405.40/GCF_000001405.40_GRCh38.p14_genomic.fna"
	aluFilePath := "DF000000002.fa"

	filePaths := []string{aluFilePath, t2tFilePath, t2tFilePath2, grch38FilePath, grch38FilePath2}

	// fmt.Println("*** Total number of matches: ", findNumExactMatches(t2tFilePath, aluFilePath))
	fileMatchCounts := make(map[string]int)

	for _, seqFilePath := range filePaths {
		numMatches := findNumExactMatches(seqFilePath, aluFilePath)
		fileMatchCounts[seqFilePath] = numMatches
		fmt.Println("*** Sequence file:", seqFilePath)
		fmt.Println("*** Total number of matches: ", numMatches)
	}

	fmt.Println("File match counts:", fileMatchCounts)
}
