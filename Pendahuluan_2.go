package main
import "fmt"

func main(){
var i, n, tabungan_awal, kenaikan, total, penambahan int

fmt.Scan(&n, &tabungan_awal, &kenaikan)
total = tabungan_awal
penambahan = kenaikan
for i=1; i <= n; i++{
total = total + penambahan
penambahan = penambahan + kenaikan
}

fmt.Println (total)
}
