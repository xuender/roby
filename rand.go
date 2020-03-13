package roby

import (
	"math/rand"
	"time"
)

// RandQueue 随机队列
type RandQueue []int

// NewRandQueue 新建随机队列
func NewRandQueue(max, size int) RandQueue {
	if size >= max {
		ret := make(RandQueue, max)
		for i := 0; i < max; i++ {
			ret[i] = i
		}
		return ret
	}
	ints := make(RandQueue, size)
	m := map[int]bool{}
	for {
		m[rand.Intn(max)] = true
		if len(m) >= size {
			break
		}
	}
	i := 0
	for num := range m {
		ints[i] = num
		i++
	}
	return ints
}

func init() {
	rand.Seed(time.Now().Unix())
}
