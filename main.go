package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	
	fmt.Printf("conferenceName is %T, conferenceTicket is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total %v out of which %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")


	var firstName string
	var lastName string
	var email string
	var userTickets int
	bookingList := []string{}

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to buy")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation on - %v \n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	bookingList = append(bookingList, firstName + " " + lastName)

	fmt.Printf("We have %v tickets left for the conference", remainingTickets)
	fmt.Printf("Booking list - %v \n", bookingList)

}
