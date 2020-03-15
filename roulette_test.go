package roby

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInt struct {
	score int
}

func (t TestInt) Score() int {
	return t.score
}

func TestRoulette_Take(t *testing.T) {
	rand.Seed(3)
	r := NewRoulette([]Scorer{TestInt{score: 2}, TestInt{score: 1}, TestInt{score: 7}})
	assert.Equal(t, 7, r.Take().(TestInt).score, "0")
	assert.Equal(t, 7, r.Take().(TestInt).score, "1")
	assert.Equal(t, 7, r.Take().(TestInt).score, "2")
	assert.Equal(t, 1, r.Take().(TestInt).score, "3")
	assert.Equal(t, 7, r.Take().(TestInt).score, "4")
	assert.Equal(t, 2, r.Take().(TestInt).score, "5")
	assert.Equal(t, 7, r.Take().(TestInt).score, "6")
	assert.Equal(t, 7, r.Take().(TestInt).score, "7")
	assert.Equal(t, 7, r.Take().(TestInt).score, "8")
	assert.Equal(t, 7, r.Take().(TestInt).score, "9")
}

func ExampleNewRoulette() {
	arr := []int{2, 2, 6}
	add := 0
	for i := range arr {
		arr[i], add = arr[i]+add, arr[i]
	}
	fmt.Println(arr)

	// Output:
	// [2 4 8]
}
