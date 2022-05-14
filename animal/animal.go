package animal

import (
	"pprof/animal/canidae/dog"
	"pprof/animal/canidae/wolf"
	"pprof/animal/felidae/cat"
	"pprof/animal/felidae/tiger"
	"pprof/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
