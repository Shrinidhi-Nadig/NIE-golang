package main
import "fmt"

func main(){
	number:=20
	if number%2==0{
		fmt.Println(number ,"is Even Number")
	}else{
		fmt.Println(number, "is Odd Number")
	}
	day:="Thursday"
	switch day{
	case "Monday":
		fmt.Println("Start the week")
	case "Saturday":
		fmt.Println("End of the week")
	default:
		fmt.Println("May be a regular day")
	}
}