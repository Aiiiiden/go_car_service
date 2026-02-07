package main

import "fmt"

type Car struct {
	Model string
	Year  int
	Price int
	Color string
}

func (c Car) Info() {
	fmt.Printf("Машина: %s, Год: %d, Цвет: %s, Цена: %d", c.Model, c.Year, c.Color, c.Price)
}

func (c *Car) Paint(newColor string) {
	c.Color = newColor
}

func (c *Car) CalculateDiscount() {
	switch c.Model {
	case "Toyota","Lada":
		c.Price = c.Price - 50000
		fmt.Println("Применена скидка для популярных марок!")
	case "Mercedes":
		c.Price = c.Price +100000
		fmt.Println("Добавлена наценка за класс")
	default:
		fmt.Println("Стандартная цена для этой модели.")
	}
}

func main() {
	var myCar Car
	myCar.Model = "Toyota"
	myCar.Year = 2009
	myCar.Color = "Белый"
	myCar.Price = 2400000

	fmt.Println("Система управления автосервисом запущена:")
	myCar.CalculateDiscount()
	myCar.Paint("Red")
	myCar.Info()
}
