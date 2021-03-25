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

const (
	UNICODE_UTF8  = iota // 使用UTF8编码
	UNICODE_ASCII        //使用ASCII编码
)

// 算法
const (
	ALGORITHM_LEVENSHTEIN = iota
	ALGORITHM_COSINE
	ALGORITHM_DICES_COEFFICIENT
	ALGORITHM_HAMMING
	ALGORITHM_JARO
	ALGORITHM_SIMHASH
)

//
// @Description: 配置项
//
type Config struct {
	Ignore    int // 忽略。空格,大小写，html标签
	Unicode   int // 使用ASCII还是UTF8
	Algorithm int // 所使用算法
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

func Default() *Config {
	return &Config{
		Ignore:    IGNORE_NONE,
		Unicode:   UNICODE_UTF8,
		Algorithm: ALGORITHM_LEVENSHTEIN,
	}
}
