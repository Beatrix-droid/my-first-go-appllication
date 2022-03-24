package helper

import "strings"

//to export functions and variables capitalise their name

func ValidateUserInput(userName string, userTickets uint, email string, remainingTickets uint)(bool, bool, bool) {
	var isValidName bool = len(userName) >=2
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets
	var isValidEmail bool = strings.Contains(email, "@") //contains function is like "in" in python
	return isValidName,  isValidTicketNumber,  isValidEmail

}
