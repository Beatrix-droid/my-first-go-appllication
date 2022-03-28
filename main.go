package main

import "fmt"
import "booking-app/helper" //to use custom packeges just import the path of the directory they are located in
import "sync"
import "time"


var wg = sync.WaitGroup{}  //waits for launched goroutine to finish
													//package "sync  provides basic Synchronization functionality"

func main(){


var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
//arrays number defines how long the array is, the type ext to it defines what type the array will contain. arrays in go cannot mix and match types.
var bookings = make([]UserData, 0) //a list/slice of maps


// a struct can be thoguht of as dictionary that can have a mixed data type
//or as a class where you specify the structure and the data type.

type UserData struct{
		username string //specify the data type for each key
		email string
		numberOfTickets uint
}

greetUsers(conferenceName, conferenceTickets, remainingTickets)



	//ask user for their name

	userName, userTickets, email := getUserInput()

	remainingTickets = remainingTickets - userTickets

	if userTickets > remainingTickets {

		fmt.Printf("We only have %v tickets remaining!", remainingTickets)
		//continue // so user has another chance to enter and book tickets.
	}

	isValidName, isValidTicketNumber, isValidEmail :=  helper.ValidateUserInput(userName, userTickets, email, remainingTickets)


	if !isValidName{
		fmt.Println("First Name you entered is too short")
	}


	if !isValidEmail{
		fmt.Println("Email address is incorrect")
	}


	if isValidEmail && isValidName && isValidTicketNumber{




		firstNames := returnName(bookings)
		bookTickets(remainingTickets, userTickets, bookings, conferenceName, userName, email)

		wg.Add(1) //add sets the number of goroutines to wait for (increases the counter by the provided number)
		//the more go routiness you wait, the more numbers in the Add() function
		go sentTicket(userTickets, userName, email)

		fmt.Printf("The names of bookings are %v \n", firstNames)

	}


	if remainingTickets == 0{
		fmt.Printf( "We have run out of tickets, come back next year!")
		//break
		}

	wg.Wait()//waitns for all the threads that were added, Blocks unit the WaitGroup counter is 0



	//an illustration of the switch syntax
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


//must tell funct the type of input paremeter it is expecting
func greetUsers(conference string, confTickets uint, remainingTickets uint){
	fmt.Printf("Welcome to our %v \n", conference)
	fmt.Printf("Get your tickets here to attend the conference\n")
	fmt.Printf("We have %v tickets and  %v remaining available tickets\n", confTickets, remainingTickets)
}

//when a func returns a value, its type must be specified after the parameters
func returnName(bookings map)(string){
	firstNames := []string{}
	for _, booking:= range bookings{
		bookings["first name"] // splits the string ith the white space as a separator, and returns a slice with split elements
		firstNames = append(firstNames, bookings.userName)
	}
	return firstNames
}


func getUserInput()(string, uint, string){
	var userName string
	var userTickets uint
	var email string

	fmt.Println("Enter your first name")
	fmt.Scan(&userName)
	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	return userName, userTickets, email

}


func bookTickets (remainingTickets uint, userTickets uint, bookings[] string, conferenceName string, userName string, email string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData) //adding elements to a list

	//create a map (python dict) for a user
	//specify type of the key and of the value
	var userData = UserData{
		username: userName,
		email: email,
		numberOfTickets: userTickets,

	}
	 // create an empty map wiht "make function"

	//adding key value paris to the map:

	userData["first name"] = userName
	userData["email"] = email
	userData["tickets booked"] = strconv.FormatUint(uint64(userTickets), 10) //convert the user tickets into a string


	fmt.Printf("Thank you %v  for booking %v tickets. You will recieve a confirmation email at %v \n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for  the %v\n", remainingTickets, conferenceName)
}

//goroutines
func sentTicket(userTickets uint, userName string, email string){
	time.Sleep(50 * time.Second)
	fmt.Printf("%v tickets for %v", userTickets, userName)
	//saving formatted strings in a variables:
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName) //to save formatted strings rather than priting them to the console
	fmt.Println("############")
	fmt.Printf("sending ticket:\n\n %v to email address %v", ticket, email)
	fmt.Println("############")

	wg.Done() //decrements the Waitgroup counter by 1. Thus this is called by the goroutine to indicate that it's finished
}

//concurrency in go is cheap and easy
//our applicaton runs in a single thread. So all code lines get executed one by one in tthis
//order in our application
//to start another thread or routien when a function is called just write "go" in front of the function. This starts a new goroutine.
// a goroutine is a lightweight thread managed by the Go runtime

//
