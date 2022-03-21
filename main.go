package main

import "fmt"
import "strings"

func main(){


var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50

fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
fmt.Printf(" Get your tickets here to attend the conference\n")
fmt.Printf("We have %v tickets and  %v remaining available tickets\n", conferenceTickets, remainingTickets)



for{
	var userName string
	var userTickets uint
	var email string


	//ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&userName)

	var isValidName bool = len(userName) >=2


	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v  for booking %v tickets. You will recieve a confirmation email at %v \n", userName, userTickets, email)

	if userTickets > remainingTickets {

		fmt.Printf("We only have %v tickets remaining!", remainingTickets)
		continue // so user has another chance to enter and book tickets.
	}



	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	var isValidEmail bool = strings.Contains(email, "@") //contains function is like "in" in python


	if !isValidName{
		fmt.Println("First Name you entered is too short")
	}


	if !isValidEmail{
		fmt.Println("Email address is incorrect")
	}


	if isValidEmail && isValidName && isValidTicketNumber{
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Thank you %v  for booking %v tickets. You will recieve a confirmation email at %v \n", userName, userTickets, email)


		//arrays number defines how long the array is, the type ext to it defines what type the array will contain. arrays in go cannot mix and match types.
		var bookings[] string

		bookings = append(bookings, userName, email) //adding elements to a list

		fmt.Printf("The whole array %v \n", bookings)
		fmt.Printf("%v tickets remaining for  the %v\n", remainingTickets, conferenceName)


		firstNames := []string{}
		for _, booking:= range bookings{
			var names = strings.Fields(booking) // splits the string ith the white space as a separator, and returns a slice with split elements
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("These are all our bookings %v\bn", bookings)
		fmt.Printf("The first names of bookings are %v \n", firstNames)

	}



	if remainingTickets == 0{
		fmt.Printf( "We have run out of tickets, come back next year!")
		break
		}

	}

	city := "London"

	switch city {

		case "New York":
		// execute code for booking New York conference tickets

		case " Singapore":
		// execute code for booking Singapore conference tickets

		case "London":
		// execute code for booking Singapore conference tickets

		case "Berlin":
		// some code here

		default:
			fmt.Print("No valid city selected")


	}

}
