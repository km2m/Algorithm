package Algo

func Sort0(x, y int) int {
	return x + y
}

func Partition(xs []int, l int, r int) int {
	// return (l + r) / 2
	target := xs[r]
	toIndex := 0
	for i := 0; i <= r-1; i++ {
		if xs[i] < target {
			xs[toIndex], xs[i] = xs[i], xs[toIndex]
			toIndex++
		}
	}
	xs[toIndex], xs[r] = xs[r], xs[toIndex]
	return toIndex
}

func QuickSort_0(xs []int, l int, r int) {
	if l < r {
		q := Partition(xs, l, r)
		QuickSort_0(xs, l, q-1)
		QuickSort_0(xs, q+1, r)
	}
}

func QuickSort(xs []int) []int {
	QuickSort_0(xs, 0, len(xs)-1)
	return xs
}

func InsertOrder(xs []int) []int {
	for i := 1; i < len(xs); i++ {
		//
		key := xs[i]
		j := i - 1
		for ; j >= 0 && xs[j] > key; j-- {
			xs[j+1] = xs[j]
		}
		xs[j+1] = key
	}
	return xs
}
