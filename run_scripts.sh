#!/bin/bash

prob1_naive() {
    echo "Running commands for prob1_naive"
    /usr/bin/time -l go run prob1_naive.go ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
    /usr/bin/time -l go run prob1_naive.go ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna

}

prob1_kmp() {
    echo "Running commands for prob1_kmp"
    /usr/bin/time -l go run prob1_kmp.go ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
    /usr/bin/time -l go run prob1_kmp.go ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
}

prob2() {
    echo "Running commands for prob2"
    /usr/bin/time -l go run prob2.go ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna
    /usr/bin/time -l go run prob2.go ncbi_dataset_GRCh38/ncbi_dataset/data/GCA_000001405.29/GCA_000001405.29_GRCh38.p14_genomic.fna
}


if [[ $# -eq 1 ]]; then
    case "$1" in
        prob1_naive)
            prob1_naive
            ;;
        prob1_kmp)
            prob1_kmp
            ;;
        prob2)
            prob2
            ;;
        *)
            echo ">>> Running all commands."
            prob1_naive
            prob1_kmp
            prob2
            ;;
    esac
else
    echo ">>> Running all scripts."
    prob1_naive
    prob1_kmp
    prob2
fi
