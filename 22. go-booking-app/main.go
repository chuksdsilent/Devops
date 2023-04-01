package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)

const conferenceTickets = 50
var remainingTicket uint = 50
var bookings = []string{}
var name = "Go Conference"

func main(){
	


	greetUser(name, conferenceTickets, remainingTicket)
	

	// isValidCity := city == "singapore" || city == "london"

	for len(bookings) < 50 && remainingTicket > 0{

		firstName, lastName, email, userTicket :=  getUserInput();
		bookings = append(bookings, firstName + " " + lastName)
	
		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)
		if(isValidEmail && isValidName && isValidTicketNumber){
			bookTicket(remainingTicket, userTicket, bookings)
			fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, userTicket, email)
			fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTicket , "are still available")
			fmt.Println("Get your tickets here to attend")
			if remainingTicket == 0 {
				fmt.Printf("Our conference is booked out. Come back next year \n")
				break
			}
		} else if userTicket == remainingTicket{} else {
			if !isValidName {
				fmt.Println("Your name is invalid")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
		}
		
	}

}

func greetUser(name string, conferenceTickets uint, remainingTicket uint){
	fmt.Printf("Welcome to %v booking application \n", name)
	
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceTickets)

}

func printFirstName(bookings []string) []string{
	firstNames := []string{}
	for _ , booking := range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint){


	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(remainingTicket uint, userTicket uint, bookings []string){
	remainingTicket = remainingTicket - userTicket
	firstName := printFirstName(bookings)
	fmt.Printf("The first names of bookings are: %v\n", firstName)
}