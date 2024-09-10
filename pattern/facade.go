package pattern

import (
	"fmt"
)

const wattMultiplier = 0.7355

type engine struct {
	hp    int
	price int
}

type electricEngine struct {
	powerInWatt int
	price       int
	capacity    int
}

func main() {
	electricEng := electricEngine{
		powerInWatt: 500,
		price:       8000,
		capacity:    300,
	}
	deiselCast := engine{
		hp:    electricEng.wattToHP(),
		price: electricEng.price,
	}
	fmt.Println(electricEng, deiselCast)
}

func (electric electricEngine) wattToHP() int {
	return int(float64(electric.powerInWatt) * wattMultiplier)
}
