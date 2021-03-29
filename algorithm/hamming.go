/**
 * @Author: xingshanghe
 * @Description: 汉明距离
 * @File:  hamming
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:22 下午
 */
package algorithm

import (
	"math"
	"unicode/utf8"
)

type Hamming struct {
}

//
// @Description: 实现比较接口
// @Date: 2021-03-29 14:21:47
// @receiver *Hamming
// @param source
// @param target
// @return float64
//
func (*Hamming) Compare(source, target string) float64 {
	count := 0

	leftSource := utf8.RuneCountInString(source)
	max := leftSource

	leftTarget := utf8.RuneCountInString(target)
	if max < leftTarget {
		max = leftTarget
	}

	for i, j := 0, 0; i < len(source) && j < len(target); {
		size := 0
		rightSource, size := utf8.DecodeRune([]byte(source[i:]))
		i += size

		rightTarget, size := utf8.DecodeRune([]byte(target[j:]))
		j += size

		if rightSource != rightTarget {
			count++
		}

	}

	return 1 - (float64(count)+math.Abs(float64(leftSource-leftTarget)))/float64(max)
}
