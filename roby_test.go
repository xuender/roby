package roby

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoby_Average(t *testing.T) {
	r := NewRoby()
	r.scores = append(r.scores, 3, 3, 2)
	assert.Equal(t, 2, r.Score(), "平均分")
}

func TestRoby_Movement(t *testing.T) {
	r := NewRoby()
	m := map[int]bool{}
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 3; c++ {
				for d := 0; d < 3; d++ {
					for e := 0; e < 3; e++ {
						m[r.Movement([5]int{a, b, c, d, e})] = true
					}
				}
			}
		}
	}
	assert.Equal(t, 243, len(m), "DNA链")
}

func TestRoby_Work(t *testing.T) {
	r := NewRoby()
	r.Work()
	assert.Equal(t, 1, len(r.scores), "积分")
	r.Work()
	assert.Equal(t, 2, len(r.scores), "积分")
}

// func TestRoby_Run(t *testing.T) {
func ExampleRoby() {
	rand.Seed(3)
	data := "455351051456452243351633534451154250254154221451354510453256220034336004402541356453453250053056215343005531151152203453243006450145534054355644033052004366341442126140213356430424341446226153426123453230616665611301614145544550552231432632141"
	r := NewRoby()
	for i, d := range data {
		r.DNA[i] = int(d - 48)
	}
	r.Work()
	fmt.Println(r.Score())
	// Output:
	// 350
}
