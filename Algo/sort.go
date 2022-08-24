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

func Merge(A []int, l int, m int, r int) {
	ll := make([]int, m-l+1)
	rr := make([]int, r-(m+1)+1)
	copy(ll, A[l:m+1])
	copy(rr, A[m+1:r+1])
	i, j := 0, 0
	for k := l; k <= r; k++ {
		if j > len(rr)-1 {
			A[k] = ll[i]
			i++
		} else if i > len(ll)-1 {
			A[k] = rr[j]
			j++
		} else if ll[i] <= rr[j] {
			A[k] = ll[i]
			i++
		} else {
			A[k] = rr[j]
			j++
		}
	}
}

func MergeSort_0(A []int, l int, r int) {
	if l < r {
		m := (l + r) / 2
		MergeSort_0(A, l, m)
		MergeSort_0(A, m+1, r)
		Merge(A, l, m, r)
	}
}

func MergeSort(A []int) (B []int) {
	MergeSort_0(A, 0, len(A)-1)
	B = A
	return
}
