/**
 * @Author: xingshanghe
 * @Description:
 * @File:  preprocessor
 * @Version: 0.1.0
 * @Date: 2021/3/25 5:42 下午
 */
package preprocess

// 和configure中配置对应
// IGNORE_SPACE
// IGNORE_CASE
// IGNORE_HTML
type PreProcessor interface {
	IgnoreSpace(input string) (output string)
	IgnoreCase(input string) (output string)
	IgnoreHtml(input string) (output string)
	ReplaceTaboo(input, replacer string) (output string)
}

type preprocess struct {
}

func (p *preprocess) IgnoreSpace(input string) (output string) {
	return ""
}

func (p *preprocess) IgnoreCase(input string) (output string) {
	return ""
}

func (p *preprocess) IgnoreHtml(input string) (output string) {
	return ""
}

func (p *preprocess) ReplaceTaboo(input, replacer string) (output string) {
	return ""
}
