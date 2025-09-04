package main

import (
	"fmt"
	"strings"
)
func isAllUnique(s1 string) bool {
    hash := make(map[string]bool)
    for _, e := range s1 {
        if _, ok := hash[strings.ToLower(string(e))]; !ok {
            hash[strings.ToLower(string(e))] = true
        } else { 
			return false 
		}
    }
    return true
}

func main() {
	sentenceOne := "abcd"
	sentenceTwo := "abCdefAaf"
	sentenceThree := "aabcd"
	fmt.Println(isAllUnique(sentenceOne))
	fmt.Println(isAllUnique(sentenceTwo))
	fmt.Println(isAllUnique(sentenceThree))
}
