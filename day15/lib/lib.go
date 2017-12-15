package day15

func exhaustGenerator(generator <-chan uint64) []uint64 {
	ret := make([]uint64, 0)
	i := 0
	for val := range generator {
		ret = append(ret, val)
		i++
	}
	return ret
}

func generator(start uint64, factor uint64, number uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		previous := start
		defer close(ch)
		for i := uint64(0); i < number; i++ {
			previous = (previous * factor) % uint64(2147483647)
			ch <- previous
		}
	}()
	return ch
}

func generator2(start uint64, factor uint64, number uint64, mod uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		previous := start
		defer close(ch)
		for i := uint64(0); i < number; {
			previous = (previous * factor) % uint64(2147483647)
			if previous%mod == 0 {
				ch <- previous
				i++
			}
		}
	}()
	return ch
}

func GeneratorA(start uint64, number uint64) <-chan uint64 {
	return generator(start, 16807, number)
}

func GeneratorB(start uint64, number uint64) <-chan uint64 {
	return generator(start, 48271, number)
}

func GeneratorA2(start uint64, number uint64) <-chan uint64 {
	return generator2(start, 16807, number, 4)
}

func GeneratorB2(start uint64, number uint64) <-chan uint64 {
	return generator2(start, 48271, number, 8)
}

func match(val1 uint64, val2 uint64) bool {
	return val1<<48 == val2<<48
}

func CountMatches(generator1 <-chan uint64, generator2 <-chan uint64) uint64 {
	sum := uint64(0)
	for v1 := range generator1 {
		v2 := <-generator2
		if match(v1, v2) {
			sum++
		}
	}
	return sum
}
