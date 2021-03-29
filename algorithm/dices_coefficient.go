/**
 * @Author: xingshanghe
 * @Description: 骰子系数？
 * @File:  dices_coefficient
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:22 下午
 */
package algorithm

import (
	"strings"
	"unicode/utf8"
)

type DicesCoefficient struct {
	Ngram int //是筛子系数需要用的一个值
}

type count struct {
	s1 int
	s2 int
}

//
// @Description:
// @Date: 2021-03-29 17:06:31
// @receiver d
// @param set
// @param s
// @param add
// @return mixed
// @return length
//
func (d *DicesCoefficient) setOrGet(set map[string]count, s string, add bool) (mixed, length int) {
	var key strings.Builder
	ngram := d.Ngram
	if ngram == 0 {
		ngram = 2
	}

	for i := 0; i < len(s); {
		firstSize := 0
		for j, total := 0, 0; j < ngram && i+total < len(s); j++ {
			r, size := utf8.DecodeRuneInString(s[i+total:])
			key.WriteRune(r)
			total += size
			if j == 0 {
				firstSize = size
			}
		}
		if utf8.RuneCountInString(key.String()) != ngram {
			break
		}
		c, ok := set[key.String()]
		if add {
			if !ok {
				c = count{}
			}

			c.s1++
		} else {
			if !ok {
				goto next
			}

			c.s2++
			if c.s1 >= c.s2 {
				mixed++
			}
		}

		set[key.String()] = c
	next:
		key.Reset()
		length++
		i += firstSize
	}

	return mixed, length
}

func (d *DicesCoefficient) Compare(source, target string) float64 {
	set := make(map[string]count, len(source)/3)

	//TODO 边界比 如字符长度小于ngram ？这里原库实现还有些问题
	mixed, l1 := d.setOrGet(set, source, true)
	mixed, l2 := d.setOrGet(set, target, false)

	return 2.0 * float64(mixed) / float64(l1+l2)
}
