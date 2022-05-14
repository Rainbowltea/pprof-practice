package muridae

import "pprof/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
