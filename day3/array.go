package main
import "fmt"

// func main(){
// 	var arr=[5]int{2,4,6,8,10}
// 	fmt.Println("Array elements are: ",arr)

// 	slice:=[]int{1,2,3,4,5}
// 	fmt.Println("Slice contains: ",slice)
// }

func main(){
	var num int
	fmt.Print("Enter the number of array elements");
	fmt.Scan(&num)
	even:=make([]int,num)
	fmt.Println("Enter ",num,"Even numbers: ")

	for i:=0;i<num;i++{
		fmt.Scan(&even[i])
	}
	fmt.Println("Even numbers are as follows: ",even)
}