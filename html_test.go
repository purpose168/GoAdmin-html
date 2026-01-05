// Package html 提供了一个HTML元素构造器，用于通过Go代码动态生成HTML标签
package html

import (
	"fmt"
	"testing"
)

// TestBase_Get 测试基础HTML元素的构造功能
// 该测试演示了如何使用链式方法调用创建复杂的HTML结构
// 测试内容包括：
// 1. 创建一个div容器，包含居中对齐样式和dropdown类
// 2. 在div中创建一个超链接(a标签)，包含下拉切换样式和属性
// 3. 在div中创建一个无序列表(ul标签)，包含下拉菜单样式和属性
// 4. 展示了如何组合多个HTML元素并设置样式和属性
func TestBase_Get(t *testing.T) {
	fmt.Println(Div(
		A("asfasdfa",
			M{"color": "#676565"},
			M{"class": "dropdown-toggle", "href": "#", "data-toggle": "dropdown"},
		)+Ul("23234234",
			M{"min-width": "20px !important", "left": "-32px", "overflow": "hidden"},
			M{"class": "dropdown-menu", "role": "menu", "aria-labelledby": "dLabel"}),
		M{"text-align": "center"}, M{"class": "dropdown"}))
}
