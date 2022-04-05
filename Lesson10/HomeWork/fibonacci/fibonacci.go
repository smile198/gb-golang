package fibonacci

type FibonacciFinder struct {
	cached map[uint]uint
}

// New creates new FibonacciFinder with default setup
func New(length int) FibonacciFinder {
	return FibonacciFinder{
		cached: make(map[uint]uint, length),
	}
}

// GetNumber returns Fibonacci number at given index
func (ff *FibonacciFinder) GetNumber(index uint) uint {
	_, exists := ff.cached[index]
	if !exists {
		ff.cached[index] = ff.calcNumber(index)
	}

	return ff.cached[index]
}

func (ff FibonacciFinder) calcNumber(number uint) uint {
	if number <= 1 {
		return number
	}

	return ff.GetNumber(number-1) + ff.GetNumber(number-2)
}
