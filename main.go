package main

import "fmt"

func luas(PI float32, r float32) float32 {
	return PI * (r * r)
}

func keliling(PI float32, r float32) float32 {
	return (2 * PI) * r
}

func main() {
	const r float32 = 12
	const PI float32 = 3.14
	fmt.Println("============================")
	fmt.Println("LUAS DAN KELILING LINGKARAN")
	fmt.Printf("============================ \n")
	fmt.Printf("Diketahui PI = %g, r = %g\n\n", PI, r)
	fmt.Printf("Keliling = %g \n", keliling(PI, r))
	fmt.Printf("Luas = %g\n", luas(PI, r))
	fmt.Println("============================")

}
