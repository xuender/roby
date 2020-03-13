package roby

import (
	"fmt"
	"sort"
)

// Group 族群
type Group []*Roby

func (g Group) Len() int           { return len(g) }
func (g Group) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g Group) Less(i, j int) bool { return g[i].Average() > g[j].Average() }

// Run 运行
func (g Group) Run(size int) {
	for i := 0; i < size; i++ {
		// 罗比工作
		for _, r := range g {
			r.Work()
		}
		// 排序
		sort.Sort(g)
		// TODO 淘汰
		// TODO 繁殖
		// TODO 变异
		fmt.Println("轮:", i, "最优罗比积分:", g[0].Average(), g[0])
	}
}

// NewGroup 新建族群
func NewGroup(size int) Group {
	group := make(Group, size)
	for i := 0; i < size; i++ {
		group[i] = NewRoby()
	}
	return group
}
