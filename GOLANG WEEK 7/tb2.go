package main
import "fmt"

func main() {
	var n, j int
	
	fmt.Scan(&n)
	for j = 1; j <= n; j += 1 {
		fmt.Print(j*j, " ")
	}
}