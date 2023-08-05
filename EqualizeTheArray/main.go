package main

import "fmt"

func equalizeArray(arr []int32) int32 {
    numbers := make(map[int32]int32)
	for _ , number := range arr{
		if _, isIn := numbers[number]; isIn {
			numbers[number] = numbers[number] + int32(1)
		}else{
			numbers[number] = 1
		}
	}

	major := int32(0)

	for _ , repeatNumber := range numbers {
		if repeatNumber > major {
			major = repeatNumber
		}
	}
	
	return int32(len(arr)) - major 
}


func main(){
	array := []int32{1,2,2,3}
	fmt.Println(equalizeArray(array))
}