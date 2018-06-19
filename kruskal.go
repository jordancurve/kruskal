// Monte-Carlo simulation for the "Kruskal Count" card trick.
// The simulation gives the probability of the magician correctly identifying the correct card.
// https://arxiv.org/pdf/math/0110143.pdf
// https://twitter.com/ProbFact/status/1009111128213946368
// Author: Josh Jordan (@jordancurve)
package main

import (
        "fmt"
        "math/rand"
        "runtime"
        "sync"
        "time"
)

var cardValue = []int{5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 5}

func main() {
        nTrial := int(1e7)
        nSuccess := parallelSum(nTrial, runtime.NumCPU(), countSuccesses)
        fmt.Printf("P(win) = %d/%d (%g)\n", nSuccess, nTrial, float64(nSuccess)/float64(nTrial))
}

// countSuccesses runs n trials and return the number of successes.
func countSuccesses(nTrial int) int {
        deck := []int{}
        for i := 0; i < 52; i++ {
                deck = append(deck, i)
        }
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        nSuccess := 0
        for i := 0; i < nTrial; i++ {
                shuffle(r, deck)
                secretNumber := r.Intn(10) + 1
                subjectCard := guess(secretNumber, deck)
                magicianCard := guess(1, deck)
                if subjectCard == magicianCard {
                        nSuccess++
                }
        }
        return nSuccess
}

func shuffle(r *rand.Rand, deck []int) {
        for i := 0; i < len(deck)-1; i++ {
                j := i + r.Intn(52-i-1)
                deck[i], deck[j] = deck[j], deck[i]
        }
}

func guess(n int, deck []int) int {
        keyCard := -1
        for _, card := range deck {
                n--
                if n <= 0 {
                        keyCard = card
                        n = cardValue[card%13]
                }
        }
        return keyCard
}

// parallelSum returns f(a) + f(b) + ... + f(z) where a+b+...+z = total.
// f should be a function that takes an integer argument specifying the number of iterations to run
// in that particular call to f().
// The sum is calculated by evaluating each call to f() in one of nParallel separate goroutines.
func parallelSum(total int, nParallel int, f func(n int) int) int {
        var sum int
        var m sync.Mutex
        var wg sync.WaitGroup
        for i := 0; i < nParallel; i++ {
                var n int
                if i == 0 {
                        n = total % nParallel
                }
                wg.Add(1)
                go func(n int) {
                        defer wg.Done()
                        x := f(n)
                        m.Lock()
                        sum += x
                        m.Unlock()
                }(total/nParallel + n)
        }
        wg.Wait()
        return sum
}
