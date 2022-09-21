package main

func main() {

}
func mul(a, b [4][4]int) [4][4]int {
	c := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0
			for k := 0; k < 4; k++ {
				sum += a[i][k] * b[k][j]
			}
		}
	}
}
