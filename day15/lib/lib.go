package day15

func exhaustGenerator(start uint, factor uint, number uint) []uint {
	ret := make([]uint, number)
	i := 0
	for val := range generator(start, factor, number) {
		ret[i] = val
		i++
	}
	return ret
}

func generator(start uint, factor uint, number uint) <-chan uint {
	ch := make(chan uint)
	go func() {
		previous := start
		defer close(ch)
		for i := uint(0); i < number; i++ {
			previous = (previous * factor) % uint(2147483647)
			ch <- previous
		}
	}()
	return ch
}

func generatorA(start uint, number uint) <-chan uint {
	return generator(start, 16807, number)
}

func generatorB(start uint, number uint) <-chan uint {
	return generator(start, 48271, number)
}
