package main

import (
	"fmt"
	"strings"
)

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

	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets you want to buy")
		fmt.Scan(&userTickets)

		if remainingTickets < userTickets {
			fmt.Printf("We only have %v tickets left, we cannot book %v tickets for you\n", remainingTickets, userTickets)
			continue
		}

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation on - %v \n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets

		bookingList = append(bookingList, firstName + " " + lastName)

		fmt.Printf("We have %v tickets left for the conference", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookingList {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Booking list first names - %v \n", firstNames)

		noTicketsLeft := remainingTickets == 0
		if noTicketsLeft {
			fmt.Println("Sorry we are house full!")
			break
		}
	}

}
