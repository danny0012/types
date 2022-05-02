package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

type Title string

type Population int

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	/*var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f Litres: %0.1f\n", carFuel, busFuel)*/

	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Liters(1.2) > Liters(3.4))

	fmt.Println()
	fmt.Println(Title("Alien") == Title("Alien"))
	fmt.Println(Title("Alien") < Title("Zodiac"))
	fmt.Println(Title("Alien") > Title("Zodiac"))
	fmt.Println(Title("Alien") + "s")

	fmt.Println()
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Gallons(1.2) == 1.2)
	fmt.Println(Liters(1.2) < 3.4)

	fmt.Println()
	fmt.Println(Liters(1.2) + Liters(Gallons(3.4)*3.785))
	fmt.Println(Gallons(1.2) == Gallons(Liters(1.2)*0.264))

	fmt.Println()

	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek County population:", population)

	fmt.Println()

	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
