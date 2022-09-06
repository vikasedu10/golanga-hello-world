package main
import "fmt"

func main()  {
	var appName = "Go Conference"
	const totalConferenceTickets uint = 50
	var remainingTickets = totalConferenceTickets

	fmt.Printf("Welcome to %v booking application!\n", appName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", totalConferenceTickets, remainingTickets)
	fmt.Println("Book your tickets now.")

	var username string;
	var email string;
	var userTickets uint;

	fmt.Printf("Enter your name to book tickets: ")
	fmt.Scan(&username)

	fmt.Printf("Enter no of tickets you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Enter your email address to book tickets: ")
	fmt.Scan(&email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thankyou %v for booking %v tickets. You will receive a confirmation e-mail at %v\n", username, userTickets, email)
	fmt.Printf("%v tickets now remaining for %v!\n", remainingTickets, appName)
}
