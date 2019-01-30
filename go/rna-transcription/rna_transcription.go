package strand

// ToRNA converts DNA string to RNA equivalent.
func ToRNA(dna string) string {
	rna := ""
	for i, _ := range dna {
		if string(dna[i]) == "G" {
			rna += "C"
		}
		if string(dna[i]) == "C" {
			rna += "G"
		}
		if string(dna[i]) == "T" {
			rna += "A"
		}
		if string(dna[i]) == "A" {
			rna += "U"
		}
	}
	return rna
}
