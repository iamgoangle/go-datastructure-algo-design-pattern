package main

import "fmt"

type CoffeeMachineHandler func(f *FrenchPressCoffeeOptions)

type CoffeeMachine interface {
	Execute(fn ...CoffeeMachineHandler)
}

type FrenchPressCoffeeOptions struct {
	Level      int
	Minutes    int
	IsHotWater bool
}

type Coffee struct {
	Name string
	FrenchPressCoffeeOptions
}

func NewCoffeeOrder(n string) CoffeeMachine {
	return &Coffee{
		Name: n,
	}
}

func (c *Coffee) Execute(fn ...CoffeeMachineHandler) {
	for _, configFn := range fn {
		configFn(&c.FrenchPressCoffeeOptions)
	}
	fmt.Println("Execute")
}

// --- library code

// production code
func SetCoffeeLevel(l int) CoffeeMachineHandler {
	return func(f *FrenchPressCoffeeOptions) {
		f.Level = l
	}
}

func main() {
	americano := NewCoffeeOrder("Americano")
	americano.Execute(SetCoffeeLevel(5))

	fmt.Printf("%+v", americano)
}
