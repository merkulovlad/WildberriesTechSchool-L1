package main 
import (
	"fmt"
	"strings"
)
var justString string = "Привет, мир!"

// func someFunc() {
//   v := createHugeString(1 &lt;&lt; 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }
// это пример утечки памяти из-за среза. Так как срез justString ссылается на массив v,
// то весь массив v не может быть удален сборщиком мусора, даже если он больше не нужен. Так что если v очень большой,
// то память будет утекать очень сильно.
// Чтобы избежать этой утечки, можно создать новый срез, который копирует только нужные данные:
// так же string хранит байты, а не руны, поэтому 100 байт может быть меньше 100 рун


// корректная реализация
func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = string([]byte(v)[:100])
}

func firstNRunes(n int) string {
	s := []rune(justString)
	if len(s) < n {
		return justString
	}
	return string(s[:n])
}

func main() {
  fmt.Println(firstNRunes(7))
  someFunc()
  fmt.Println(firstNRunes(7))
}
