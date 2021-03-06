/**
 * @Author: xingshanghe
 * @Description: 莱文斯顿 编辑距离
 * @File:  levenshtein
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:21 下午
 */
package algorithm

type Levenshtein struct {
}

//
// @Description: 实现比较接口
// @Date: 2021-03-29 14:21:31
// @receiver *Levenshtein
// @param source
// @param target
// @return float64
//
func (*Levenshtein) Compare(source, target string) float64 {
	s1, s2 := []rune(source), []rune(target)
	cacheX := make([]int, len(s2))

	diagonal := 0
	for y, yLen := 0, len(s1); y < yLen; y++ {
		for x, xLen := 0, len(cacheX); x < xLen; x++ {
			on := x + 1
			left := y + 1
			if x == 0 {
				diagonal = y
			} else if y == 0 {
				diagonal = x
			}
			if y > 0 {
				on = cacheX[x]
			}
			if x-1 >= 0 {
				left = cacheX[x-1]
			}

			same := 0
			if s1[y] != s2[x] {
				same = 1
			}

			oldDiagonal := cacheX[x]
			cacheX[x] = min(min(on+1, left+1), same+diagonal)
			diagonal = oldDiagonal

		}
	}

	return 1.0 - float64(cacheX[len(cacheX)-1])/float64(max(len(s1), len(s2)))
}

//
// @Description: 取两个数最小值
// @Date: 2021-03-29 14:28:13
// @param x
// @param y
// @return int
//
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//
// @Description: 取两个数最大值
// @Date: 2021-03-29 14:28:16
// @param x
// @param y
// @return int
//
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
