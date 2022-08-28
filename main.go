package main
import "fmt"

func main()  {
	fmt.Println("Hello DevOps world!! :)")

	var isActive bool = true
	fmt.Printf("Number of DevOps engineers are active and allrounder? - %T \n", isActive)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable type: %T \n", smallVal)

	// Default values and aliases
	var message string;
	fmt.Println(message)
	var marks int;
	fmt.Println(marks)
}