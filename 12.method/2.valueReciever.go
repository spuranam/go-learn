package main

import "fmt"

type automobile struct {
	make  string
	model string
	year  int
	doors int
	color string
}

// In Go methods reciever can be either a value of type or pointer to value.
func (a automobile) start() {
	fmt.Printf("%s %s has started and is ready to go\n", a.make, a.model)
}

func (a automobile) stop() {
	fmt.Printf("%s %s has come to stop\n", a.make, a.model)

}

func (a automobile) honk() {
	fmt.Printf("%s %s is honking its horn\n", a.make, a.model)
}

func (a automobile) paint(c string) {
	a.color = c
	fmt.Printf("%s %s has been painted %s\n", a.make, a.model, c)

}

func main() {
	fusion := automobile{"Ford", "Fusion", 2014, 4, "White"}
	mustang := automobile{"Ford", "Mustang", 2010, 2, "Red"}

	fusion.start()
	fusion.honk()
	fusion.stop()

	mustang.start()
	mustang.honk()
	mustang.stop()

	fmt.Printf("%s %s is of %s color\n", mustang.make, mustang.model, mustang.color)
	mustang.paint("Black")
	fmt.Printf("The color of %s %s is %s\n", mustang.make, mustang.model, mustang.color)
}
