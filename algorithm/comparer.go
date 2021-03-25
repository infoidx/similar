/**
 * @Author: xingshanghe
 * @Description:
 * @File:  comparer
 * @Version: 0.1.0
 * @Date: 2021/3/24 5:35 下午
 */
package algorithm

import (
	"fmt"

	"github.com/infoidx/similar/configure"
)

//
// @Description: 比较器接口
//
type Comparer interface {
	Compare(source, target string) float64
}

func New(opts ...configure.Option) Comparer {
	cfg := configure.Default().Configure(opts...)
	// TODO 根据cfg实现，初始化正确的方法

	fmt.Println(cfg)
	return nil
}
