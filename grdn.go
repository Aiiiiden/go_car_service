package main

import "fmt"

type Palm struct {
	Type   string
	Yield  int
	Health int
}

func (p *Palm) Water() {
	p.Health += 10
	fmt.Printf("Вы полили сорт %s. Здоровье теперь %d%%\n", p.Type, p.Health)
}

func main() {
	fmt.Println("--- Добро пожаловать в твой Сад в Медине ---")

	myGarden := []Palm{
		{Type: "Меджул", Yield: 50, Health: 80},
		{Type: "Аджва", Yield: 30, Health: 95},
		{Type: "Суккари", Yield: 40, Health: 40},
	}

	newPalm :=Palm{Type: "Мазфати", Yield: 60, Health: 90}
	myGarden = append(myGarden, newPalm)
	fmt.Println("--- Добавлено новое дерево сорта Мазфати! ---")

	totalYield := 0

	for i := 0; i < len(myGarden); i++ {
		switch {
		case myGarden[i].Health < 50:
			fmt.Printf("!!! Дерево %s требует внимания!\n", myGarden[i].Type)
			myGarden[i].Water()
		default:
			fmt.Printf("Дерево %s в порядке.\n", myGarden[i].Type)
		}

		totalYield = totalYield + myGarden[i].Yield
	}

	fmt.Println("---------------------------------------")
	fmt.Printf("Итого урожая со всех деревьев %d кг\n", totalYield)
	fmt.Println("Бисмилягь, работа на сегодня закончена.")
}
