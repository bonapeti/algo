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

func main() {
     	rand.Seed(42)
	var sizePar = flag.Int("m", 100, "Magnitude")
	flag.Parse()
	size := int(math.Pow10(*sizePar))
	a := []int{}
	for i := 0; i < size; i++ {
	    a = append(a, rand.Intn(size))
	}
	//log.Printf("Before: %v", a)
	insertionSort(a)
	//log.Printf("After: %v", a)

}
