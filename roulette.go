package roby

import (
	"math/rand"
	"sort"
)

// IntMax 最大整数
const IntMax = int(^uint(0) >> 1)

// IntMin 最小整数
const IntMin = ^IntMax

// Roulette 轮盘
type Roulette struct {
	Scores []Scorer
	list   [][2]int
	max    int
}

// Add 增加
func (r *Roulette) Add(score Scorer) {
	r.Scores = append(r.Scores, score)
}

func (r *Roulette) init() {
	if len(r.Scores) == 0 {
		return
	}
	if r.list == nil || len(r.list) != len(r.Scores) {
		// 轮盘初始化
		r.list = make([][2]int, len(r.Scores))
		// 初始化坐标，并寻找最小积分
		min := IntMax
		for i, s := range r.Scores {
			r.list[i][1] = i
			a := s.Score()
			r.list[i][0] = a
			if min > a {
				min = a
			}
		}
		// 将所有积分设置成大于1的正数
		if min < 1 {
			abs := -1*min + 1
			for i := range r.list {
				r.list[i][0] += abs
			}
		}
		// 按照积分排序
		sort.Slice(r.list, func(i, j int) bool {
			return r.list[i][0] < r.list[j][0]
		})
		// 后面的值进行偏移,扩大面积
		add := 0
		for i := range r.list {
			r.list[i][0], add = r.list[i][0]+add, r.list[i][0]
		}
		// 最大值比最大积分大1
		r.max = r.list[len(r.list)-1][0] + 1
	}
}

// Take 获取
func (r *Roulette) Take() Scorer {
	r.init()
	s := rand.Intn(r.max)
	for _, l := range r.list {
		if s < l[0] {
			return r.Scores[l[1]]
		}
	}
	return r.Scores[0]
}

// Scorer 积分
type Scorer interface {
	Score() int
}

// NewRoulette 新建轮盘赌
func NewRoulette(scores []Scorer) *Roulette {
	ret := &Roulette{Scores: scores}
	ret.init()
	return ret
}
