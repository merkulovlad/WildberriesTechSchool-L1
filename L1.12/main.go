package main
import (
	"fmt"
)

func makeSet(s1 []string) (uniq []string) {
    hash := make(map[string]bool)
    for _, e := range s1 {
        if _, ok := hash[e]; !ok {
            uniq = append(uniq, e)
            hash[e] = true
        }
    }
    
    return uniq
}
func main() {
	fmt.Println(makeSet([]string{"apple", "banana", "apple", "orange"}))
}