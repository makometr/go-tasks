package task8

// Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0.

// Task8 sets bit on bitPos on n to val
func Task8(n *int64, bitPos int, val bool) {
	switch val {
	case false:
		*n &= ^(1 << bitPos)
	case true:
		*n |= (1 << bitPos)
	}
}
