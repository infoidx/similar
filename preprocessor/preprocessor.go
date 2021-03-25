/**
 * @Author: xingshanghe
 * @Description:
 * @File:  preprocessor
 * @Version: 0.1.0
 * @Date: 2021/3/25 5:42 下午
 */
package preprocessor

// 和configure中配置对应
// IGNORE_SPACE
// IGNORE_CASE
// IGNORE_HTML
type PreProcessor interface {
	SpaceProcess(input string) (output string)
	SpaceCase(input string) (output string)
	HtmlProcess(input string) (output string)
	TabooProcess(input string) (output string)
}
