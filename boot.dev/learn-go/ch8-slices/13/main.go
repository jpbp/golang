package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i:=0;i<rows;i++{

		var linha [] int
		for j:=0;j<cols;j++{
			linha = append(linha, i*j)
		}
		matrix = append(matrix, linha)
	}
	return matrix
}
func main(){
	m1 := createMatrix(5,10)
	for _, linha := range m1 {
		fmt.Println(linha)
	}
}