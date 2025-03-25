package mymap

import (
	"fmt"
	"testing"
)

// 嵌套结构统计
// • 任务：接收 []struct{Name string; Score int}，返回按姓名统计的分数平均值（结果类型 map[string]float64）。

type Person struct {
	Name  string
	Score int
}
type myscores struct {
	n           int
	totalScores int
}

func ans5(scores []Person) (hashMap map[string]float64) {

	myscoresList := make(map[string]*myscores)

	for _, v := range scores {
		if _, ok := myscoresList[v.Name]; !ok {
			myscoresList[v.Name] = &myscores{n: 1, totalScores: v.Score}
		} else {
			temp := myscoresList[v.Name]
			temp.n++
			temp.totalScores += v.Score
		}
	}

	hashMap = make(map[string]float64)
	for k, v := range myscoresList {
		hashMap[k] = (float64(v.totalScores) / float64(v.n))
	}

	return hashMap
}

func Test5(t *testing.T) {
	fmt.Println(ans5([]Person{{"Alice", 80}, {"Bob", 70}, {"Alice", 90}}))
}
