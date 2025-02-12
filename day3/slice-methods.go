package main
import "fmt"

func main(){
	slice :=[]int{3,6,9}
	fmt.Println("Before appending new elements ,Slice: ",slice)

	//APPEND THE NEW ELEMENTS
	slice=append(slice,12,15)
	slice=append(slice,15,18)
	fmt.Println("After append new elements ,Slice: ",slice)
	fmt.Println("Length of slice is:",len(slice))
	
	//COPY SLICE
	newslice:=make([]int,6)
	copyslice:=copy(newslice,slice)
	fmt.Println("The existing slice elements are: ",slice)
	fmt.Println("The newly copied slice elements are: ",newslice)
	fmt.Println("Number of elements copied are: ",copyslice)

	//CLEAR ELEMENTS
	newslice=nil
	fmt.Println("The newslice is empty: ",newslice)
}