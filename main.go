//package main//

//---------------------------------------Basic Startup To A File---------------------------------------//
//all our code must belong to a package
//the first line of code in our go file should be this
//when we create a go project we must initialise it with "go mod init (name of the folder of the project)"

//---------------------------------------How To Import Module's---------------------------------------//
//To import a module we have to import it using import "the name".
//fmt stands for format and allows us to do basic things like print hello to the cmd.
//import "fmt"//

//---------------------------------------Starting Point Of The File---------------------------------------//
//The go language need a starting point in every file so the func main is the starting point of every file.
//A program can only have one main function as there is only one entrypoint.

//---------------------------------------Make The File Work---------------------------------------//
//func main() {//
//this line accesses the fmt module import before and using Print to say something to the terminal.
//fmt.Println("Hello, World")//
//}//

//---------------------------------------Packages---------------------------------------//
//GO files are organized into packages. Go has a standard library, which provides different core packages for us to use.
//A package is a collection of source files.

//---------------------------------------Compiling---------------------------------------//
//To run the file in the terminal use "go run (the name of the file)"

//---------------------------------------Booking Application---------------------------------------//

package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	//---------------------------------------Variables---------------------------------------//

	//Variables are used to store values.
	//They are like containers for values.
	//You give the variable a name & referance anywhere in the app.
	//Benefits: you only have to update the variable once.
	//			Make our app more flexible.

	//Example: var conferenceName = "Python Conference"
	//We have to let go know that we are creating a variable using the var keyword.
	//A keyword is a reserved word in every language.
	//The code below will give us an error because in go the variable created must always be used.

	//var conferenceName = "Go Conference"
	//fmt.Println(conferenceName)

	//------Constants------//

	//Constants are variables, except that they cannot be changed
	//To make a variable a constant all you have to do is instead of using "var" use "const"

	//---------------------------------------Print Formated Data---------------------------------------//

	//Example = fmt.printf("Some text with a variable %s", myVariabel)
	//It takes a template string that contains the text that need to be formated.
	//Plus some annotation verbs (or place holders) that tells the fmt functions how to format the variable passed.

	//Example use in our code:

	//fmt.Printf("Welcome to our %v conference booking application\n", conferenceName)
	//fmt.Println("We have a total of %v tickets and %v tickets left\n", conferenceTickets, remainingTickets)
	//fmt.Println("Get your tickets here to attend.")

	//%v is a place holder that we use in the place the variable needs to be printed.
	//But we still have to refernce the variable.
	//To do that we use a comma and the name of the variable after the quotation line.
	//We also need to use "\n" to create a new line so that they are not printed together.
	//If we have two variable to be printed and want to use printf, we put %v in both places and then we reference them in the end IN ORDER.
	//%v is the default format but there are other formats.

	//---------------------------------------Data Types In GO---------------------------------------//

	//There are  5 main data types.
	//They are Strings, Integers, Booleans, Maps and Arrays
	//The most basic types are Strings and Integers
	//Strings: Used for textual data, defined with double quotes, e.g "This is a string"
	//Integers: Used to represent WHOLE NUMBERS, positive and negative e.g: 5, 120, -20
	//There are many more numeric data types.

	//Go is a statically typed language.
	//You need to tell the GO compiler, the data type when declaring the variable.
	//Go also has type inference meaning it can also infer the data type of a variable when you assign it.
	//Example where you would have to manulayy declare the variable's data type:

	//var userName string
	//varu userTickets int

	//userName = "Tom"
	//userTickets = 2
	//fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
