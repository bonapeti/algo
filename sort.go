package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"time"
)

func trackTime(start time.Time, name string, size int) {
	elapsed := time.Since(start)
	log.Printf("%s of %d took %s", name, size, elapsed)
}

func mergeSort(a []int) {
	defer trackTime(time.Now(), "mergeSort", len(a))
	//log.Printf("Array: %v", a)
	if len(a) < 2 {
		return
	}
	doMergeSort(a, 1, len(a))
}

func doMergeSort(a []int, p, r int) {
	if p < r {
		//log.Printf("doMergeSort: %v from %d to %d", a, p, r)
		q := (p + r) / 2
		//log.Printf("Middle: %d", q)
		doMergeSort(a, p, q)
		doMergeSort(a, q+1, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p, q, r int) {
	//log.Printf("Merge: %v from %d through %d to %d", a, p, q, r)
	n1 := q - p + 1
	n2 := r - q
	leftArray := make([]int, n1+1)
	rightArray := make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		leftArray[i-1] = a[p+i-2]
	}
	leftArray[n1] = math.MaxInt64
	//log.Printf("Left: %v", leftArray)
	for j := 1; j <= n2; j++ {
		rightArray[j-1] = a[q+j-1]
	}
	rightArray[n2] = math.MaxInt64
	//log.Printf("Right: %v", rightArray)
	i := 1
	j := 1
	for k := p; k <= r; k++ {
		if leftArray[i-1] <= rightArray[j-1] {
			a[k-1] = leftArray[i-1]
			i++
		} else {
			a[k-1] = rightArray[j-1]
			j++
		}
	}
	//log.Printf("Array: %v", a)
}

func insertionSort(a []int) {
	defer trackTime(time.Now(), "insertionSort", len(a))
	for i := 1; i < len(a); i++ {
		key := a[i]
		//		log.Printf("i: %d, key: %d", i, key)
		for j := i - 1; j >= 0; j-- {
			//			log.Printf("j: %d, a[j]: %d", j, a[j])
			if a[j] > key {
				a[j+1] = a[j]
				a[j] = key
				//				log.Printf("Swapped: %s", a)
			}

		}
	}
}

func randomSamples(size int) []int {
	rand.Seed(42)
	a := []int{}
	for i := 0; i < size; i++ {
		a = append(a, rand.Intn(size))
	}
	return a
}

func worstCase(size int) []int {
	a := []int{}
	for i := 0; i < size; i++ {
		a = append(a, size-i)
	}
	return a
}

func bestCase(size int) []int {
	a := []int{}
	for i := 0; i < size; i++ {
		a = append(a, i)
	}
	return a
}

func main() {
	rand.Seed(42)
	var sizePar = flag.Int("m", 1, "Magnitude")
	flag.Parse()
	size := int(math.Pow10(*sizePar))
	a := worstCase(size)
	//log.Printf("Before: %v", a)
	mergeSort(a)
	//insertionSort(a)
	//insertionSort(bestCase(size))

	//log.Printf("After: %v", a)

}
