package roby

import (
	"bytes"
	"math/rand"
	"strconv"
)

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
	return buffer.String()
}

// Average 平均分
func (r *Roby) Average() int {
	sum := 0
	for _, s := range r.scores {
		sum += s
	}
	return sum / len(r.scores)
}

// Movement 动作
func (r *Roby) Movement(around [5]int) int {
	ret := 0
	for i, a := range around {
		w := a
		for f := 0; f < i; f++ {
			w *= 3
		}
		ret += w
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
	r.run(NewStage())
}
func (r *Roby) run(stage Stage) {
	// 初始位置
	x := rand.Intn(10)
	y := rand.Intn(10)
	score := 0

	// 走50步
	for i := 0; i < 50; i++ {
		state := stage.State(x, y)
		// 获取行动
		movement := r.Movement(state)
		// 行动
		switch movement {
		case 0:
		case 1:
		case 2:
		case 3:
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

// NewRoby 新罗比
func NewRoby() *Roby {
	// 初始化罗比
	r := Roby{}
	for i := 0; i < 243; i++ {
		r.DNA[i] = rand.Intn(7)
	}
	return &r
}
