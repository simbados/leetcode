package main

func mergeStringsAlternately(word1 string, word2 string) string {
    returnVal := ""
    highest := max(len(word1), len(word2))
    for i := 0; i < highest; i++ {
        if i < len(word1) {
            returnVal = returnVal + string(word1[i])
        }
        if i < len(word2) {
            returnVal = returnVal + string(word2[i])
        }
    }
    return returnVal
}
