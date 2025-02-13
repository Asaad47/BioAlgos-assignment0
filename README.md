# Introduction

This is the first assignment of CS 249 Algorithms in Bioinformatics. The assignment is mainly composed of two implementation parts: exact pattern matching and matching up to one mismatch. In both parts, the goal is to find the pattern sequence AluY in the sequence files of human genome assemblies GRCh38 (hg38) and T2T CHM13v2.0 retrieved from NCBI datasets. The AluY sequence is provided in the assignment description and available in the [FASTA](./DF000000002.fa) file. I used naive character matching and the Knuth-Morris-Pratt (KMP) algorithm for exact pattern matching and implemented single mismatch, single insertion, and single deletion for matching up to one mismatch. The results are reported in the [Results](#results-on-aluy) section.

## Brief implementation description

I have three main files that address the implementation problems: prob1_naive.go, prob1_kmp.go, and prob2.go. In all three files, I read the sequence file line by line and start searching for the pattern in the sequence whenever I have enough characters in the `queue` to match the pattern. Essentially, the `queue` is a buffer that holds read characters from the sequence file and searching happens whenever the `queue` is at least the length of the pattern. This is to avoid loading the entire sequence file into memory.

For finding exact matches, I used two implementations for searching: naive character matching and KMP algorithm found in prob1_naive.go and prob1_kmp.go respectively. The naive matching directly compares the characters in the `queue` with the pattern in each iteration while the KMP algorithm pre-computes the border array and uses it to skip unnecessary comparisons. However, both implementations achieved roughly the same performance.

For finding approximate matches, I seperated the search mechanism to four functions: exact match, single mismatch, single insertion, and single deletion. The exact match is the same as the naive implementation of prob1. The single mismatch resovles to true if the Hamming distance is exactly 1. The single insertion and single deletion look for the first mismatch index and then skip the mismatched character in the `queue` and the pattern respectively. Some additional checks are used to avoid counting duplicates. For example, if there a was an exact match, the single deletion function will not be called as it is a subset of the exact match. Also, the single addition function asserts that the first and last charcaters are the same as the pattern and that the mismatch is in the middle to avoid duplicates.

Note: I have done pattern match searching per header contig in the sequnce file, so when reading a new header, the `queue` and `start` index are reset and the search starts from the beginning of the new contig.

## Results on AluY

Dropdowns have more information on the results on specified files. The _start_ and _end_ indices are inclusive.

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

| File Name  | T2T - Time (s)| T2T - Memory (MB) | GRCh38 - Time (s) | GRCh38 - Memory (MB) |
|-------------|----------------|------------------|----------------|------------------|
| prob1_naive.go | 17.40 | 53.706752 | 19.06 | 56.131584 |
| prob1_kmp.go   | 18.66 | 56.164352 | 19.89 | 56.082432 |
| prob2.go       | 39.22 | 55.394304 | 41.37 | 56.115200 |

The T2T file is 3.16 GB, and the GRCh38 file is 3.34 GB, so the implementations take a little bit more time processing the GRCh38 file. The naive implementation is a little bit faster than KMP in both files. The memory usage is similar for all implementations. Counting mismatches takes almost double the time compared to counting exact matches, which seems to be good performing since the implementation could potentially make up to four times the number of comparisons. There still might be room for improvement in the implementation, but the current performance is acceptable with runtime less than a minute with a maximum memory footprint of 56 MB.

<details>
<summary> T2T file has 25 header lines</summary>
<br>

```bash
Header: >CP068277.2 Homo sapiens isolate CHM13 chromosome 1
Header: >CP068276.2 Homo sapiens isolate CHM13 chromosome 2
Header: >CP068275.2 Homo sapiens isolate CHM13 chromosome 3
Header: >CP068274.2 Homo sapiens isolate CHM13 chromosome 4
Header: >CP068273.2 Homo sapiens isolate CHM13 chromosome 5
Header: >CP068272.2 Homo sapiens isolate CHM13 chromosome 6
Header: >CP068271.2 Homo sapiens isolate CHM13 chromosome 7
Header: >CP068270.2 Homo sapiens isolate CHM13 chromosome 8
Header: >CP068269.2 Homo sapiens isolate CHM13 chromosome 9
Header: >CP068268.2 Homo sapiens isolate CHM13 chromosome 10
Header: >CP068267.2 Homo sapiens isolate CHM13 chromosome 11
Header: >CP068266.2 Homo sapiens isolate CHM13 chromosome 12
Header: >CP068265.2 Homo sapiens isolate CHM13 chromosome 13
Header: >CP068264.2 Homo sapiens isolate CHM13 chromosome 14
Header: >CP068263.2 Homo sapiens isolate CHM13 chromosome 15
Header: >CP068262.2 Homo sapiens isolate CHM13 chromosome 16
Header: >CP068261.2 Homo sapiens isolate CHM13 chromosome 17
Header: >CP068260.2 Homo sapiens isolate CHM13 chromosome 18
Header: >CP068259.2 Homo sapiens isolate CHM13 chromosome 19
Header: >CP068258.2 Homo sapiens isolate CHM13 chromosome 20
Header: >CP068257.2 Homo sapiens isolate CHM13 chromosome 21
Header: >CP068256.2 Homo sapiens isolate CHM13 chromosome 22
Header: >CP068255.2 Homo sapiens isolate CHM13 chromosome X
Header: >CP086569.2 Homo sapiens isolate NA24385 chromosome Y
Header: >CP068254.1 Homo sapiens isolate CHM13 mitochondrion, complete genome
```

</details>

<details>
<summary>The GRCh38 file has 709 headers</summary>
<br>

```bash
Header: >CM000663.2 Homo sapiens chromosome 1, GRCh38 reference primary assembly
Header: >KI270706.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270707.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270708.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270709.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270710.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270711.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270712.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270713.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270714.1 Homo sapiens chromosome 1 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000664.2 Homo sapiens chromosome 2, GRCh38 reference primary assembly
Header: >KI270715.1 Homo sapiens chromosome 2 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270716.1 Homo sapiens chromosome 2 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000665.2 Homo sapiens chromosome 3, GRCh38 reference primary assembly
Header: >GL000221.1 Homo sapiens chromosome 3 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000666.2 Homo sapiens chromosome 4, GRCh38 reference primary assembly
Header: >GL000008.2 Homo sapiens chromosome 4 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000667.2 Homo sapiens chromosome 5, GRCh38 reference primary assembly
Header: >GL000208.1 Homo sapiens chromosome 5 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000668.2 Homo sapiens chromosome 6, GRCh38 reference primary assembly
Header: >CM000669.2 Homo sapiens chromosome 7, GRCh38 reference primary assembly
Header: >CM000670.2 Homo sapiens chromosome 8, GRCh38 reference primary assembly
Header: >CM000671.2 Homo sapiens chromosome 9, GRCh38 reference primary assembly
Header: >KI270717.1 Homo sapiens chromosome 9 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270718.1 Homo sapiens chromosome 9 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270719.1 Homo sapiens chromosome 9 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270720.1 Homo sapiens chromosome 9 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000672.2 Homo sapiens chromosome 10, GRCh38 reference primary assembly
Header: >CM000673.2 Homo sapiens chromosome 11, GRCh38 reference primary assembly
Header: >KI270721.1 Homo sapiens chromosome 11 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000674.2 Homo sapiens chromosome 12, GRCh38 reference primary assembly
Header: >CM000675.2 Homo sapiens chromosome 13, GRCh38 reference primary assembly
Header: >CM000676.2 Homo sapiens chromosome 14, GRCh38 reference primary assembly
Header: >GL000009.2 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >GL000225.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270722.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >GL000194.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270723.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270724.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270725.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270726.1 Homo sapiens chromosome 14 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000677.2 Homo sapiens chromosome 15, GRCh38 reference primary assembly
Header: >KI270727.1 Homo sapiens chromosome 15 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000678.2 Homo sapiens chromosome 16, GRCh38 reference primary assembly
Header: >KI270728.1 Homo sapiens chromosome 16 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000679.2 Homo sapiens chromosome 17, GRCh38 reference primary assembly
Header: >GL000205.2 Homo sapiens chromosome 17 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270729.1 Homo sapiens chromosome 17 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270730.1 Homo sapiens chromosome 17 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000680.2 Homo sapiens chromosome 18, GRCh38 reference primary assembly
Header: >CM000681.2 Homo sapiens chromosome 19, GRCh38 reference primary assembly
Header: >CM000682.2 Homo sapiens chromosome 20, GRCh38 reference primary assembly
Header: >CM000683.2 Homo sapiens chromosome 21, GRCh38 reference primary assembly
Header: >CM000684.2 Homo sapiens chromosome 22, GRCh38 reference primary assembly
Header: >KI270731.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270732.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270733.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270734.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270735.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270736.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270737.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270738.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270739.1 Homo sapiens chromosome 22 unlocalized genomic contig, GRCh38 reference primary assembly
Header: >CM000685.2 Homo sapiens chromosome X, GRCh38 reference primary assembly
Header: >CM000686.2 Homo sapiens chromosome Y, GRCh38 reference primary assembly
Header: >KI270740.1 Homo sapiens chromosome Y unlocalized genomic contig, GRCh38 reference primary assembly
Header: >KI270302.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270304.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270303.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270305.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270322.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270320.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270310.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270316.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270315.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270312.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270311.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270317.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270412.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270411.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270414.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270419.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270418.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270420.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270424.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270417.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270422.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270423.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270425.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270429.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270442.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270466.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270465.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270467.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270435.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270438.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270468.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270510.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270509.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270518.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270508.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270516.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270512.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270519.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270522.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270511.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270515.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270507.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270517.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270529.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270528.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270530.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270539.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270538.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270544.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270548.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270583.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270587.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270580.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270581.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270579.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270589.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270590.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270584.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270582.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270588.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270593.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270591.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270330.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270329.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270334.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270333.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270335.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270338.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270340.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270336.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270337.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270363.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270364.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270362.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270366.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270378.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270379.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270389.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270390.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270387.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270395.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270396.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270388.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270394.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270386.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270391.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270383.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270393.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270384.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270392.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270381.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270385.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270382.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270376.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270374.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270372.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270373.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270375.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270371.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270448.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270521.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000195.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000219.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000220.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000224.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270741.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000226.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000213.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270743.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270744.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270745.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270746.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270747.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270748.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270749.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270750.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270751.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270752.1 UNVERIFIED_ORG: Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270753.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270754.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270755.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270756.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270757.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000214.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KI270742.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000216.2 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >GL000218.1 Homo sapiens unplaced genomic contig, GRCh38 reference primary assembly
Header: >KQ031383.1 Homo sapiens chromosome 1 genomic contig HG1342_HG2282_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ983255.1 Homo sapiens chromosome 1 genomic contig HSCHR1_5_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273333.1 Homo sapiens chromosome 1 genomic contig HG1343_HG173_HG459_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538361.1 Homo sapiens chromosome 1 genomic contig HG2095_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ458383.1 Homo sapiens chromosome 1 genomic contig HSCHR1_4_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN196473.1 Homo sapiens chromosome 1 genomic contig HG2058_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208904.1 Homo sapiens chromosome 1 genomic contig HSCHR1_8_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559100.1 Homo sapiens chromosome 1 genomic contig HG460_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196472.1 Homo sapiens chromosome 1 genomic contig HG986_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208905.1 Homo sapiens chromosome 1 genomic contig HSCHR1_9_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ458382.1 Homo sapiens chromosome 1 genomic contig HSCHR1_3_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV880763.1 Homo sapiens chromosome 1 genomic contig HSCHR1_6_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN196474.1 Homo sapiens chromosome 1 genomic contig HG2104_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273330.1 Homo sapiens chromosome 1 genomic contig HSCHR1_12_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273335.1 Homo sapiens chromosome 1 genomic contig HG2515_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273336.1 Homo sapiens chromosome 1 genomic contig HG2577_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273331.1 Homo sapiens chromosome 1 genomic contig HSCHR1_5_CTG31, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN538360.1 Homo sapiens chromosome 1 genomic contig HG1832_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208906.1 Homo sapiens chromosome 1 genomic contig HG2002_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ458384.1 Homo sapiens chromosome 1 genomic contig HSCHR1_5_CTG32_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273334.1 Homo sapiens chromosome 1 genomic contig HG2571_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273332.1 Homo sapiens chromosome 1 genomic contig HSCHR1_6_CTG31, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143342.1 Homo sapiens chromosome 2 genomic contig HG1384_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273344.1 Homo sapiens chromosome 2 genomic contig HG2231_HG2496_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273345.1 Homo sapiens chromosome 2 genomic contig HG2140_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273343.1 Homo sapiens chromosome 2 genomic contig HG2052_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273340.1 Homo sapiens chromosome 2 genomic contig HSCHR2_6_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ031384.1 Homo sapiens chromosome 2 genomic contig HG2290_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273337.1 Homo sapiens chromosome 2 genomic contig HSCHR2_10_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273342.1 Homo sapiens chromosome 2 genomic contig HG2275_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208907.1 Homo sapiens chromosome 2 genomic contig HSCHR2_7_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273339.1 Homo sapiens chromosome 2 genomic contig HSCHR2_12_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273338.1 Homo sapiens chromosome 2 genomic contig HSCHR2_11_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273341.1 Homo sapiens chromosome 2 genomic contig HG2494_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ983256.1 Homo sapiens chromosome 2 genomic contig HSCHR2_6_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208908.1 Homo sapiens chromosome 2 genomic contig HSCHR2_8_CTG7_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN538363.1 Homo sapiens chromosome 2 genomic contig HG2232_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143341.1 Homo sapiens chromosome 2 genomic contig HG721_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538362.1 Homo sapiens chromosome 2 genomic contig HG2233_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766192.1 Homo sapiens chromosome 3 genomic contig HG2236_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273347.1 Homo sapiens chromosome 3 genomic contig HG2077_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273348.1 Homo sapiens chromosome 3 genomic contig HG2069_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196475.1 Homo sapiens chromosome 3 genomic contig HG2066_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ031385.1 Homo sapiens chromosome 3 genomic contig HG2235_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538364.1 Homo sapiens chromosome 3 genomic contig HG126_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143343.1 Homo sapiens chromosome 3 genomic contig HSCHR3_5_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208909.1 Homo sapiens chromosome 3 genomic contig HSCHR3_4_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ031386.1 Homo sapiens chromosome 3 genomic contig HG2237_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196476.1 Homo sapiens chromosome 3 genomic contig HG2022_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559104.1 Homo sapiens chromosome 3 genomic contig HG2133_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559105.1 Homo sapiens chromosome 3 genomic contig HSCHR3_6_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559103.1 Homo sapiens chromosome 3 genomic contig HSCHR3_9_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559102.1 Homo sapiens chromosome 3 genomic contig HSCHR3_8_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559101.1 Homo sapiens chromosome 3 genomic contig HSCHR3_7_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273346.1 Homo sapiens chromosome 3 genomic contig HG2264_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143344.1 Homo sapiens chromosome 4 genomic contig HG699_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143347.1 Homo sapiens chromosome 4 genomic contig HG1298_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090013.1 Homo sapiens chromosome 4 genomic contig HSCHR4_2_CTG4, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273351.1 Homo sapiens chromosome 4 genomic contig HG287_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143346.1 Homo sapiens chromosome 4 genomic contig HG1299_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143348.1 Homo sapiens chromosome 4 genomic contig HG1296_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143345.1 Homo sapiens chromosome 4 genomic contig HG2525_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143349.1 Homo sapiens chromosome 4 genomic contig HG705_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090014.1 Homo sapiens chromosome 4 genomic contig HSCHR4_8_CTG12, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ090015.1 Homo sapiens chromosome 4 genomic contig HSCHR4_9_CTG12, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273349.1 Homo sapiens chromosome 4 genomic contig HSCHR4_2_CTG8_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV766193.1 Homo sapiens chromosome 4 genomic contig HSCHR4_12_CTG12, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273350.1 Homo sapiens chromosome 4 genomic contig HG2155_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ983258.1 Homo sapiens chromosome 4 genomic contig HSCHR4_11_CTG12, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ983257.1 Homo sapiens chromosome 4 genomic contig HG2023_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273354.1 Homo sapiens chromosome 5 genomic contig HG2405_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208910.1 Homo sapiens chromosome 5 genomic contig HSCHR5_9_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273353.1 Homo sapiens chromosome 5 genomic contig HG2476_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196477.1 Homo sapiens chromosome 5 genomic contig HSCHR5_7_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575243.1 Homo sapiens chromosome 5 genomic contig HSCHR5_8_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273356.1 Homo sapiens chromosome 5 genomic contig HSCHR5_10_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143350.1 Homo sapiens chromosome 5 genomic contig HG1395_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273355.1 Homo sapiens chromosome 5 genomic contig HG2308_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV575244.1 Homo sapiens chromosome 5 genomic contig HG30_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273352.1 Homo sapiens chromosome 5 genomic contig HG1046_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208911.1 Homo sapiens chromosome 6 genomic contig HG2057_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090017.1 Homo sapiens chromosome 6 genomic contig HSCHR6_1_CTG10, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273357.1 Homo sapiens chromosome 6 genomic contig HSCHR6_1_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ031387.1 Homo sapiens chromosome 6 genomic contig HG1651_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196478.1 Homo sapiens chromosome 6 genomic contig HG2128_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090016.1 Homo sapiens chromosome 6 genomic contig HG2072_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766194.1 Homo sapiens chromosome 6 genomic contig HG2121_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143351.1 Homo sapiens chromosome 6 genomic contig HG563_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143352.1 Homo sapiens chromosome 7 genomic contig HG1309_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559106.1 Homo sapiens chromosome 7 genomic contig HSCHR7_3_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273358.1 Homo sapiens chromosome 7 genomic contig HSCHR7_4_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV880764.1 Homo sapiens chromosome 7 genomic contig HG2088_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV880765.1 Homo sapiens chromosome 7 genomic contig HG2266_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208912.1 Homo sapiens chromosome 7 genomic contig HG708_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208913.1 Homo sapiens chromosome 7 genomic contig HSCHR7_3_CTG4_4, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ031388.1 Homo sapiens chromosome 7 genomic contig HG2239_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273362.1 Homo sapiens chromosome 8 genomic contig HG2267_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208915.1 Homo sapiens chromosome 8 genomic contig HG76_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV880767.1 Homo sapiens chromosome 8 genomic contig HG2068_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273359.1 Homo sapiens chromosome 8 genomic contig HG2176_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV880766.1 Homo sapiens chromosome 8 genomic contig HG2067_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273361.1 Homo sapiens chromosome 8 genomic contig HG2408_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559107.1 Homo sapiens chromosome 8 genomic contig HSCHR8_7_CTG7, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273363.1 Homo sapiens chromosome 8 genomic contig HG2031_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208914.1 Homo sapiens chromosome 8 genomic contig HG2419_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273360.1 Homo sapiens chromosome 8 genomic contig HG1047_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090018.1 Homo sapiens chromosome 9 genomic contig HSCHR9_1_CTG6, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ090019.1 Homo sapiens chromosome 9 genomic contig HSCHR9_1_CTG7, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273366.1 Homo sapiens chromosome 9 genomic contig HG1206_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273364.1 Homo sapiens chromosome 9 genomic contig HG2158_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273365.1 Homo sapiens chromosome 9 genomic contig HG1012_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196479.1 Homo sapiens chromosome 9 genomic contig HG2030_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143353.1 Homo sapiens chromosome 9 genomic contig HG613_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143354.1 Homo sapiens chromosome 10 genomic contig HG545_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538367.1 Homo sapiens chromosome 10 genomic contig HG2244_HG2245_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143355.1 Homo sapiens chromosome 10 genomic contig HG1277_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090020.1 Homo sapiens chromosome 10 genomic contig HSCHR10_1_CTG6, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN196480.1 Homo sapiens chromosome 10 genomic contig HG2191_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090021.1 Homo sapiens chromosome 10 genomic contig HG2334_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273367.1 Homo sapiens chromosome 10 genomic contig HG2576_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538366.1 Homo sapiens chromosome 10 genomic contig HG2242_HG2243_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538365.1 Homo sapiens chromosome 10 genomic contig HG2241_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ759759.2 Homo sapiens chromosome 11 genomic contig HG107_HG2565_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273369.1 Homo sapiens chromosome 11 genomic contig HG152_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143358.1 Homo sapiens chromosome 11 genomic contig HG28_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273371.1 Homo sapiens chromosome 11 genomic contig HG2578_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538368.1 Homo sapiens chromosome 11 genomic contig HSCHR11_1_CTG1_2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143360.1 Homo sapiens chromosome 11 genomic contig HG2111_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559109.1 Homo sapiens chromosome 11 genomic contig HG2114_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559108.1 Homo sapiens chromosome 11 genomic contig HG2060_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766195.1 Homo sapiens chromosome 11 genomic contig HG1708_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273370.1 Homo sapiens chromosome 11 genomic contig HG2568_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559111.1 Homo sapiens chromosome 11 genomic contig HSCHR11_1_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143359.1 Homo sapiens chromosome 11 genomic contig HG2115_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143357.1 Homo sapiens chromosome 11 genomic contig HG1445_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559110.1 Homo sapiens chromosome 11 genomic contig HSCHR11_2_CTG8, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ090022.1 Homo sapiens chromosome 11 genomic contig HG2116_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143356.1 Homo sapiens chromosome 11 genomic contig HG1521_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273368.1 Homo sapiens chromosome 11 genomic contig HSCHR11_2_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN196481.1 Homo sapiens chromosome 11 genomic contig HG2217_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090023.1 Homo sapiens chromosome 12 genomic contig HSCHR12_2_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208916.1 Homo sapiens chromosome 12 genomic contig HG1815_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143362.1 Homo sapiens chromosome 12 genomic contig HG1398_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538369.1 Homo sapiens chromosome 12 genomic contig HG1362_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196482.1 Homo sapiens chromosome 12 genomic contig HG23_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273372.1 Homo sapiens chromosome 12 genomic contig HG2554_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208918.1 Homo sapiens chromosome 12 genomic contig HSCHR12_8_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ759760.1 Homo sapiens chromosome 12 genomic contig HG2063_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208917.1 Homo sapiens chromosome 12 genomic contig HG2047_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538370.1 Homo sapiens chromosome 12 genomic contig HG2247_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559112.1 Homo sapiens chromosome 12 genomic contig HSCHR12_9_CTG2_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143361.1 Homo sapiens chromosome 12 genomic contig HG2246_HG2248_HG2276_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143366.1 Homo sapiens chromosome 13 genomic contig HG2509_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538372.1 Homo sapiens chromosome 13 genomic contig HG2291_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090024.1 Homo sapiens chromosome 13 genomic contig HSCHR13_1_CTG7, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143363.1 Homo sapiens chromosome 13 genomic contig HG1817_1_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196483.1 Homo sapiens chromosome 13 genomic contig HG2216_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538373.1 Homo sapiens chromosome 13 genomic contig HG2249_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090025.1 Homo sapiens chromosome 13 genomic contig HSCHR13_1_CTG8, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143364.1 Homo sapiens chromosome 13 genomic contig HG1523_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143365.1 Homo sapiens chromosome 13 genomic contig HG1524_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN538371.1 Homo sapiens chromosome 13 genomic contig HG2288_HG2289_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143367.1 Homo sapiens chromosome 14 genomic contig HG2510_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273373.1 Homo sapiens chromosome 14 genomic contig HG2526_HG2573_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208920.1 Homo sapiens chromosome 14 genomic contig HG1_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143368.1 Homo sapiens chromosome 14 genomic contig HSCHR14_9_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208919.1 Homo sapiens chromosome 14 genomic contig HSCHR14_8_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN538374.1 Homo sapiens chromosome 15 genomic contig HG2139_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143372.1 Homo sapiens chromosome 15 genomic contig HG2511_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143371.1 Homo sapiens chromosome 15 genomic contig HG2365_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ031389.1 Homo sapiens chromosome 15 genomic contig HSCHR15_6_CTG8, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273375.1 Homo sapiens chromosome 15 genomic contig HSCHR15_9_CTG8, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143370.1 Homo sapiens chromosome 15 genomic contig HG2198_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273374.1 Homo sapiens chromosome 15 genomic contig HG2280_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143369.1 Homo sapiens chromosome 15 genomic contig HG2499_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273376.1 Homo sapiens chromosome 16 genomic contig HG401_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090026.1 Homo sapiens chromosome 16 genomic contig HSCHR16_5_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559113.1 Homo sapiens chromosome 16 genomic contig HG2263_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV880768.1 Homo sapiens chromosome 16 genomic contig HG926_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143373.1 Homo sapiens chromosome 16 genomic contig HG2471_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090027.1 Homo sapiens chromosome 16 genomic contig HSCHR16_4_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208921.1 Homo sapiens chromosome 16 genomic contig HSCHR16_5_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ031390.1 Homo sapiens chromosome 16 genomic contig HSCHR16_3_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273377.1 Homo sapiens chromosome 16 genomic contig HG405_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766196.1 Homo sapiens chromosome 17 genomic contig HG2285_HG106_HG2252_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143374.1 Homo sapiens chromosome 17 genomic contig HG2087_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV575245.1 Homo sapiens chromosome 17 genomic contig HG2046_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766198.1 Homo sapiens chromosome 17 genomic contig HSCHR17_3_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273380.1 Homo sapiens chromosome 17 genomic contig HG2407_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273378.1 Homo sapiens chromosome 17 genomic contig HSCHR17_13_CTG4, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV766197.1 Homo sapiens chromosome 17 genomic contig HSCHR17_11_CTG4, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559114.1 Homo sapiens chromosome 17 genomic contig HSCHR17_12_CTG4, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273383.1 Homo sapiens chromosome 17 genomic contig HG2580_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273379.1 Homo sapiens chromosome 17 genomic contig HG2118_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273382.1 Homo sapiens chromosome 17 genomic contig HG1369_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143375.1 Homo sapiens chromosome 17 genomic contig HG1320_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273381.1 Homo sapiens chromosome 17 genomic contig HG2251_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ559116.1 Homo sapiens chromosome 18 genomic contig HSCHR18_1_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ458385.1 Homo sapiens chromosome 18 genomic contig HSCHR18_5_CTG1_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ559115.1 Homo sapiens chromosome 18 genomic contig HG2412_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ090028.1 Homo sapiens chromosome 18 genomic contig HG2213_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208922.1 Homo sapiens chromosome 18 genomic contig HG2442_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273386.1 Homo sapiens chromosome 19 genomic contig HG2469_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273387.1 Homo sapiens chromosome 19 genomic contig HSCHR19_6_CTG2, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273384.1 Homo sapiens chromosome 19 genomic contig HG2461_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143376.1 Homo sapiens chromosome 19 genomic contig HG109_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ458386.1 Homo sapiens chromosome 19 genomic contig HG26_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273385.1 Homo sapiens chromosome 19 genomic contig HG2569_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196484.1 Homo sapiens chromosome 19 genomic contig HG2021_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV575246.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_0019-4656-A_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575247.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_CA01-TA01_1_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575248.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_CA01-TA01_2_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575249.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_CA01-TB04_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575250.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_CA01-TB01_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575251.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_HG2394_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575252.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_502960008-2_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575253.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_502960008-1_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575254.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_0010-5217-AB_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575255.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_7191059-1_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575256.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_0019-4656-B_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575257.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_CA04_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575259.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_7191059-2_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575260.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_HG2396_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KV575258.1 Homo sapiens chromosome 19 genomic contig HSCHR19KIR_HG2393_CTG3_1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >MU273388.1 Homo sapiens chromosome 20 genomic contig HG2225_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273389.1 Homo sapiens chromosome 20 genomic contig HG410_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143377.1 Homo sapiens chromosome 21 genomic contig HG2513_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273390.1 Homo sapiens chromosome 21 genomic contig HG2219_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273391.1 Homo sapiens chromosome 21 genomic contig HG2265_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273392.1 Homo sapiens chromosome 21 genomic contig HG2521_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143380.1 Homo sapiens chromosome 22 genomic contig HG2512_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143378.1 Homo sapiens chromosome 22 genomic contig HG1485_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196485.1 Homo sapiens chromosome 22 genomic contig HSCHR22_4_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ458387.1 Homo sapiens chromosome 22 genomic contig HSCHR22_6_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ458388.1 Homo sapiens chromosome 22 genomic contig HSCHR22_7_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KN196486.1 Homo sapiens chromosome 22 genomic contig HSCHR22_5_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KQ759761.1 Homo sapiens chromosome 22 genomic contig HSCHR22_8_CTG1, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143379.1 Homo sapiens chromosome 22 genomic contig HG494_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KQ759762.2 Homo sapiens chromosome 22 genomic contig HG1311_HG2539_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KV766199.1 Homo sapiens chromosome X genomic contig HSCHRX_3_CTG7, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143385.1 Homo sapiens chromosome X genomic contig HG1466_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273397.1 Homo sapiens chromosome X genomic contig HSCHRX_3_CTG3, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143382.1 Homo sapiens chromosome X genomic contig HG1506_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273393.1 Homo sapiens chromosome X genomic contig HG2527_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143383.1 Homo sapiens chromosome X genomic contig HG1507_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273394.1 Homo sapiens chromosome X genomic contig HG2541_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >ML143381.1 Homo sapiens chromosome X genomic contig HG439_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273396.1 Homo sapiens chromosome X genomic contig HSCHRX_2_CTG14, GRC reference assembly NOVEL PATCH for GRCh38
Header: >ML143384.1 Homo sapiens chromosome X genomic contig HG1509_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273395.1 Homo sapiens chromosome X genomic contig HSCHRX_1_CTG14, GRC reference assembly NOVEL PATCH for GRCh38
Header: >KZ208923.1 Homo sapiens chromosome Y genomic contig HG1531_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >MU273398.1 Homo sapiens chromosome Y genomic contig HG1532_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KZ208924.1 Homo sapiens chromosome Y genomic contig HG1535_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KN196487.1 Homo sapiens chromosome Y genomic contig HG2062_PATCH, GRC reference assembly FIX PATCH for GRCh38
Header: >KI270762.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270766.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270760.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270765.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383518.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383519.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383520.2 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270764.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270763.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270759.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270761.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270770.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270773.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270774.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270769.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383521.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270772.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270775.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270771.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270768.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL582966.2 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383522.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270776.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270767.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >JH636055.2 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270783.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270780.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383526.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270777.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270778.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270781.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270779.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270782.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270784.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270790.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383528.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270787.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL000257.2 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270788.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383527.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270785.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270789.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270786.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270793.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270792.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270791.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383532.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL949742.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270794.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL339449.2 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383530.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270796.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383531.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270795.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL000250.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270800.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270799.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383533.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270801.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270802.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KB021644.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270797.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270798.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270804.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270809.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270806.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383534.2 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270803.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270808.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270807.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270805.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270818.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270812.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270811.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270821.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270813.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270822.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270814.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270810.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270819.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270820.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270817.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270816.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270815.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383539.1 Homo sapiens chromosome 9 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383540.1 Homo sapiens chromosome 9 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383541.1 Homo sapiens chromosome 9 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383542.1 Homo sapiens chromosome 9 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270823.1 Homo sapiens chromosome 9 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383545.1 Homo sapiens chromosome 10 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270824.1 Homo sapiens chromosome 10 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383546.1 Homo sapiens chromosome 10 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270825.1 UNVERIFIED_ORG: Homo sapiens chromosome 10 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270832.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270830.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270831.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270829.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383547.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >JH159136.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >JH159137.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270827.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270826.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL877875.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL877876.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270837.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383549.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270835.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383550.2 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383552.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383553.2 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270834.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383551.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270833.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270836.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270840.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270839.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270843.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270841.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270838.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270842.1 Homo sapiens chromosome 13 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270844.1 Homo sapiens chromosome 14 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270847.1 Homo sapiens chromosome 14 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270845.1 Homo sapiens chromosome 14 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270846.1 Homo sapiens chromosome 14 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270852.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270851.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270848.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383554.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270849.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383555.2 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270850.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270854.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270856.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270855.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270853.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383556.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383557.1 Homo sapiens chromosome 16 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383563.3 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270862.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270861.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270857.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >JH159146.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >JH159147.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383564.2 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL000258.2 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383565.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270858.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270859.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383566.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270860.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270864.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383567.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383570.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383571.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383568.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383569.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383572.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270863.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270868.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270865.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383573.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383575.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383576.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383574.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270866.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270867.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL949746.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383577.2 Homo sapiens chromosome 20 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270869.1 Homo sapiens chromosome 20 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270871.1 Homo sapiens chromosome 20 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270870.1 Homo sapiens chromosome 20 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383578.2 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270874.1 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270873.1 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383579.2 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383580.2 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383581.2 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270872.1 Homo sapiens chromosome 21 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270875.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270878.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270879.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270876.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270877.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383583.2 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >GL383582.2 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270880.1 Homo sapiens chromosome X genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270881.1 Homo sapiens chromosome X genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_1
Header: >KI270892.1 Homo sapiens chromosome 1 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270894.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270893.1 Homo sapiens chromosome 2 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270895.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270896.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270897.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270898.1 Homo sapiens chromosome 5 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >GL000251.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270899.1 Homo sapiens chromosome 7 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270901.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270900.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270902.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270903.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270904.1 Homo sapiens chromosome 12 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270906.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270905.1 Homo sapiens chromosome 15 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270907.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270910.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270909.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >JH159148.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270908.1 Homo sapiens chromosome 17 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270912.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270911.1 Homo sapiens chromosome 18 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >GL949747.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KB663609.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270913.1 Homo sapiens chromosome X genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_2
Header: >KI270924.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >KI270925.1 Homo sapiens chromosome 4 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >GL000252.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >KI270926.1 Homo sapiens chromosome 8 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >KI270927.1 Homo sapiens chromosome 11 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >GL949748.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >KI270928.1 Homo sapiens chromosome 22 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_3
Header: >KI270934.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_4
Header: >GL000253.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_4
Header: >GL949749.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_4
Header: >KI270935.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_5
Header: >GL000254.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_5
Header: >GL949750.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_5
Header: >KI270936.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_6
Header: >GL000255.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_6
Header: >GL949751.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_6
Header: >KI270937.1 Homo sapiens chromosome 3 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_7
Header: >GL000256.2 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_7
Header: >GL949752.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_7
Header: >KI270758.1 Homo sapiens chromosome 6 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_8
Header: >GL949753.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_8
Header: >KI270938.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_9
Header: >KI270882.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_10
Header: >KI270883.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_11
Header: >KI270884.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_12
Header: >KI270885.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_13
Header: >KI270886.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_14
Header: >KI270887.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_15
Header: >KI270888.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_16
Header: >KI270889.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_17
Header: >KI270890.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_18
Header: >KI270891.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_19
Header: >KI270914.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_20
Header: >KI270915.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_21
Header: >KI270916.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_22
Header: >KI270917.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_23
Header: >KI270918.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_24
Header: >KI270919.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_25
Header: >KI270920.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_26
Header: >KI270921.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_27
Header: >KI270922.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_28
Header: >KI270923.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_29
Header: >KI270929.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_30
Header: >KI270930.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_31
Header: >KI270931.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_32
Header: >KI270932.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_33
Header: >KI270933.1 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_34
Header: >GL000209.2 Homo sapiens chromosome 19 genomic contig, GRCh38 reference assembly alternate locus group ALT_REF_LOCI_35
Header: >J01415.2 Homo sapiens mitochondrion, complete genome
```

</details>

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
