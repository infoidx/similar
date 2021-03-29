/**
 * @Author: xingshanghe
 * @Description: 配置及配置项相关
 * @File:  configure
 * @Version: 0.1.0
 * @Date: 2021/3/24 4:19 下午
 */
package configure

const IGNORE_NONE = iota
const (
	IGNORE_SPACE  = 1 << iota // 忽略空格
	IGNORE_CASE               // 忽略大小写
	IGNORE_HTML               // 忽略HTML
	REPLACE_TABOO             // 替换敏感词
)

// 算法
const (
	ALGORITHM_LEVENSHTEIN = iota
	ALGORITHM_COSINES
	ALGORITHM_DICES_COEFFICIENT
	ALGORITHM_HAMMING
	ALGORITHM_JARO
	ALGORITHM_SINES
)

const (
	DICES_COEFFICIENT_NGRAM = 2
	JARO_MATCH_WINDOW       = 0
)

//
// @Description: 配置项
//
type Config struct {
	Ignore    int // 忽略。空格,大小写，html标签
	Algorithm int // 所使用算法
	DicesCoefficient
	Jaro
}

type DicesCoefficient struct {
	Ngram int
}

type Jaro struct {
	MatchWindow int
}

//
// @Description: 配置选项
// @Date: 2021-03-24 16:37:01
// @param *Config
//
type Option func(*Config)

//
// @Description: 配置
// @Date: 2021-03-24 16:36:51
// @receiver cfg
// @param opts
// @return *Config
//
func (cfg *Config) Configure(opts ...Option) *Config {
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

//
// @Description: 默认配置
// @Date: 2021-03-26 17:12:38
// @return *Config
//
func Default() *Config {
	return &Config{
		Ignore:    IGNORE_NONE,
		Algorithm: ALGORITHM_LEVENSHTEIN,
	}
}
