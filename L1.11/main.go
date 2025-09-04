package main
import (
	"fmt"
)

func intersection(s1, s2 []int) (inter []int) {
    hash := make(map[int]bool)
    for _, e := range s1 {
        hash[e] = true
    }
    for _, e := range s2 {
        if hash[e] {
            inter = append(inter, e)
        }
    }
    inter = removeDups(inter)
    return
}
func removeDups(elements []int)(nodups []int) {
    encountered := make(map[int]bool)
    for _, element := range elements {
        if !encountered[element] {
            nodups = append(nodups, element)
            encountered[element] = true
        }
    }
    return
}

func main() {
	fmt.Println(intersection([]int{1,2,3}, []int{2,3,4}))
}