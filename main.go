package main

import (
	"fmt"
	"math/rand"

	"github.com/azolus/dstructOn/heap"
)

func main() {
	keys := []interface{}{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
	}

	heap := heap.NewMinHeap()
	heap.SetNodes(keys, score)

	// heap.Print()

	_, sScores := heap.Sort()

	fmt.Println(sScores, keys)
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func score(inf interface{}) int {
	return int([]rune(inf.(string))[0])
}
