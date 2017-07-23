package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func MyPrint(d mat64.Matrix) {
	r, c := d.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Printf("%8.2f\t", d.At(i, j))
		}
		fmt.Println()
	}
}

func main() {
	dense1 := mat64.NewDense(5, 7, make([]float64, 35))
	dense_a := mat64.NewDense(5, 6, make([]float64, 30))
	dense_b := mat64.NewDense(6, 7, make([]float64, 42))

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			dense_a.Set(i, j, float64(i*j))
		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			dense_b.Set(i, j, float64(i*j))
		}
	}

	fmt.Println("=================Mul之前的Matrix=================")
	fmt.Println("----------dense_a------------")
	MyPrint(dense_a)
	fmt.Println("----------dense_b------------")
	MyPrint(dense_b)
	fmt.Println("----------dense1------------")
	MyPrint(dense1)

	dense1.Mul(dense_a, dense_b)

	fmt.Println("=================Mul之后的Matrix=================")
	fmt.Println("----------dense_a------------")
	MyPrint(dense_a)
	fmt.Println("----------dense_b------------")
	MyPrint(dense_b)
	fmt.Println("----------dense1------------")
	MyPrint(dense1)

}
