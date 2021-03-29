/**
 * @Author: xingshanghe
 * @Description: 亚罗距离
 * @File:  jaro
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:22 下午
 */
package algorithm

import (
	"math"
	"sort"
	"sync"
	"unicode/utf8"
)

type Jaro struct {
	MatchWindow int
}

type check struct {
	index int
	c     rune
}

var checkPool = sync.Pool{
	New: func() interface{} {
		return &check{}
	},
}

//
// @Description:
// @Date: 2021-03-29 17:38:06
// @receiver j
// @param source
// @param target
// @return float64
//
func (j *Jaro) Compare(source, target string) float64 {
	mw := max(utf8.RuneCountInString(source), utf8.RuneCountInString(target))/2 - 1
	if j.MatchWindow != 0 {
		mw = j.MatchWindow
	}

	m := 0
	matchSet := make(map[rune][]int, len(source)/3)
	lenSource := 0
	for _, c := range source {
		matchSet[c] = append(matchSet[c], lenSource)
		lenSource++
	}

	t := 0
	lenTarget := 0
	indexAndRuneSource := make([]*check, 0, 8)
	indexAndRuneTarget := make([]rune, 0, 8)

	defer func() {
		for _, v := range indexAndRuneSource {
			checkPool.Put(v)
		}
	}()

	for _, c := range target {
		indexes, ok := matchSet[c]
		lenTarget++
		if !ok {
			continue
		}

		for k, i := range indexes {
			if i == -1 {
				continue
			}

			if math.Abs(float64(lenTarget-1-i)) <= float64(mw) {
				m++

				currCheck := checkPool.Get().(*check)
				currCheck.index = i
				currCheck.c = c

				indexAndRuneSource = append(indexAndRuneSource, currCheck)
				indexAndRuneTarget = append(indexAndRuneTarget, c)

				indexes[k] = -1
				break
			}
		}
	}

	m2 := float64(m)
	if m2 == 0 {
		return 0.0
	}

	sort.Slice(indexAndRuneSource, func(i, j int) bool {
		return indexAndRuneSource[i].index < indexAndRuneSource[j].index
	})

	for i, v := range indexAndRuneSource {
		if v.c != indexAndRuneTarget[i] {
			t++
		}
	}

	return 1.0 / 3.0 * (m2/float64(lenSource) + m2/float64(lenTarget) + (m2-float64(t)/2.0)/m2)
}
