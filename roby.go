package roby

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/xuender/oil/integer"
)

// NUM 平均验证次数
const NUM = 10

// Roby 机器人
type Roby struct {
	DNA    [243]int
	scores []int
}

func (r *Roby) String() string {
	var buffer bytes.Buffer
	for _, dna := range r.DNA {
		buffer.WriteString(strconv.Itoa(dna))
	}
	return fmt.Sprintf("{DNA:%s, AGE:%d, SCORE:%d}", buffer.String()[:20], len(r.scores)/NUM, r.Score())
}

// Score 平均分
func (r *Roby) Score() int {
	sum := 0
	max := integer.MinInt
	for _, s := range r.scores[len(r.scores)-NUM:] {
		sum += s
		if max < s {
			max = s
		}
	}
	return max
	// return sum*100/NUM + sum%NUM
	// if sum%NUM == 0 {
	// 	return ret
	// }
	// return ret + 1
}

// Movement 动作
func (r *Roby) Movement(around [5]int) int {
	ret := 0
	for i, a := range around {
		if a > 0 {
			ret += integer.Exp(3, 4-i) * a
		}
	}
	return r.DNA[ret]
}

// move 移动检查
func (r *Roby) move(movement int, state [5]int, x, y int) (bool, int, int) {
	switch movement {
	case 0:
		// 上
		if state[0] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x, y - 1
	case 1:
		// 下
		if state[1] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x, y + 1
	case 2:
		// 左
		if state[2] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x - 1, y
	case 3:
		// 右
		if state[3] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x + 1, y
	}
	return true, x, y
}

// Work 捡罐子
func (r *Roby) Work() {
	// r.scores = append([]int{})
	for i := 0; i < NUM; i++ {
		r.Run(NewStage())
	}
}

// Run 运行
func (r *Roby) Run(stage Stage) {
	// 初始位置
	x := rand.Intn(10)
	y := rand.Intn(10)
	score := 0

	// 走50步
	for i := 0; i < 200; i++ {
		state := stage.State(x, y)
		// 获取行动
		movement := r.Movement(state)
		// 行动
		switch movement {
		case 0, 1, 2, 3:
			// 上下左右
			if ok, nx, ny := r.move(movement, state, x, y); ok {
				x = nx
				y = ny
			} else {
				score -= 5 // 撞墙扣5分
			}
		case 4:
			// 随机移动
			if ok, nx, ny := r.move(rand.Intn(4), state, x, y); ok {
				x = nx
				y = ny
			} else {
				score -= 5 // 撞墙扣5分
			}
		case 5:
			// 捡罐头
			if stage[y][x] == 1 {
				stage[y][x] = 0
				score += 10 // 捡到罐子10分
			} else {
				score-- // 没有罐子扣1分
			}
		case 6:
			// 不动
		}
	}
	r.scores = append(r.scores, score)
}

// Variation 变异
func (r *Roby) Variation() {
	// 变异5次
	for i := 0; i < 20; i++ {
		r.DNA[rand.Intn(243)] = rand.Intn(7)
	}
}

// NewRoby 新罗比
func NewRoby(robys ...*Roby) *Roby {
	// 初始化罗比
	r := Roby{}
	if len(robys) == 0 {
		for i := 0; i < 243; i++ {
			r.DNA[i] = rand.Intn(7)
		}
	} else {
		// 遗传
		if rand.Intn(2) == 0 {
			reverse(robys)
		}
		for i := range r.DNA {
			r.DNA[i] = robys[i%len(robys)].DNA[i]
		}
		// 变异
		r.Variation()
	}
	return &r
}
func reverse(s []*Roby) []*Roby {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
