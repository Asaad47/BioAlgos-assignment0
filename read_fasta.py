t2t_file_path = "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"

i = 0

with open(t2t_file_path, "r") as file:
    for line in file:
        if line.startswith(">"):  # Header line
            print(f"Header: {line.strip()}")
        # else:  # Sequence line
        #     print(f"Sequence: {line.strip()}")
        i += 1
        
        # if i == 15:
        #     break
print(i)

"""
(base) {9:15}~/Desktop/classes/cs249/assignment0 ➭ python read_fasta.py
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
"""