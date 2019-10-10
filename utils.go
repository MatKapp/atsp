package main

func shuffle(array []int) []int {
	var r RNG
	SIZE := len(array)

	for i := 0; i < SIZE; i++ {
		array[i] = i
	}

	for i := uint32(SIZE - 1); i >= 1; i-- {

		index := r.Uint32n(i)
		tmp := array[index]
		array[index] = array[i]
		array[i] = tmp
	}

	return array
}
