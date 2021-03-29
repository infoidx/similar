/**
 * @Author: xingshanghe
 * @Description:
 * @File:  similar
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:09 下午
 */
package similar

import (
	"strings"

	"github.com/infoidx/similar/algorithm"
	"github.com/infoidx/similar/configure"
	"github.com/infoidx/similar/preprocess"
)

//
// @Description: 比较文本获取文本近似度
// @Date: 2021-03-28 23:21:18
// @param source
// @param target
// @param opts
// @return float64
//
func Compare(source, target string, opts ...configure.Option) float64 {
	if strings.TrimSpace(source) == strings.TrimSpace(target) {
		return 1.0
	}
	if len(source) == 0 || len(target) == 0 {
		return 0.0
	}

	cfg := configure.Default().Configure(opts...)

	preprocessors := preprocess.Array(cfg)
	for _, v := range preprocessors {
		source = v.PreProcess(source)
		target = v.PreProcess(target)
	}

	comparer := algorithm.New(cfg)
	return comparer.Compare(source, target)
}

//
// @Description: 使用莱文斯顿距离算法比较
// @Date: 2021-03-25 13:39:34
// @return configure.Option
//
func UseLevenshtein() configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_LEVENSHTEIN
	}
}

//
// @Description: 使用正弦相似度算法比较
// @Date: 2021-03-25 13:39:24
// @return configure.Option
//
func UseSines() configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_SINES
	}
}

//
// @Description: 适应余弦相似度算法比较
// @Date: 2021-03-25 13:39:08
// @return configure.Option
//
func UseCosines() configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_COSINES
	}
}

//
// @Description: 使用骰子🎲算法比较
// @Date: 2021-03-25 13:38:40
// @return configure.Option
//
func UseDicesCoefficient(ngrams ...int) configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_DICES_COEFFICIENT
		cfg.DicesCoefficient.Ngram = configure.DICES_COEFFICIENT_NGRAM

		if len(ngrams) > 0 {
			cfg.DicesCoefficient.Ngram = ngrams[0]
		}
	}
}

//
// @Description: 使用汉明距离算法比较
// @Date: 2021-03-25 13:38:30
// @return configure.Option
//
func UseHamming() configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_HAMMING
	}
}

//
// @Description: 使用亚罗算法比较
// @Date: 2021-03-25 13:38:13
// @return configure.Option
//
func UseJaro(matchWindows ...int) configure.Option {
	return func(cfg *configure.Config) {
		cfg.Algorithm = configure.ALGORITHM_JARO
		cfg.Jaro.MatchWindow = configure.JARO_MATCH_WINDOW

		if len(matchWindows) > 0 {
			cfg.Jaro.MatchWindow = matchWindows[0]
		}
	}
}

//
// @Description: 忽略大小写
// @Date: 2021-03-29 14:10:35
// @return configure.Option
//
func IgnoreCase() configure.Option {
	return func(config *configure.Config) {
		config.Ignore |= configure.IGNORE_CASE
	}
}

//
// @Description: 忽略空格
// @Date: 2021-03-29 14:10:48
// @return configure.Option
//
func IgnoreSpace() configure.Option {
	return func(config *configure.Config) {
		config.Ignore |= configure.IGNORE_SPACE
	}
}

//
// @Description: 忽略HTML
// @Date: 2021-03-29 14:10:58
// @return configure.Option
//
func IgnoreHtml() configure.Option {
	return func(config *configure.Config) {
		config.Ignore |= configure.IGNORE_HTML
	}
}

//
// @Description: 替换敏感词
// @Date: 2021-03-29 14:11:09
// @return configure.Option
//
func ReplaceTaboo() configure.Option {
	return func(config *configure.Config) {
		config.Ignore |= configure.REPLACE_TABOO
	}
}
