package main

import "fmt"

func main() {
	fmt.Println("hello welcome to this game!")
	fmt.Printf("what is your name?: ")
	host := "sogi"
	var name string
	fmt.Scan(&name)
	fmt.Printf("nice to see you %v. This is %v \n", name, host)
	fmt.Printf("enter your age: ")
	var age int
	fmt.Scan(&age)
	if age >= 16 {
		fmt.Println("you can play the game!")
	} else {
		fmt.Println("you cannot play the game")
		return
	}
	fmt.Printf("What is 3 + 5?: ")
	var answer string
	fmt.Scan(&answer)
	if answer == "8" {
		fmt.Println("you are correct sekai des")
	} else {
		println("you are wrong sedly we gotta end this game here! learn basic calculations firs then cum again.")
		return
	}
	fmt.Println("game comming soon")
}
