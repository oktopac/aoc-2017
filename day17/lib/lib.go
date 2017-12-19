package day17

import "log"

type Spinlock struct {
	step     int
	position int
	count    uint
	spinlock []uint
}

func NewSpinlock(step int) *Spinlock {
	s := new(Spinlock)
	s.step = step
	s.position = int(0)
	s.count = uint(1)
	s.spinlock = make([]uint, 1)
	s.spinlock[0] = 0
	return s
}

func (s *Spinlock) insertValue(index int, value uint) {
	s.spinlock = append(s.spinlock, 0)
	copy(s.spinlock[index+1:], s.spinlock[index:])
	s.spinlock[index] = value
}

func (s *Spinlock) Spin() int {
	s.position = (s.position + s.step) % len(s.spinlock)
	s.insertValue(s.position, s.count)
	s.count++
	return s.position
}

func (s *Spinlock) Lookup(index int) uint {
	log.Println(index)
	return s.spinlock[index%len(s.spinlock)]
}
