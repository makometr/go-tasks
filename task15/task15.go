package task15

func Swap1(a, b *int) {
	*a, *b = *b, *a
}

func Swap2(a, b *int) {
	*b = *b ^ *a
	*a = *b ^ *a
	*b = *b ^ *a
}

// Можно писать и с доп. переменной
