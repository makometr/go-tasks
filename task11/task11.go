package task11

// 11.	Написать пересечение двух неупорядоченных массивов.

// Task11 provide intersection of int arrays by hashing
func Task11(a, b []int) []int {
	unique := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		unique[a[i]] = false
	}
	for i := 0; i < len(b); i++ {
		if _, ok := unique[b[i]]; ok {
			unique[b[i]] = true
		}
	}

	if len(unique) == 0 {
		return nil
	}

	result := make([]int, 0)
	for k, v := range unique {
		if v == true {
			result = append(result, k)
		}
	}

	return result
}

// Task11Interface provide interfaces realiztation
func Task11Interface(a, b []interface{}) []interface{} {
	unique := make(map[interface{}]bool)
	for i := 0; i < len(a); i++ {
		unique[a[i]] = false
	}
	for i := 0; i < len(b); i++ {
		if _, ok := unique[b[i]]; ok {
			unique[b[i]] = true
		}
	}

	if len(unique) == 0 {
		return nil
	}

	result := make([]interface{}, 0)
	for k, v := range unique {
		if v == true {
			result = append(result, k)
		}
	}

	return nil
}
