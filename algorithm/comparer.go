/**
 * @Author: xingshanghe
 * @Description:
 * @File:  comparer
 * @Version: 0.1.0
 * @Date: 2021/3/24 5:35 下午
 */
package algorithm

import (
	"github.com/infoidx/similar/configure"
)

//
// @Description: 比较器接口
//
type Comparer interface {
	Compare(source, target string) float64
}

func New(cfg *configure.Config) Comparer {
	var comparer Comparer
	switch cfg.Algorithm {
	case configure.ALGORITHM_COSINES:
		comparer = &Cosines{}
	case configure.ALGORITHM_DICES_COEFFICIENT:
		comparer = &DicesCoefficient{}
	case configure.ALGORITHM_HAMMING:
		comparer = &Hamming{}
	case configure.ALGORITHM_JARO:
		comparer = &Jaro{}
	case configure.ALGORITHM_SINES:
		comparer = &Sines{}
	case configure.ALGORITHM_LEVENSHTEIN:
		comparer = &Levenshtein{}
	default:
		comparer = &Levenshtein{}
	}
	return comparer
}
