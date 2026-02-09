package main

import "fmt"

type Ticket struct {
	To string
}

func (t Ticket) Info() {
	fmt.Printf("Билет куплен до города: %s \n", t.To )
}

func SayBismillah() {
	fmt.Println("Бисмиллягьи Рах1мани Рах1им! В добрый путь с Богом.")
}

func main() {
	myTicket := Ticket{To: "Медина"}

	myTicket.Info()

	SayBismillah()
}