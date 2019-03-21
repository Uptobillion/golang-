package inandout

import "fmt"

func In() []int {
	var cl2 = make([]int, 3)
	fmt.Println("Please input x1 y1 and boomsnumber:")
	fmt.Scanf("%d %d %d", &cl2[0], &cl2[1], &cl2[2])
	return cl2
}
func Out(x1, y1 int, a, b [][]int) {
	fmt.Printf("   ")
	for x := 1; x < x1+1; x++ {
		fmt.Printf("%02d ", x)
	}
	fmt.Printf("\n")
	for y := 1; y < y1+1; y++ {
		fmt.Printf("%02d ", y)
		for x := 1; x < x1+1; x++ {
			switch {
			case b[x][y] == 1:
				fmt.Printf("%02d ", a[x][y])
			case b[x][y] == 2:
				fmt.Printf("f  ")
			case b[x][y] == 0:
				fmt.Printf("x  ")
			}
		}
		fmt.Printf("\n")
	}
}

func Click() []int {
	var cl2 = make([]int, 3)
	fmt.Printf("Please input x y c:\n")
	fmt.Scanf("%d %d %d", &cl2[0], &cl2[1], &cl2[2])
	return cl2
}
func Fail() {
	fmt.Println("failed!!")
}
func Success() {
	fmt.Println("yes!")
}
