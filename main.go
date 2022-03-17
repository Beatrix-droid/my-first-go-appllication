package main

import "fmt"

func main(){


var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50

fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
fmt.Printf(" Get your tickets here to attend the conference")
fmt.Printf("We have %v tickets and  %v remaining available tickets\n", conferenceTickets, remainingTickets)


var userName string
var userTickets uint
var email string


//ask user for their name
fmt.Println("Enter your first name")
fmt.Scan(&userName)

fmt.Println("Enter the number of tickets")
fmt.Scan(&userTickets)


fmt.Println("Enter your email address")
fmt.Scan(&email)


remainingTickets = remainingTickets - userTickets
fmt.Printf("Thank you %v  for booking %v tickets. You will recieve a confirmation email at %v \n", userName, userTickets, email)


//arrays number defines how long the array is, the type ext to it defines what type the array will contain. arrays in go cannot mix and match types.
var bookings[50] string

bookings[0] = userName

fmt.Printf("The whole array %v \n", bookings)

}
