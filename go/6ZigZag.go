package main

func convert(s string, numRows int) string {
	result := ""
	multi := 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i-multi < len(s); j = j + (numRows + (numRows - 2)) {
			if i != 0 && i != numRows-1 && i+j-multi > 0 {
				result += string(s[i+j-multi])
			} 
      if j+i < len(s) {
				result += string(s[i+j])
			}
		}
		if i != 0 && i != numRows-1 {
			multi += 2
		}
	}
	return result
}
