package main

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int)
	j := 0
	for _, s := range input {
		if _, ok := m[s]; ok == false {
			m[s] = j
			j++
		}
	}
	output := make([]string, len(m))
	for k, v := range m {
		output[v] = k
	}
	return output
	
}