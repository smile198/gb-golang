package algorithm

var fibbCache = map[int]int{
	0: 0,
	1: 1,
}

func fibbWithCache(x int) int {
	if val, ok := fibbCache[x]; ok {
		return val
	}

	fibbCache[x] = fibbWithCache(x-1) + fibbWithCache(x-2)

	return fibbCache[x]
}

func fibbWithoutCache(x int) int {
	if x < 2 {
		return x
	}

	return fibbWithoutCache(x-1) + fibbWithoutCache(x-2)
}

func Fibb(x int) int {
	//return fibbWithoutCache(x)
	return fibbWithCache(x)
}
