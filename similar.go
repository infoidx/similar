/**
 * @Author: xingshanghe
 * @Description:
 * @File:  similar
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:09 下午
 */
package similar

import (
	"github.com/infoidx/similar/algorithm"
	"github.com/infoidx/similar/configure"
)

func Compare(source, target string, opts ...configure.Option) float64 {
	return algorithm.New(opts...).Compare(source, target)
}

//
// @Description: 使用莱文斯顿距离算法比较
// @Date: 2021-03-25 13:39:34
// @return configure.Option
//
func UseLevenshtein() configure.Option {
	return func(config *configure.Config) {

	}
}

//
// @Description: 使用正弦相似度算法比较
// @Date: 2021-03-25 13:39:24
// @return configure.Option
//
func UseSimHash() configure.Option {
	return func(config *configure.Config) {
		// TODO
	}
}

//
// @Description: 适应余弦相似度算法比较
// @Date: 2021-03-25 13:39:08
// @return configure.Option
//
func UseCosine() configure.Option {
	return func(config *configure.Config) {
		// TODO
	}
}

//
// @Description: 使用骰子🎲算法比较
// @Date: 2021-03-25 13:38:40
// @return configure.Option
//
func UseDicesCoefficient() configure.Option {
	return func(config *configure.Config) {
		// TODO
	}
}

//
// @Description: 使用汉明距离算法比较
// @Date: 2021-03-25 13:38:30
// @return configure.Option
//
func UseHamming() configure.Option {
	return func(config *configure.Config) {
		// TODO
	}
}

//
// @Description: 使用亚罗算法比较
// @Date: 2021-03-25 13:38:13
// @return configure.Option
//
func UseJaro() configure.Option {
	return func(config *configure.Config) {
		// TODO
	}
}
