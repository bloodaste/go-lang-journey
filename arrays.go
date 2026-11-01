package main

import "fmt"

func main() {
	arry := [5]int{}

	for i := 0; i < len(arry); i++ {
		arry[i] = i
	}
	fmt.Println(arry)
	fmt.Println(arry[2])

	arry2 := [5]int {1,2,3,4,5}
	var sums int
	for i :=0; i<len(arry2); i++ {
		fmt.Println("arry shit",arry2[i])
		sums += arry2[i]
	}
	fmt.Println("output",sums)
	fmt.Println(arry2[0])

	var valueindex int
	for index, value := range arry2 {
		valueindex += value
		fmt.Printf("the value of %d, is %d \n", index,value)
	}
	fmt.Println(valueindex)
	_,b  :=num();

	fmt.Println("number a",b)

	multidementionalarray := [3][3] int{}
	
	for i:=0; i<4; i++{
		for j := i+1; j<3; j++{
		multidementionalarray[i][j] = i+j 
		}
	}
	fmt.Println(multidementionalarray)
}

func num ()(int ,int){
	return 1,2
}