package main
import "fmt"

func main() {
	var n, i, j int
	
	fmt.Scan(&i, &n)
	for j = i; j <= n; j+=2 {
		fmt.Println(j)
	}
}
