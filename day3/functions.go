package main
import "fmt"

func add(a int, b int)int{
	return a+b
}

func swap(x,y string)(string,string){
	return y,x
}
func main(){
	num1, num2 :=0, 0
	fmt.Print("Enter any two elements: ")
	fmt.Scan(&num1, &num2)

	sum:=add(num1,num2)
	fmt.Println("The sum of two numbers: ",sum)

	// name:="Shrinidhi"
	// fmt.Println("HI Shrinidhi %s",name)
	// fmt.Print("HI Shrinidhi %s",name)

	first,last:=swap("North","Nie")
	fmt.Println("Swapped: ",first,last)
}

