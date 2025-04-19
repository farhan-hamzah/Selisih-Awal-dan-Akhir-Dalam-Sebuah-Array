package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, akhir, awal, hasil int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
		akhir = A[i]
	}
	awal = A[0]
	hasil = awal - akhir
	if hasil < 0{
		hasil = -hasil
	}
	fmt.Print(hasil)
}