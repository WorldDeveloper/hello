// package main

// import "fmt"

// func main() {
// 	words := [...]string{"Hello444444444", "World", "in", "a", "Frame"}

// 	for _, word := range words {
// 		fmt.Printf("* %-5s *\n", word)
// 	}

// 	fmt.Println(getLen("abcde"))

// }

// func getLen(src string) int {
// 	return len(src)
// }

// func findMaxWord(s) int {
// 	return
// }

// package main

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {

// 	picture := make([][]uint8, 0, 0)
// //picture := [][]uint8{}

// 	//row := make([]uint8, dx, dx)

// 	for y:= 0; y < dy; y++ {
// 	row := []uint8{}
// 	for x := 0; x < dx; x++ {
// 		row = append(row, uint8((x^y)))
// 	}
// 	picture = append(picture, row)
// 	}

// 	return picture
// }

// func main() {
// 	pic.Show(Pic)
// }
