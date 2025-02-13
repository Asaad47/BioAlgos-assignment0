# Introduction

This is the first assignment of CS 249 Algorithms in Bioinformatics. The assignment is mainly composed of two implementation parts: exact pattern matching and matching up to one mismatch. In both parts, the goal is to find the pattern sequence AluY in the sequence files of human genome assemblies GRCh38 (hg38) and T2T CHM13v2.0 retrieved from NCBI datasets. The AluY sequence is provided in the assignment description and available in the [FASTA](./DF000000002.fa) file. I used naive character matching and the Knuth-Morris-Pratt (KMP) algorithm for exact pattern matching and implemented single mismatch, single insertion, and single deletion for matching up to one mismatch. The results are reported in the [Results](#results-on-aluy) section.

## Brief implementation description

I have three main files that address the implementation problems: prob1_naive.go, prob1_kmp.go, and prob2.go. In all three files, I read the sequence file line by line and start searching for the pattern in the sequence whenever I have enough characters in the `queue` to match the pattern. Essentially, the `queue` is a buffer that holds read characters from the sequence file and searching happens whenever the `queue` is at least the length of the pattern. This is to avoid loading the entire sequence file into memory.

For finding exact matches, I used two implementations for searching: naive character matching and KMP algorithm found in prob1_naive.go and prob1_kmp.go respectively. The naive matching directly compares the characters in the `queue` with the pattern in each iteration while the KMP algorithm pre-computes the border array and uses it to skip unnecessary comparisons. However, both implementations achieved roughly the same performance.

For finding approximate matches, I seperated the search mechanism to four functions: exact match, single mismatch, single insertion, and single deletion. The exact match is the same as the naive implementation of prob1. The single mismatch resovles to true of the Hamming distance is exactly 1. The single insertion and single deletion look for the first mismatch index and then skip the mismatched character in the `queue` and the pattern respectively. Some additional checks are used to avoid counting duplicates. For example, if there a was an exact match, the single deletion function will not be called as it is a subset of the exact match. Also, the single addition function asserts that the first and last charcaters are the same as the pattern and that the mismatch is in the middle to avoid duplicates.

## Results on AluY

Dropdowns have more information on the results on specified files. The start and end indices are inclusive.

<details>
<summary>Number of exact matches in T2T file: 2</summary>
<br>

```bash
Run: go run prob1_naive.go <seq_file>
<seq_file> =  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Matched: (chromosome, start, end) = (>CP068276.2 Homo sapiens isolate CHM13 chromosome 2, 186833277, 186833587)
Matched: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 39966041, 39966351)
*** Total number of matches:  2
       17.78 real        16.89 user         1.02 sys
            55427072  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               15700  page reclaims
                1826  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               22927  signals received
                1170  voluntary context switches
              112810  involuntary context switches
           714061707  instructions retired
           269071124  cycles elapsed
            13616768  peak memory footprint
```

</details>

<details>
<summary>Number of exact matches in GRCh38 file: 3</summary>
<br>

```bash
Run: go run prob1_naive.go <seq_file>
<seq_file> =  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Matched: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629832, 88630142)
Matched: (chromosome, start, end) = (>CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly, 39808489, 39808799)
Matched: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 18689763, 18690073)
*** Total number of matches:  3
       19.21 real        18.27 user         1.11 sys
            56164352  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               16266  page reclaims
                2660  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               24626  signals received
                2557  voluntary context switches
              129995  involuntary context switches
           757152512  instructions retired
           282639290  cycles elapsed
            13338112  peak memory footprint
```

</details>

<details>
<summary>Number of matches up to one mismatch in T2T file: 21</summary>
<br>

```bash
Run: go run prob2_naive.go <seq_file>
<seq_file> =  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Pattern: ggccgggcgcggtggctcacgcctgtaatcccagcactttgggaggccgaggcgggcggatcacgaggtcaggagatcgagaccatcctggctaacacggtgaaaccccgtctctactaaaaatacaaaaaattagccgggcgtggtggcgggcgcctgtagtcccagctactcgggaggctgaggcaggagaatggcgtgaacccgggaggcggagcttgcagtgagccgagatcgcgccactgcactccagcctgggcgacagagcgagactccgtctcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Length of pattern: 311
Starting sequence file processing:  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
-----------------------------------
Single mismatch: (chromosome, start, end) = (>CP068277.2 Homo sapiens isolate CHM13 chromosome 1, 143946938, 143947248)
Single mismatch: (chromosome, start, end) = (>CP068277.2 Homo sapiens isolate CHM13 chromosome 1, 190858894, 190859204)
Exact match: (chromosome, start, end) = (>CP068276.2 Homo sapiens isolate CHM13 chromosome 2, 186833277, 186833587)
Single mismatch: (chromosome, start, end) = (>CP068275.2 Homo sapiens isolate CHM13 chromosome 3, 69350794, 69351104)
Single mismatch: (chromosome, start, end) = (>CP068274.2 Homo sapiens isolate CHM13 chromosome 4, 136221657, 136221967)
Single mismatch: (chromosome, start, end) = (>CP068273.2 Homo sapiens isolate CHM13 chromosome 5, 126911808, 126912118)
Single mismatch: (chromosome, start, end) = (>CP068272.2 Homo sapiens isolate CHM13 chromosome 6, 51234801, 51235111)
Single mismatch: (chromosome, start, end) = (>CP068272.2 Homo sapiens isolate CHM13 chromosome 6, 68212687, 68212997)
Exact match: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 39966041, 39966351)
Single mismatch: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 49189460, 49189770)
Single mismatch: (chromosome, start, end) = (>CP068270.2 Homo sapiens isolate CHM13 chromosome 8, 101916771, 101917081)
Single mismatch: (chromosome, start, end) = (>CP068267.2 Homo sapiens isolate CHM13 chromosome 11, 86916177, 86916487)
Single mismatch: (chromosome, start, end) = (>CP068266.2 Homo sapiens isolate CHM13 chromosome 12, 62817411, 62817721)
Single mismatch: (chromosome, start, end) = (>CP068266.2 Homo sapiens isolate CHM13 chromosome 12, 79100751, 79101061)
Single mismatch: (chromosome, start, end) = (>CP068265.2 Homo sapiens isolate CHM13 chromosome 13, 66839102, 66839412)
Additional character: (chromosome, start, end) = (>CP068265.2 Homo sapiens isolate CHM13 chromosome 13, 66839102, 66839413)
Single mismatch: (chromosome, start, end) = (>CP068264.2 Homo sapiens isolate CHM13 chromosome 14, 17947515, 17947825)
Single mismatch: (chromosome, start, end) = (>CP068261.2 Homo sapiens isolate CHM13 chromosome 17, 2142369, 2142679)
Single mismatch: (chromosome, start, end) = (>CP068261.2 Homo sapiens isolate CHM13 chromosome 17, 19541649, 19541959)
Single mismatch: (chromosome, start, end) = (>CP068260.2 Homo sapiens isolate CHM13 chromosome 18, 66026304, 66026614)
Single mismatch: (chromosome, start, end) = (>CP068257.2 Homo sapiens isolate CHM13 chromosome 21, 24874951, 24875261)
** Number of exact matches: 2
** Number of single mismatches: 18
** Number of additional character matches: 1
** Number of missing character matches: 0
*** Total number of matches:  21
       40.20 real        38.58 user         1.25 sys
            55377920  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               16412  page reclaims
                2550  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               23628  signals received
                2402  voluntary context switches
              139833  involuntary context switches
           761351157  instructions retired
           284404222  cycles elapsed
            13764160  peak memory footprint
```

</details>

<details>
<summary>Number of matches up to one mismatch in GRCh38 file: 17</summary>
<br>

```bash
Run: go run prob2_naive.go <seq_file>
<seq_file> =  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Pattern: ggccgggcgcggtggctcacgcctgtaatcccagcactttgggaggccgaggcgggcggatcacgaggtcaggagatcgagaccatcctggctaacacggtgaaaccccgtctctactaaaaatacaaaaaattagccgggcgtggtggcgggcgcctgtagtcccagctactcgggaggctgaggcaggagaatggcgtgaacccgggaggcggagcttgcagtgagccgagatcgcgccactgcactccagcctgggcgacagagcgagactccgtctcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Length of pattern: 311
Starting sequence file processing:  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
-----------------------------------
Single mismatch: (chromosome, start, end) = (>CM000663.2 Homo sapiens chromosome 1, GRCh38 reference primary assembly, 85940122, 85940432)
Single mismatch: (chromosome, start, end) = (>CM000663.2 Homo sapiens chromosome 1, GRCh38 reference primary assembly, 144846234, 144846544)
Single mismatch: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 69313857, 69314167)
Additional character: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629831, 88630142)
Exact match: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629832, 88630142)
Single mismatch: (chromosome, start, end) = (>CM000666.2 Homo sapiens chromosome 4, GRCh38 reference primary assembly, 132898458, 132898768)
Single mismatch: (chromosome, start, end) = (>CM000667.2 Homo sapiens chromosome 5, GRCh38 reference primary assembly, 125028156, 125028466)
Exact match: (chromosome, start, end) = (>CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly, 39808489, 39808799)
Single mismatch: (chromosome, start, end) = (>CM000670.2 Homo sapiens chromosome 8, GRCh38 reference primary assembly, 100790760, 100791070)
Single mismatch: (chromosome, start, end) = (>CM000674.2 Homo sapiens chromosome 12, GRCh38 reference primary assembly, 62838639, 62838949)
Single mismatch: (chromosome, start, end) = (>CM000675.2 Homo sapiens chromosome 13, GRCh38 reference primary assembly, 63082060, 63082370)
Single mismatch: (chromosome, start, end) = (>CM000676.2 Homo sapiens chromosome 14, GRCh38 reference primary assembly, 34836097, 34836407)
Single mismatch: (chromosome, start, end) = (>CM000677.2 Homo sapiens chromosome 15, GRCh38 reference primary assembly, 53225105, 53225415)
Exact match: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 18689763, 18690073)
Single mismatch: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 19593503, 19593813)
Single mismatch: (chromosome, start, end) = (>CM000683.2 Homo sapiens chromosome 21, GRCh38 reference primary assembly, 26516833, 26517143)
Single mismatch: (chromosome, start, end) = (>KI270778.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1, 54847, 55157)
** Number of exact matches: 3
** Number of single mismatches: 13
** Number of additional character matches: 1
** Number of missing character matches: 0
*** Total number of matches:  17
       42.11 real        40.83 user         1.24 sys
            56360960  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               16402  page reclaims
                2599  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               25270  signals received
                2730  voluntary context switches
              143309  involuntary context switches
           762603845  instructions retired
           295351590  cycles elapsed
            13370880  peak memory footprint
```

</details>

## Performance analysis

Time is reported as the real time reported by the `usr/bin/time -l` command in seconds. Memory is reported as the maximum resident set size in MB, where `usr/bin/time -l` command reports it in bytes.

<details>
<summary>script run output</summary>
<br>

```bash
>> sh ./run_scripts.sh

>>> Running all scripts.
Running commands for prob1_naive
Run: go run prob1_naive.go <seq_file>
<seq_file> =  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Matched: (chromosome, start, end) = (>CP068276.2 Homo sapiens isolate CHM13 chromosome 2, 186833277, 186833587)
Matched: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 39966041, 39966351)
*** Total number of matches:  2
       17.40 real        16.98 user         0.93 sys
            53706752  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               17533  page reclaims
                 113  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               23148  signals received
                1708  voluntary context switches
              110999  involuntary context switches
           725820987  instructions retired
           265748696  cycles elapsed
            13616640  peak memory footprint
Run: go run prob1_naive.go <seq_file>
<seq_file> =  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Matched: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629832, 88630142)
Matched: (chromosome, start, end) = (>CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly, 39808489, 39808799)
Matched: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 18689763, 18690073)
*** Total number of matches:  3
       19.06 real        18.40 user         1.07 sys
            56131584  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               18806  page reclaims
                 107  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               24319  signals received
                1117  voluntary context switches
              120336  involuntary context switches
           717251083  instructions retired
           288758227  cycles elapsed
            13534720  peak memory footprint
Running commands for prob1_kmp
Run: go run prob1_kmp.go <seq_file>
<seq_file> =  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Match: (chromosome, start, end) = (>CP068276.2 Homo sapiens isolate CHM13 chromosome 2, 186833277, 186833587)
Match: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 39966041, 39966351)
*** Total number of matches:  2
       18.66 real        18.11 user         1.03 sys
            56164352  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               17852  page reclaims
                 111  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               23310  signals received
                1271  voluntary context switches
              112335  involuntary context switches
           716548566  instructions retired
           276243778  cycles elapsed
            13780416  peak memory footprint
Run: go run prob1_kmp.go <seq_file>
<seq_file> =  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Match: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629832, 88630142)
Match: (chromosome, start, end) = (>CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly, 39808489, 39808799)
Match: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 18689763, 18690073)
*** Total number of matches:  3
       19.89 real        19.29 user         1.07 sys
            56082432  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               17793  page reclaims
                 109  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               24601  signals received
                1223  voluntary context switches
              118741  involuntary context switches
           707703017  instructions retired
           268069290  cycles elapsed
            13911552  peak memory footprint
Running commands for prob2
Run: go run prob2_naive.go <seq_file>
<seq_file> =  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Pattern: ggccgggcgcggtggctcacgcctgtaatcccagcactttgggaggccgaggcgggcggatcacgaggtcaggagatcgagaccatcctggctaacacggtgaaaccccgtctctactaaaaatacaaaaaattagccgggcgtggtggcgggcgcctgtagtcccagctactcgggaggctgaggcaggagaatggcgtgaacccgggaggcggagcttgcagtgagccgagatcgcgccactgcactccagcctgggcgacagagcgagactccgtctcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Length of pattern: 311
Starting sequence file processing:  ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
-----------------------------------
Single mismatch: (chromosome, start, end) = (>CP068277.2 Homo sapiens isolate CHM13 chromosome 1, 143946938, 143947248)
Single mismatch: (chromosome, start, end) = (>CP068277.2 Homo sapiens isolate CHM13 chromosome 1, 190858894, 190859204)
Exact match: (chromosome, start, end) = (>CP068276.2 Homo sapiens isolate CHM13 chromosome 2, 186833277, 186833587)
Single mismatch: (chromosome, start, end) = (>CP068275.2 Homo sapiens isolate CHM13 chromosome 3, 69350794, 69351104)
Single mismatch: (chromosome, start, end) = (>CP068274.2 Homo sapiens isolate CHM13 chromosome 4, 136221657, 136221967)
Single mismatch: (chromosome, start, end) = (>CP068273.2 Homo sapiens isolate CHM13 chromosome 5, 126911808, 126912118)
Single mismatch: (chromosome, start, end) = (>CP068272.2 Homo sapiens isolate CHM13 chromosome 6, 51234801, 51235111)
Single mismatch: (chromosome, start, end) = (>CP068272.2 Homo sapiens isolate CHM13 chromosome 6, 68212687, 68212997)
Exact match: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 39966041, 39966351)
Single mismatch: (chromosome, start, end) = (>CP068271.2 Homo sapiens isolate CHM13 chromosome 7, 49189460, 49189770)
Single mismatch: (chromosome, start, end) = (>CP068270.2 Homo sapiens isolate CHM13 chromosome 8, 101916771, 101917081)
Single mismatch: (chromosome, start, end) = (>CP068267.2 Homo sapiens isolate CHM13 chromosome 11, 86916177, 86916487)
Single mismatch: (chromosome, start, end) = (>CP068266.2 Homo sapiens isolate CHM13 chromosome 12, 62817411, 62817721)
Single mismatch: (chromosome, start, end) = (>CP068266.2 Homo sapiens isolate CHM13 chromosome 12, 79100751, 79101061)
Single mismatch: (chromosome, start, end) = (>CP068265.2 Homo sapiens isolate CHM13 chromosome 13, 66839102, 66839412)
Additional character: (chromosome, start, end) = (>CP068265.2 Homo sapiens isolate CHM13 chromosome 13, 66839102, 66839413)
Single mismatch: (chromosome, start, end) = (>CP068264.2 Homo sapiens isolate CHM13 chromosome 14, 17947515, 17947825)
Single mismatch: (chromosome, start, end) = (>CP068261.2 Homo sapiens isolate CHM13 chromosome 17, 2142369, 2142679)
Single mismatch: (chromosome, start, end) = (>CP068261.2 Homo sapiens isolate CHM13 chromosome 17, 19541649, 19541959)
Single mismatch: (chromosome, start, end) = (>CP068260.2 Homo sapiens isolate CHM13 chromosome 18, 66026304, 66026614)
Single mismatch: (chromosome, start, end) = (>CP068257.2 Homo sapiens isolate CHM13 chromosome 21, 24874951, 24875261)
** Number of exact matches: 2
** Number of single mismatches: 18
** Number of additional character matches: 1
** Number of missing character matches: 0
*** Total number of matches:  21
       39.22 real        38.40 user         1.29 sys
            55394304  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               17634  page reclaims
                 105  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               23846  signals received
                1187  voluntary context switches
              116855  involuntary context switches
           713959285  instructions retired
           278799439  cycles elapsed
            13452864  peak memory footprint
Run: go run prob2_naive.go <seq_file>
<seq_file> =  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
Pattern-- Header: >DF000000002.4 AluY
Pattern: ggccgggcgcggtggctcacgcctgtaatcccagcactttgggaggccgaggcgggcggatcacgaggtcaggagatcgagaccatcctggctaacacggtgaaaccccgtctctactaaaaatacaaaaaattagccgggcgtggtggcgggcgcctgtagtcccagctactcgggaggctgaggcaggagaatggcgtgaacccgggaggcggagcttgcagtgagccgagatcgcgccactgcactccagcctgggcgacagagcgagactccgtctcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
Length of pattern: 311
Starting sequence file processing:  ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
-----------------------------------
Single mismatch: (chromosome, start, end) = (>CM000663.2 Homo sapiens chromosome 1, GRCh38 reference primary assembly, 85940122, 85940432)
Single mismatch: (chromosome, start, end) = (>CM000663.2 Homo sapiens chromosome 1, GRCh38 reference primary assembly, 144846234, 144846544)
Single mismatch: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 69313857, 69314167)
Additional character: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629831, 88630142)
Exact match: (chromosome, start, end) = (>CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly, 88629832, 88630142)
Single mismatch: (chromosome, start, end) = (>CM000666.2 Homo sapiens chromosome 4, GRCh38 reference primary assembly, 132898458, 132898768)
Single mismatch: (chromosome, start, end) = (>CM000667.2 Homo sapiens chromosome 5, GRCh38 reference primary assembly, 125028156, 125028466)
Exact match: (chromosome, start, end) = (>CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly, 39808489, 39808799)
Single mismatch: (chromosome, start, end) = (>CM000670.2 Homo sapiens chromosome 8, GRCh38 reference primary assembly, 100790760, 100791070)
Single mismatch: (chromosome, start, end) = (>CM000674.2 Homo sapiens chromosome 12, GRCh38 reference primary assembly, 62838639, 62838949)
Single mismatch: (chromosome, start, end) = (>CM000675.2 Homo sapiens chromosome 13, GRCh38 reference primary assembly, 63082060, 63082370)
Single mismatch: (chromosome, start, end) = (>CM000676.2 Homo sapiens chromosome 14, GRCh38 reference primary assembly, 34836097, 34836407)
Single mismatch: (chromosome, start, end) = (>CM000677.2 Homo sapiens chromosome 15, GRCh38 reference primary assembly, 53225105, 53225415)
Exact match: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 18689763, 18690073)
Single mismatch: (chromosome, start, end) = (>CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly, 19593503, 19593813)
Single mismatch: (chromosome, start, end) = (>CM000683.2 Homo sapiens chromosome 21, GRCh38 reference primary assembly, 26516833, 26517143)
Single mismatch: (chromosome, start, end) = (>KI270778.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1, 54847, 55157)
** Number of exact matches: 3
** Number of single mismatches: 13
** Number of additional character matches: 1
** Number of missing character matches: 0
*** Total number of matches:  17
       41.37 real        40.47 user         1.37 sys
            56115200  maximum resident set size
                   0  average shared memory size
                   0  average unshared data size
                   0  average unshared stack size
               18831  page reclaims
                 102  page faults
                   0  swaps
                   0  block input operations
                   0  block output operations
                   0  messages sent
                   0  messages received
               24911  signals received
                1220  voluntary context switches
              125431  involuntary context switches
           726368999  instructions retired
           291797706  cycles elapsed
            14091840  peak memory footprint
```

</details>

| File Name  | T2T - Time (s)| T2T - Memory (MB) | GRCh38 - Time (S) | GRCh38 - Memory (MB) |
|-------------|----------------|------------------|----------------|------------------|
| prob1_naive.go | 17.40 | 53.706752 | 19.06 | 56.131584 |
| prob1_kmp.go   | 18.66 | 56.164352 | 19.89 | 56.082432 |
| prob2.go       | 39.22 | 55.394304 | 41.37 | 56.115200 |

The T2T file is 3.16 GB, and the GRCh38 file is 3.34 GB.

I initially implemented prob1_naive.go in Python, but the implementation took 16 minutes to run on the T2T file. I then implemented the same code in Go, and it took 20 seconds to run. Hence, I made the decision to continue with Go for the rest of the implementations. Interestingly, the Python implementation had a maximum resident set size of 9 MB, which is 5 times less than Go, but I haven't investigated this further. 

## Running the code

Go version: go1.23.0 darwin/arm64.
I ran scripts locally on M2 Macbook air.

- download the sequence files from the [link provided](https://www.ncbi.nlm.nih.gov/datasets/genome/?taxon=9606) and decompress them in the main directory.

To run the code, you can use the following commands:

```bash
sh ./run_scripts.sh [prob1_naive | prob1_kmp | prob2] 
```

Running the above command without additional arguments will run all three implementations. Alternatively, you can run each implementation seperately and pass the path to the sequence file as an argument. For example:

```bash
go run prob1_naive.go <path_to_sequence_file>
```

Note that in the bash script, I used the `time` command of format `/usr/bin/time -l` to measure the memory usage of the program, and this can be different depending on the system.
