package main

import (
	"fmt"
)

func findTriplets(arr []int, n int){
	found := true
	var i,j,k int
	
	for i = 0; i < n-2; i++ { 
	    
            for j = i+1; j < n-1; j++ { 
		
                for k = j+1; k < n; k++ { 
                    if arr[i] + arr[j] + arr[k] == 0 {
		        fmt.Println(arr[i], arr[j], arr[k])
                        found = true; 
                    } 
                } 
            } 
        } 
 
     if found == false {
	fmt.Println("not exist ===>>")
     }      
}

func main() {
	arr := []int{0, -1, 2, -3, 1}
	n := len(arr)
	findTriplets(arr, n)
}
