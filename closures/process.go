package closures

func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
