# solve exact match




def find_num_exact_matches(seq_file_name, pattern_file_name):
    pattern = ""
    with open(pattern_file_name, "r") as file:
        for line in file:
            if line.startswith(">"):  # Header line
                print(f"Pattern-- Header: {line.strip()}")
            else:  # Sequence line
                pattern += line.strip()
                
    print("pattern: ", pattern)
    print("length of pattern: ", len(pattern))
    num_matches = 0
    chromosome = ""
    valid_chromosome = False
    start = 0
    
    queue = []
    

    with open(seq_file_name, "r") as file:
        for line in file:
            if line.startswith(">"):  # Header line
                print(f"Header: {line.strip()}")
                print(f"-- Number of matches so far: {num_matches}")
                start = 0
                queue = []
                chromosome = line.strip().split("chromosome ")
                if len(chromosome) > 1:
                    valid_chromosome = True
                    chromosome = chromosome[1]
                else:
                    chromosome = chromosome[0]
                    valid_chromosome = False
            else:  # Sequence line
                # print(f"Sequence: {line.strip()}")
                # if not valid_chromosome:
                #     continue
                queue += list(line.strip())

                while len(queue) >= len(pattern):
                    matched = True
                    for i in range(len(pattern)):
                        if queue[i].lower() != pattern[i].lower():
                            matched = False
                            break
                    if matched:
                        num_matches += 1
                        print(f"Matched: (chromosome, start, end) = ({chromosome}, {start}, {start + len(pattern) - 1})")
                    queue.pop(0)
                    start += 1
                    
                    
    
    return num_matches
        
        
if __name__ == "__main__":
    t2t_file_path = "ncbi_dataset_T2T/ncbi_dataset/data/GCA_009914755.4/GCA_009914755.4_T2T-CHM13v2.0_genomic.fna"
    
    alu_file_path = "DF000000002.fa"
    
    
    print(find_num_exact_matches(t2t_file_path, alu_file_path))