/**
 * @Author: xingshanghe
 * @Description: 正弦相似度
 * @File:  simhash
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:25 下午
 */
package algorithm

import "github.com/infoidx/similar/configure"

type Sines struct {
	*configure.Config
}

func (s *Sines) Compare(source, target string) float64 {
	return 0
}
