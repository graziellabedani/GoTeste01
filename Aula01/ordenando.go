package main

import (
	"fmt"
	"Sort"
)

func SortByLen(arr ...string) []string {
	for i := 0; i < len(arr) - 1; i++{
		for j := i + 1; j < len(arr); j++{
			if len(arr[i]) > len(arr[j]){
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main(){
	entrada :=[]string{"maça", "ameixa", "laranja", "pera", "banana"}

	saida := SortByLen(entrada ...)

	fmt.Println("Entrada: ", entrada)
}
