package task10

// 10.	Дана последовательность температурных колебаний (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5).
// Объединить данный значения в группы с шагов в 10 градусов. Последовательность в подмножностве не важна.

// Task10 temperatures diaposon div 10
func Task10(ts ...int) map[int][]int {
	temps := make(map[int][]int)

	for _, t := range ts {
		key := t - t%10
		temps[key] = append(temps[key], t)
	}
	return temps
}
