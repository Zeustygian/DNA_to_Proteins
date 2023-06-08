package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

func translate_strand(strand byte) byte {
	if strand == 'a' {
		return 't'
	}
	if strand == 't' {
		return 'a'
	}
	if strand == 'g' {
		return 'c'
	}
	if strand == 'c' {
		return 'g'
	}
	return ' '
}

func translate_to_mrna(strand byte) byte {
	if strand == 'a' {
		return 'u'
	}
	if strand == 't' {
		return 'a'
	}
	if strand == 'g' {
		return 'c'
	}
	if strand == 'c' {
		return 'g'
	}
	return ' '
}

func translate_to_trna(strand byte) byte {
	if strand == 'a' {
		return 'u'
	}
	if strand == 'u' {
		return 'a'
	}
	if strand == 'g' {
		return 'c'
	}
	if strand == 'c' {
		return 'g'
	}
	return ' '
}

func dna_complementary(dna string) string {
	var c_dna bytes.Buffer

	for i := 0; i < len(dna); i++ {
		c_dna.WriteByte(translate_strand(dna[i]))
	}
	return c_dna.String()
}

func message_arn(dna string) string {
	var m_rna bytes.Buffer

	for i := 0; i < len(dna); i++ {
		m_rna.WriteByte(translate_to_mrna(dna[i]))
	}
	return m_rna.String()
}

func mrna_to_trna(m_rna string) string {
	var t_rna bytes.Buffer

	for i := 0; i < len(m_rna); i++ {
		t_rna.WriteByte(translate_to_trna(m_rna[i]))
	}
	return t_rna.String()
}

func get_protein(m_rna string) string {
	var array map[string]string = make(map[string]string)

	array["uuu"] = "Phe"
	array["uuc"] = "Phe"
	array["uua"] = "Leu"
	array["uug"] = "Leu"

	array["cuu"] = "Leu"
	array["cuc"] = "Leu"
	array["cua"] = "Leu"
	array["cug"] = "Leu"

	array["auu"] = "Ile"
	array["auc"] = "Ile"
	array["aua"] = "Ile"
	array["aug"] = "Met"

	array["guu"] = "Val"
	array["guc"] = "Val"
	array["gua"] = "Val"
	array["gug"] = "Val"

	array["ucu"] = "Ser"
	array["ucc"] = "Ser"
	array["uca"] = "Ser"
	array["ucg"] = "Ser"

	array["ccu"] = "Pro"
	array["ccc"] = "Pro"
	array["cca"] = "Pro"
	array["ccg"] = "Pro"

	array["acu"] = "Thr"
	array["acc"] = "Thr"
	array["aca"] = "Thr"
	array["acg"] = "Thr"

	array["gcu"] = "Ala"
	array["gcc"] = "Ala"
	array["gca"] = "Ala"
	array["gcg"] = "Ala"

	array["uau"] = "Tyr"
	array["uac"] = "Tyr"
	array["uaa"] = "STOP"
	array["uag"] = "STOP"

	array["cau"] = "His"
	array["cac"] = "His"
	array["caa"] = "Gln"
	array["cag"] = "Gln"

	array["aau"] = "Asn"
	array["aac"] = "Asn"
	array["aaa"] = "Lys"
	array["aag"] = "Lys"

	array["gau"] = "Asp"
	array["gac"] = "Asp"
	array["gaa"] = "Glu"
	array["gag"] = "Glu"

	array["ugu"] = "Cys"
	array["ugc"] = "Cys"
	array["uga"] = "STOP"
	array["ugg"] = "Trp"

	array["cgu"] = "Arg"
	array["cgc"] = "Arg"
	array["cga"] = "Arg"
	array["cgg"] = "Arg"

	array["agu"] = "Ser"
	array["agc"] = "Ser"
	array["aga"] = "Arg"
	array["agg"] = "Arg"

	array["ggu"] = "Gly"
	array["ggc"] = "Gly"
	array["gga"] = "Gly"
	array["ggg"] = "Gly"

	return array[m_rna]
}

func split_three(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func loop_proteins(m_rna string) {
	arr := split_three(m_rna, 3)

	for i := 0; i != len(arr); i++ {
		print(" " + get_protein(arr[i]))
	}
}

func main() {
	body, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("Unable to read the file: %v", err)
	}
	dna := strings.ToLower(string(body))
	// fmt.Println("DNA: " + string(dna))
	c_dna := dna_complementary(string(dna))
	// fmt.Println("DNA complementary: " + c_dna)
	m_rna := message_arn(c_dna)
	// fmt.Println("Messager RNA: " + m_rna)
	// t_rna := mrna_to_trna(m_rna)
	// fmt.Println("Transfer RNA: " + t_rna)

	// os.Exit(0)

	print("\nThe proteins synthesized by this RNA sample are:")
	loop_proteins(m_rna)
}
