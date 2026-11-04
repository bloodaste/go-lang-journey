package main

import (
	"fmt"
	"slices"
)

func main() {
	// var slice1 []int
	numbers := []int{1, 2, 4, 5, 6, 7}
		numbers2 := []int{1, 2, 4, 5, 6, 7}
	a := numbers[1:4] // it limits the value of copied slice
	fmt.Println(numbers[2])
	fmt.Println(a)
	a = append(a,6,7,8) // this part will append on the sliceor array 
	fmt.Println(a)
	sum := 0
	for _,v := range numbers{
		sum += v 
	}
	fmt.Println(sum)
	sum2 := 0
	
	
	for {
		sum2 *= sum2
		sum2 ++
		println("sum",sum2)
		if (sum2 >= 10){
			break
		}
	}
	if slices.Equal (numbers2,numbers){
		fmt.Println("is equal")
	}else {
		fmt.Println("not equak")
	}

	copyslice := make([]int, len(numbers2))

	fmt.Println(copyslice)

	//multidimensional array
	twod := make([][]int, 3) // this part allows you to make an slice with 3 rows and 3 collumns
	fmt.Println(twod) // the output of this is [[],[],[]]
	for i := 0; i < 3; i++{
		innerlenght := i+1 //this 
		twod[i] = make([]int, innerlenght) // this part will create the inner lenght slice
		for j := 0; j<innerlenght; j++ {
			twod[i][j] = j+1 // this part is the value of it 
		}
	}
	fmt.Println(twod)
	secondd := make([][]int , 3)

	for i:= 0; i<3; i++{
		secondd[i] = make([]int, 3)
	
	}
	fmt.Println(secondd)
}