package main

import "fmt"

func luas(PI float32, r float32) float32 {
	return PI * (r * r)
}

func keliling(PI float32, r float32) float32 {
	return (2 * PI) * r
}

func main() {
	var r float32
	const PI float32 = 3.14
	fmt.Println("============================")
	fmt.Println("LUAS DAN KELILING LINGKARAN")
	fmt.Printf("============================ \n")
	fmt.Print("Masukan jari-jari :")
	fmt.Scanf("%g", &r)
	fmt.Printf("\nDiketahui : \nPI = %g\nr = %g\n\n", PI, r)
	fmt.Printf("Keliling = %.1f \n", keliling(PI, r))
	fmt.Printf("Luas = %.1f\n", luas(PI, r))
	fmt.Println("============================")
}
