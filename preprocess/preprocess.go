/**
 * @Author: xingshanghe
 * @Description:
 * @File:  preprocessor
 * @Version: 0.1.0
 * @Date: 2021/3/25 5:42 下午
 */
package preprocess

import (
	"regexp"
	"strings"

	"github.com/infoidx/similar/configure"
)

var Taboos []string                                 // 默认敏感词
var Spacers = []string{"\r", "\n", "\t", "\f", " "} // 默认空格

//
// @Description: 预处理器映射
// @Date: 2021-03-28 01:34:10
// @return map[int]PreProcessor
//
func preProcessorMap() map[int]PreProcessor {
	maps := make(map[int]PreProcessor)
	maps[configure.IGNORE_SPACE] = NewSpacePreProcessor()
	maps[configure.IGNORE_CASE] = NewCasePreProcessor()
	maps[configure.IGNORE_HTML] = NewHtmlPreProcessor()
	maps[configure.REPLACE_TABOO] = NewTabooPreProcessor()
	return maps
}

//
// @Description: 根据配置生成预处理器数组/切片
// @Date: 2021-03-28 23:13:55
// @param cfg
// @return []PreProcessor
//
func Array(cfg *configure.Config) []PreProcessor {
	var result []PreProcessor
	if cfg.Ignore > 0 {
		maps := preProcessorMap()
		for k, v := range maps {
			if cfg.Ignore&k > 0 {
				result = append(result, v)
			}
		}
	}
	return result
}

//
// @Description: 预处理器
//
type PreProcessor interface {
	PreProcess(input string) string
}

//
// @Description: 创建空格预处理器
// @Date: 2021-03-28 23:18:26
// @return PreProcessor
//
func NewSpacePreProcessor() PreProcessor {
	return &SpacePreProcessor{Spacers: Spacers}
}

//
// @Description: 创建大小写预处理器
// @Date: 2021-03-28 23:18:40
// @return PreProcessor
//
func NewCasePreProcessor() PreProcessor {
	return &CasePreProcessor{}
}

//
// @Description: 创建html标签预处理器
// @Date: 2021-03-28 23:18:51
// @return PreProcessor
//
func NewHtmlPreProcessor() PreProcessor {
	return &HtmlPreProcessor{}
}

//
// @Description: 创建敏感词处理器
// @Date: 2021-03-28 23:19:08
// @return PreProcessor
//
func NewTabooPreProcessor() PreProcessor {
	return &TabooPreProcessor{
		Taboos: Taboos,
	}
}

//
// @Description: 空格预处理器
//
type SpacePreProcessor struct {
	Spacers []string
}

//
// @Description: 实现预处理
// @Date: 2021-03-28 23:19:38
// @receiver p
// @param input
// @return string
//
func (p *SpacePreProcessor) PreProcess(input string) string {
	var oldnew []string
	for _, v := range p.Spacers {
		oldnew = append(oldnew, v, "")
	}
	replacer := strings.NewReplacer(oldnew...)
	return replacer.Replace(input)
}

//
// @Description: 大小写预处理器
//
type CasePreProcessor struct {
}

//
// @Description: 实现预处理
// @Date: 2021-03-28 23:20:01
// @receiver *CasePreProcessor
// @param input
// @return string
//
func (*CasePreProcessor) PreProcess(input string) string {
	return strings.ToLower(input)
}

//
// @Description: Html预处理器
//
type HtmlPreProcessor struct {
}

//
// @Description: 实现预处理
// @Date: 2021-03-28 23:20:22
// @receiver *HtmlPreProcessor
// @param input
// @return string
//
func (*HtmlPreProcessor) PreProcess(input string) string {
	// 将所有html标签转换成小写标签
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	input = reg.ReplaceAllStringFunc(input, strings.ToLower)
	// 替换掉注释和标签
	reg = regexp.MustCompile(`<!--[^>]+>|<iframe[\S\s]+?</iframe>|<a[^>]+>|</a>|<script[\S\s]+?</script>|<style[\S\s]+?</style>|<[\s]+?>|</[\s]+?>`)
	input = reg.ReplaceAllString(input, "")
	reg, _ = regexp.Compile("\\<[\\s]+?\\>")
	input = reg.ReplaceAllString(input, "")
	return strings.TrimSpace(input)
}

//
// @Description: 敏感词预处理器
//
type TabooPreProcessor struct {
	Taboos []string
}

//
// @Description: 实现预处理
// @Date: 2021-03-28 23:20:47
// @receiver t
// @param input
// @return string
//
func (t *TabooPreProcessor) PreProcess(input string) string {
	for _, taboo := range t.Taboos {
		input = strings.Replace(input, taboo, "**", -1)
	}
	return input
}
