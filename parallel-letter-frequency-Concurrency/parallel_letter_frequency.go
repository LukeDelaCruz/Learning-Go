package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

var lock sync.Mutex

func processData(wg *sync.WaitGroup, text *string, freqMap FreqMap) {
	defer wg.Done()

	newMap := Frequency(*text)

	// Mutex prevents fatal error: concurrent map writes
	lock.Lock()
	for k, v := range newMap {
		freqMap[k] += v
	}
	lock.Unlock()
}

// ConcurrentFrequency counts the frequency of each rune in the given strings, by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	res := FreqMap{}
	for i := range texts {
		wg.Add(1)

		// Confinement is done using &texts[i]
		go processData(&wg, &texts[i], res)
	}

	wg.Wait()

	return res
}
