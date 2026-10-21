package main
import "fmt"

func main() {
	var n, k, p int

	fmt.Scan(&n)
	p = n
	for k = 1; k <= n; k++ {
		fmt.Println("Request", k, "diproses dengan prioritas", p)
		p--
	}
}
