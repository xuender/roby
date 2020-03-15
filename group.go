package roby

import (
	"fmt"
	"sort"

	"github.com/xuender/oil/random"
)

// Group 族群
type Group struct {
	robys []*Roby
	size  int
}

// Sort 族群排序
func (g *Group) Sort() {
	sort.Slice(g.robys, func(i, j int) bool {
		return g.robys[i].Score() > g.robys[j].Score()
	})
}

// Run 运行
func (g *Group) Run(size int) {
	for i := 0; i < size; i++ {
		// 罗比工作
		for _, r := range g.robys {
			r.Work()
		}
		// 排序
		g.Sort()
		fmt.Println(i, "最优罗比:", g.robys[0])
		// 轮盘优选
		scores := make([]random.Scorer, len(g.robys))
		for i, r := range g.robys {
			scores[i] = r
		}
		r := random.NewRoulette(scores)
		// 保留得分最高的1/4
		for i := g.size / 10; i < g.size; i++ {
			// 繁殖变异
			g.robys[i] = NewRoby(r.Take().(*Roby), r.Take().(*Roby))
		}
	}
}

// NewGroup 新建族群
func NewGroup(size int) *Group {
	group := Group{
		size:  size,
		robys: make([]*Roby, size),
	}
	for i := 0; i < size; i++ {
		group.robys[i] = NewRoby()
	}
	return &group
}
