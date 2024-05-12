package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	email string
	noOfTickets uint
}
var wg = sync.WaitGroup{}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookingList = make([]UserData, 0)
	
	fmt.Printf("conferenceName is %T, conferenceTicket is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total %v out of which %v are available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	firstName, lastName, email, userTickets := helper.GetUserInput()

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		noOfTickets: userTickets,
	}

	if remainingTickets < userTickets {
		fmt.Printf("We only have %v tickets left, we cannot book %v tickets for you\n", remainingTickets, userTickets)
		return
	}

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation on - %v \n", firstName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	bookingList = append(bookingList, userData)

	fmt.Printf("We have %v tickets left for the conference", remainingTickets)

	firstNames := printFirstNames(bookingList)
	fmt.Printf("\nBooking list first names - %v \n", firstNames)

	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName)

	noTicketsLeft := remainingTickets == 0
	if noTicketsLeft {
		fmt.Println("Sorry we are house full!")
	}
	wg.Wait()

}

func printFirstNames(bookingList []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookingList {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames;
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets are sent to %v  %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("%v. Hope will see you there\n", ticket)
	fmt.Println("###############")
	wg.Done()
}
