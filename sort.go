package main

import (
	"log"
	"time"
	"flag"
	"math/rand"
	"math"
)

func trackTime(start time.Time, name string, size int) {
	elapsed := time.Since(start)
	log.Printf("%s of %d took %s", name, size, elapsed)
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

func randomSamples(size int)  []int {
	rand.Seed(42)
	a := []int{}
	for i := 0; i < size; i++ {
	    a = append(a, rand.Intn(size))
	}
	return a
}

func worstCase(size int)  []int {
	a := []int{}
	for i := 0; i < size; i++ {
	    a = append(a, size - i)
	}
	return a
}

func bestCase(size int)  []int {
	a := []int{}
	for i := 0; i < size; i++ {
	    a = append(a, i)
	}
	return a
}

func main() {
     	rand.Seed(42)
	var sizePar = flag.Int("m", 100, "Magnitude")
	flag.Parse()
	size := int(math.Pow10(*sizePar))
	//log.Printf("Before: %v", a)
	//insertionSort(randomSamples(size))
	insertionSort(worstCase(size))
	//log.Printf("After: %v", a)

}
