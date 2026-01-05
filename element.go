// Package html 提供了一个HTML元素构造器，用于通过Go代码动态生成HTML标签
//
// 该包提供了一组简洁的API，用于构建HTML元素，支持链式调用和灵活的样式/属性设置。
// 主要特性包括：
//   - 支持常见的HTML标签（div、span、button、form、a、ul、li等）
//   - 支持内联样式和HTML属性设置
//   - 支持元素的嵌套和组合
//   - 使用template.HTML类型确保HTML内容的安全输出
//
// 使用示例：
//
//	// 创建一个简单的div元素
//	html.Div("Hello World")
//
//	// 创建带有样式和属性的div元素
//	html.Div("内容",
//	    html.M{"color": "red", "font-size": "14px"},  // 样式
//	    html.M{"class": "container", "id": "main"})   // 属性
//
//	// 创建嵌套的HTML结构
//	html.Div(
//	    html.A("点击这里",
//	        html.M{"color": "blue"},
//	        html.M{"href": "https://example.com"}),
//	    html.M{"text-align": "center"},
//	    html.M{"class": "wrapper"})
//
// 注意事项：
//   - 所有返回template.HTML的函数都可以直接用于Go的html/template包
//   - 使用M类型简化样式和属性的定义
//   - 支持链式调用，但推荐使用便捷函数（如Div、Span等）来创建元素
package html

import (
	"fmt"
	"html/template"
	"strings"
)

// M 是一个字符串映射类型的别名，用于简化样式和属性的定义
//
// 使用示例：
//
//	style := html.M{"color": "red", "font-size": "14px"}
//	attr := html.M{"class": "btn", "id": "submit"}
type M map[string]string

// Style 是M的别名，用于表示HTML元素的样式
//
// Style类型提供了String()方法，可以将样式映射转换为HTML样式字符串。
// 例如: Style{"color": "red", "font-size": "12px"} 将转换为 ` style="color:red;font-size:12px;"`
type Style M

// Attribute 是M的别名，用于表示HTML元素的属性
//
// Attribute类型提供了String()方法，可以将属性映射转换为HTML属性字符串。
// 例如: Attribute{"class": "btn", "id": "submit"} 将转换为 ` class="btn" id="submit"`
type Attribute M

// String 将Style转换为HTML样式字符串
//
// 该方法遍历Style映射中的所有键值对，将它们转换为CSS样式格式。
// 每个样式属性之间用分号分隔，最后添加style属性前缀。
//
// 返回值示例：
//   - 空样式返回空字符串
//   - 非空样式返回 ` style="color:red;font-size:12px;"`
//
// 使用示例：
//
//	style := html.Style{"color": "red", "font-size": "12px"}
//	htmlStr := style.String()  // 返回 ` style="color:red;font-size:12px;"`
func (s Style) String() template.HTML {
	res := ""
	for k, v := range s {
		res += k + ":" + v + ";"
	}
	if res != "" {
		res = ` style="` + res + `"`
	}
	return template.HTML(res)
}

// String 将Attribute转换为HTML属性字符串
//
// 该方法遍历Attribute映射中的所有键值对，将它们转换为HTML属性格式。
// 每个属性之间用空格分隔，属性值用双引号包裹。
//
// 返回值示例：
//   - 空属性返回空字符串
//   - 非空属性返回 ` class="btn" id="submit"`
//
// 使用示例：
//
//	attr := html.Attribute{"class": "btn", "id": "submit"}
//	htmlStr := attr.String()  // 返回 ` class="btn" id="submit"`
func (s Attribute) String() template.HTML {
	res := ""
	for k, v := range s {
		res += k + `="` + v + `" `
	}
	if res != "" {
		res = ` ` + res[:len(res)-1]
	}
	return template.HTML(res)
}

// Element 表示一个HTML元素，包含标签名、内容、样式和属性
//
// Element结构体是HTML元素的核心表示，包含了构建HTML标签所需的所有信息。
// 通过链式方法调用可以灵活地设置元素的各个属性。
//
// 字段说明：
//   - Tag: HTML标签名，如"div"、"span"、"button"等
//   - Content: 元素的内容，可以是文本或其他HTML元素
//   - Style: 元素的内联样式，使用Style类型表示
//   - Attribute: 元素的HTML属性，使用Attribute类型表示
//
// 使用示例：
//
//	el := html.BaseEl()
//	el = el.SetTag("div")
//	el = el.SetContent("Hello")
//	el = el.SetStyle("color", "red")
//	el = el.SetAttr("class", "container")
//	htmlStr := el.Get()  // 返回 <div style="color:red;" class="container">Hello</div>
type Element struct {
	Tag       template.HTML // HTML标签名，如"div"、"span"等
	Content   template.HTML // 元素的内容，可以是文本或其他HTML元素
	Style     Style         // 元素的内联样式
	Attribute Attribute     // 元素的HTML属性
}

// BaseEl 创建并返回一个基础的Element实例
//
// 该函数初始化了一个空的Element结构体，其中Style和Attribute字段都被初始化为空的映射。
// 这是创建自定义HTML元素的起点，可以通过链式方法调用进一步设置元素的属性。
//
// 返回值：
//   - 返回一个初始化后的Element实例，Tag、Content为空，Style和Attribute为空映射
//
// 使用示例：
//
//	el := html.BaseEl()
//	el = el.SetTag("div").SetContent("内容").SetStyle("color", "red")
//	htmlStr := el.Get()
func BaseEl() Element {
	return Element{Style: make(map[string]string), Attribute: make(map[string]string)}
}

// SetTag 设置元素的HTML标签名
//
// 该方法设置Element的Tag字段，指定HTML元素的标签类型。
// 常见的标签名包括："div"、"span"、"p"、"button"、"form"、"a"、"ul"、"li"等。
//
// 参数：
//   - tag: HTML标签名，类型为template.HTML
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div")
func (b Element) SetTag(tag template.HTML) Element {
	b.Tag = tag
	return b
}

// SetContent 设置元素的内容
//
// 该方法设置Element的Content字段，指定HTML元素内部的内容。
// 内容可以是纯文本，也可以是其他HTML元素（通过template.HTML类型）。
//
// 参数：
//   - content: 元素的内容，类型为template.HTML
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetContent("Hello World")
func (b Element) SetContent(content template.HTML) Element {
	b.Content = content
	return b
}

// SetStyle 设置元素的单个样式属性
//
// 该方法设置Element的Style字段中的单个样式属性。
// 如果该样式属性已存在，将被覆盖；如果不存在，将被添加。
//
// 参数：
//   - key: 样式属性名，如"color"、"font-size"、"background-color"等
//   - value: 样式属性值，如"red"、"14px"、"#ffffff"等
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetStyle("color", "red").SetStyle("font-size", "14px")
func (b Element) SetStyle(key, value string) Element {
	b.Style[key] = value
	return b
}

// SetClass 设置元素的class属性
//
// 该方法设置Element的Attribute字段中的class属性，支持添加多个class名称。
// 如果元素已有class属性，新的class会追加到现有class后面，用空格分隔。
//
// 参数：
//   - class: 一个或多个class名称，使用可变参数
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetClass("container", "wrapper", "main")
//	// 结果: class="container wrapper main"
func (b Element) SetClass(class ...string) Element {
	if b.Attribute["class"] != "" {
		b.Attribute["class"] += " " + strings.Join(class, " ")
	} else {
		b.Attribute["class"] += strings.Join(class, " ")
	}
	return b
}

// SetId 设置元素的id属性
//
// 该方法设置Element的Attribute字段中的id属性。
// id属性在HTML文档中应该是唯一的，用于标识特定的元素。
//
// 参数：
//   - id: 元素的唯一标识符
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetId("main-container")
//	// 结果: id="main-container"
func (b Element) SetId(id string) Element {
	b.Attribute["id"] = id
	return b
}

// SetAttr 设置元素的单个属性
//
// 该方法设置Element的Attribute字段中的单个HTML属性。
// 可以设置任何HTML属性，如"href"、"src"、"alt"、"data-*"等。
//
// 参数：
//   - key: 属性名，如"href"、"src"、"data-toggle"等
//   - value: 属性值，如"https://example.com"、"image.jpg"、"dropdown"等
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("a").SetAttr("href", "https://example.com").SetAttr("target", "_blank")
func (b Element) SetAttr(key, value string) Element {
	b.Attribute[key] = value
	return b
}

// SetStyleAndAttr 同时设置元素的样式和属性
//
// 该方法接受一个M类型的切片，用于批量设置元素的样式和属性。
// 切片的第一个元素（如果存在）用于设置样式，第二个元素（如果存在）用于设置属性。
// 这是一种便捷的方法，可以一次性设置多个样式和属性。
//
// 参数：
//   - ms: M类型的切片，最多包含两个元素
//   - ms[0]: 样式映射，用于设置元素的样式
//   - ms[1]: 属性映射，用于设置元素的属性
//
// 返回值：
//   - 返回设置后的Element实例，支持链式调用
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetContent("内容")
//	el = el.SetStyleAndAttr([]html.M{
//	    {"color": "red", "font-size": "14px"},  // 样式
//	    {"class": "container", "id": "main"},   // 属性
//	})
func (b Element) SetStyleAndAttr(ms []M) Element {
	if len(ms) > 0 {
		for k, v := range ms[0] {
			b.Style[k] = v
		}
	}
	if len(ms) > 1 {
		for k, v := range ms[1] {
			b.Attribute[k] = v
		}
	}
	return b
}

// Get 将Element转换为完整的HTML标签字符串
//
// 该方法将Element的所有字段组合成一个完整的HTML标签字符串。
// 生成的HTML字符串遵循标准HTML格式，包含标签名、样式属性、其他属性和内容。
//
// 返回值：
//   - 返回完整的HTML标签字符串，类型为template.HTML
//
// 返回值示例：
//   - <div style="color:red;" class="container">内容</div>
//   - <a href="https://example.com" target="_blank">链接</a>
//
// 使用示例：
//
//	el := html.BaseEl().SetTag("div").SetContent("Hello").SetStyle("color", "red")
//	htmlStr := el.Get()  // 返回 <div style="color:red;">Hello</div>
func (b Element) Get() template.HTML {
	return template.HTML(fmt.Sprintf(`<%s%s%s>%s</%s>`, b.Tag, b.Style.String(), b.Attribute.String(), b.Content, b.Tag))
}

// BodyEl 创建并返回一个body标签的Element实例
//
// body标签是HTML文档的主体部分，包含了页面的所有可见内容。
//
// 返回值：
//   - 返回一个tag为"body"的Element实例
//
// 使用示例：
//
//	el := html.BodyEl().SetContent("页面内容").SetStyle("background-color", "#f0f0f0")
func BodyEl() Element {
	return BaseEl().SetTag("body")
}

// Body 创建一个完整的body标签
//
// 该函数创建一个包含指定内容和样式/属性的body标签。
// body标签通常用于包裹HTML文档的主体内容。
//
// 参数：
//   - content: body标签的内容，可以是文本或其他HTML元素
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的body标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Body(
//	    "页面主体内容",
//	    html.M{"background-color": "#f0f0f0"},
//	    html.M{"class": "page-body"},
//	)
//	// 结果: <body style="background-color:#f0f0f0;" class="page-body">页面主体内容</body>
func Body(content template.HTML, ms ...M) template.HTML {
	return BodyEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// DivEl 创建并返回一个div标签的Element实例
//
// div标签是HTML中最常用的块级元素，用于组织和布局页面内容。
// div元素默认占据一行，可以包含其他HTML元素。
//
// 返回值：
//   - 返回一个tag为"div"的Element实例
//
// 使用示例：
//
//	el := html.DivEl().SetContent("div内容").SetClass("container")
func DivEl() Element {
	return BaseEl().SetTag("div")
}

// Div 创建一个完整的div标签
//
// 该函数创建一个包含指定内容和样式/属性的div标签。
// div标签是HTML中最常用的容器元素，用于布局和组织内容。
//
// 参数：
//   - content: div标签的内容，可以是文本或其他HTML元素
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的div标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Div(
//	    "div内容",
//	    html.M{"color": "red", "font-size": "14px"},
//	    html.M{"class": "container", "id": "main"},
//	)
//	// 结果: <div style="color:red;font-size:14px;" class="container" id="main">div内容</div>
func Div(content template.HTML, ms ...M) template.HTML {
	return DivEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// SpanEl 创建并返回一个span标签的Element实例
//
// span标签是HTML中的行内元素，用于对文本的一部分进行样式化或分组。
// span元素不会独占一行，可以与其他行内元素并排显示。
//
// 返回值：
//   - 返回一个tag为"span"的Element实例
//
// 使用示例：
//
//	el := html.SpanEl().SetContent("span内容").SetStyle("color", "blue")
func SpanEl() Element {
	return BaseEl().SetTag("span")
}

// Span 创建一个完整的span标签
//
// 该函数创建一个包含指定内容和样式/属性的span标签。
// span标签是行内元素，常用于对文本的一部分进行样式化。
//
// 参数：
//   - content: span标签的内容，通常是文本
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的span标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Span(
//	    "重要文本",
//	    html.M{"color": "red", "font-weight": "bold"},
//	    html.M{"class": "highlight"},
//	)
//	// 结果: <span style="color:red;font-weight:bold;" class="highlight">重要文本</span>
func Span(content template.HTML, ms ...M) template.HTML {
	return SpanEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// IEl 创建并返回一个i标签的Element实例
//
// i标签通常用于表示斜体文本，但在现代Web开发中更常用于显示图标。
// 许多图标库（如Font Awesome、Bootstrap Icons）都使用i标签来显示图标。
//
// 返回值：
//   - 返回一个tag为"i"的Element实例
//
// 使用示例：
//
//	el := html.IEl().SetAttr("class", "fa fa-home")
func IEl() Element {
	return BaseEl().SetTag("i")
}

// I 创建一个完整的i标签
//
// 该函数创建一个包含指定样式/属性的i标签。
// i标签通常用于显示图标，通过设置class属性来指定具体的图标。
//
// 参数：
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的i标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.I(
//	    html.M{},
//	    html.M{"class": "fa fa-home", "aria-hidden": "true"},
//	)
//	// 结果: <i class="fa fa-home" aria-hidden="true"></i>
func I(ms ...M) template.HTML {
	return IEl().SetStyleAndAttr(ms).Get()
}

// PEl 创建并返回一个p标签的Element实例
//
// p标签用于表示段落，是HTML中最常用的文本容器之一。
// 浏览器会在段落前后自动添加一些空白。
//
// 返回值：
//   - 返回一个tag为"p"的Element实例
//
// 使用示例：
//
//	el := html.PEl().SetContent("这是一个段落").SetStyle("line-height", "1.6")
func PEl() Element {
	return BaseEl().SetTag("p")
}

// P 创建一个完整的p标签
//
// 该函数创建一个包含指定内容和样式/属性的p标签。
// p标签用于表示段落，是HTML文档中最常用的文本容器。
//
// 参数：
//   - content: p标签的内容，通常是段落文本
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的p标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.P(
//	    "这是一个段落文本",
//	    html.M{"line-height": "1.6", "color": "#333"},
//	    html.M{"class": "paragraph"},
//	)
//	// 结果: <p style="line-height:1.6;color:#333;" class="paragraph">这是一个段落文本</p>
func P(content template.HTML, ms ...M) template.HTML {
	return PEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// ButtonEl 创建并返回一个button标签的Element实例
//
// button标签用于创建可点击的按钮，常用于表单提交或触发JavaScript操作。
// 按钮可以包含文本或图标等内容。
//
// 返回值：
//   - 返回一个tag为"button"的Element实例
//
// 使用示例：
//
//	el := html.ButtonEl().SetContent("点击我").SetAttr("type", "button")
func ButtonEl() Element {
	return BaseEl().SetTag("button")
}

// Button 创建一个完整的button标签
//
// 该函数创建一个包含指定内容和样式/属性的button标签。
// button标签用于创建可点击的按钮，可以用于表单提交或触发操作。
//
// 参数：
//   - content: button标签的内容，通常是按钮文本
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的button标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Button(
//	    "提交",
//	    html.M{"padding": "10px 20px", "background-color": "#007bff"},
//	    html.M{"type": "submit", "class": "btn btn-primary"},
//	)
//	// 结果: <button style="padding:10px 20px;background-color:#007bff;" type="submit" class="btn btn-primary">提交</button>
func Button(content template.HTML, ms ...M) template.HTML {
	return ButtonEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// FormEl 创建并返回一个form标签的Element实例
//
// form标签用于创建HTML表单，用于收集用户输入。
// 表单可以包含各种输入元素，如文本框、复选框、单选按钮等。
//
// 返回值：
//   - 返回一个tag为"form"的Element实例
//
// 使用示例：
//
//	el := html.FormEl().SetAttr("action", "/submit").SetAttr("method", "post")
func FormEl() Element {
	return BaseEl().SetTag("form")
}

// Form 创建一个完整的form标签
//
// 该函数创建一个包含指定内容和样式/属性的form标签。
// form标签用于创建HTML表单，用于收集和提交用户数据。
//
// 参数：
//   - content: form标签的内容，通常包含各种表单元素
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的form标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Form(
//	    "表单内容",
//	    html.M{"margin": "20px"},
//	    html.M{"action": "/submit", "method": "post", "class": "login-form"},
//	)
//	// 结果: <form style="margin:20px;" action="/submit" method="post" class="login-form">表单内容</form>
func Form(content template.HTML, ms ...M) template.HTML {
	return FormEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// AEl 创建并返回一个a标签的Element实例
//
// a标签用于创建超链接，可以链接到其他页面、文件或同一页面的其他位置。
// href属性指定链接的目标地址。
//
// 返回值：
//   - 返回一个tag为"a"的Element实例
//
// 使用示例：
//
//	el := html.AEl().SetContent("点击这里").SetAttr("href", "https://example.com")
func AEl() Element {
	return BaseEl().SetTag("a")
}

// A 创建一个完整的a标签
//
// 该函数创建一个包含指定内容和样式/属性的a标签。
// a标签用于创建超链接，可以链接到其他页面或资源。
//
// 参数：
//   - content: a标签的内容，通常是链接文本
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的a标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.A(
//	    "点击这里",
//	    html.M{"color": "#007bff", "text-decoration": "none"},
//	    html.M{"href": "https://example.com", "target": "_blank", "class": "link"},
//	)
//	// 结果: <a style="color:#007bff;text-decoration:none;" href="https://example.com" target="_blank" class="link">点击这里</a>
func A(content template.HTML, ms ...M) template.HTML {
	return AEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// UlEl 创建并返回一个ul标签的Element实例
//
// ul标签用于创建无序列表，列表项使用li标签表示。
// 无序列表的每一项前面通常会有一个项目符号（如圆点）。
//
// 返回值：
//   - 返回一个tag为"ul"的Element实例
//
// 使用示例：
//
//	el := html.UlEl().SetContent("列表项").SetClass("list-group")
func UlEl() Element {
	return BaseEl().SetTag("ul")
}

// Ul 创建一个完整的ul标签
//
// 该函数创建一个包含指定内容和样式/属性的ul标签。
// ul标签用于创建无序列表，通常与li标签配合使用。
//
// 参数：
//   - content: ul标签的内容，通常包含多个li标签
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的ul标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Ul(
//	    "列表项内容",
//	    html.M{"list-style-type": "none", "padding": "0"},
//	    html.M{"class": "menu-list", "id": "main-menu"},
//	)
//	// 结果: <ul style="list-style-type:none;padding:0;" class="menu-list" id="main-menu">列表项内容</ul>
func Ul(content template.HTML, ms ...M) template.HTML {
	return UlEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// LiEl 创建并返回一个li标签的Element实例
//
// li标签用于表示列表项，必须包含在ul（无序列表）或ol（有序列表）标签内。
// li标签可以包含文本或其他HTML元素。
//
// 返回值：
//   - 返回一个tag为"li"的Element实例
//
// 使用示例：
//
//	el := html.LiEl().SetContent("列表项1").SetClass("list-item")
func LiEl() Element {
	return BaseEl().SetTag("li")
}

// Li 创建一个完整的li标签
//
// 该函数创建一个包含指定内容和样式/属性的li标签。
// li标签用于表示列表项，必须包含在ul或ol标签内。
//
// 参数：
//   - content: li标签的内容，可以是文本或其他HTML元素
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的li标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.Li(
//	    "列表项内容",
//	    html.M{"padding": "10px", "border-bottom": "1px solid #eee"},
//	    html.M{"class": "list-item", "data-id": "1"},
//	)
//	// 结果: <li style="padding:10px;border-bottom:1px solid #eee;" class="list-item" data-id="1">列表项内容</li>
func Li(content template.HTML, ms ...M) template.HTML {
	return LiEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// BEl 创建并返回一个b标签的Element实例
//
// b标签用于表示粗体文本，使文本以粗体显示。
// 在HTML5中，推荐使用CSS的font-weight属性来控制文本粗细。
//
// 返回值：
//   - 返回一个tag为"b"的Element实例
//
// 使用示例：
//
//	el := html.BEl().SetContent("粗体文本")
func BEl() Element {
	return BaseEl().SetTag("b")
}

// B 创建一个完整的b标签
//
// 该函数创建一个包含指定内容和样式/属性的b标签。
// b标签用于表示粗体文本，使文本以粗体显示。
//
// 参数：
//   - content: b标签的内容，通常是需要加粗的文本
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的b标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.B(
//	    "重要文本",
//	    html.M{},
//	    html.M{"class": "bold-text"},
//	)
//	// 结果: <b class="bold-text">重要文本</b>
func B(content template.HTML, ms ...M) template.HTML {
	return BEl().SetContent(content).SetStyleAndAttr(ms).Get()
}

// Br 创建一个或多个br标签（换行标签）
//
// br标签用于在文本中插入换行，是自闭合标签。
// 该函数可以创建指定数量的br标签。
//
// 参数：
//   - num: 可选参数，指定br标签的数量
//   - 如果不提供参数，默认创建1个br标签
//   - 如果提供参数，创建指定数量的br标签
//
// 返回值：
//   - 返回一个或多个br标签的HTML字符串
//
// 使用示例：
//
//	htmlStr1 := html.Br()           // 结果: <br>
//	htmlStr2 := html.Br(3)          // 结果: <br><br><br>
//	htmlStr3 := html.Br(2)          // 结果: <br><br>
func Br(num ...int) template.HTML {
	c := template.HTML("<br>")
	if len(num) > 0 {
		c = ""
		for range num {
			c += "<br>"
		}
	}
	return c
}

// H1El 创建并返回一个h1标签的Element实例
//
// h1标签用于表示一级标题，是HTML中最重要的标题。
// h1标题通常用于页面的主标题，字体最大。
//
// 返回值：
//   - 返回一个tag为"h1"的Element实例
//
// 使用示例：
//
//	el := html.H1El().SetContent("页面标题").SetStyle("color", "#333")
func H1El() Element {
	return BaseEl().SetTag("h1")
}

// H1 创建一个完整的h1标签
//
// 该函数创建一个包含指定内容和样式/属性的h1标签。
// h1标签用于表示一级标题，是HTML中最重要的标题级别。
//
// 参数：
//   - content: h1标签的内容，通常是页面主标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h1标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H1(
//	    "页面主标题",
//	    html.M{"color": "#333", "font-size": "32px"},
//	    html.M{"class": "main-title", "id": "page-title"},
//	)
//	// 结果: <h1 style="color:#333;font-size:32px;" class="main-title" id="page-title">页面主标题</h1>
func H1(content template.HTML, ms ...M) template.HTML {
	return H1El().SetContent(content).SetStyleAndAttr(ms).Get()
}

// H2El 创建并返回一个h2标签的Element实例
//
// h2标签用于表示二级标题，重要性仅次于h1。
// h2标题通常用于章节标题或子标题。
//
// 返回值：
//   - 返回一个tag为"h2"的Element实例
//
// 使用示例：
//
//	el := html.H2El().SetContent("章节标题").SetStyle("color", "#444")
func H2El() Element {
	return BaseEl().SetTag("h2")
}

// H2 创建一个完整的h2标签
//
// 该函数创建一个包含指定内容和样式/属性的h2标签。
// h2标签用于表示二级标题，通常用于章节标题或子标题。
//
// 参数：
//   - content: h2标签的内容，通常是章节标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h2标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H2(
//	    "章节标题",
//	    html.M{"color": "#444", "font-size": "28px"},
//	    html.M{"class": "section-title"},
//	)
//	// 结果: <h2 style="color:#444;font-size:28px;" class="section-title">章节标题</h2>
func H2(content template.HTML, ms ...M) template.HTML {
	return H2El().SetContent(content).SetStyleAndAttr(ms).Get()
}

// H3El 创建并返回一个h3标签的Element实例
//
// h3标签用于表示三级标题，重要性次于h2。
// h3标题通常用于子章节标题或小节标题。
//
// 返回值：
//   - 返回一个tag为"h3"的Element实例
//
// 使用示例：
//
//	el := html.H3El().SetContent("小节标题").SetStyle("color", "#555")
func H3El() Element {
	return BaseEl().SetTag("h3")
}

// H3 创建一个完整的h3标签
//
// 该函数创建一个包含指定内容和样式/属性的h3标签。
// h3标签用于表示三级标题，通常用于子章节标题或小节标题。
//
// 参数：
//   - content: h3标签的内容，通常是小节标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h3标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H3(
//	    "小节标题",
//	    html.M{"color": "#555", "font-size": "24px"},
//	    html.M{"class": "subsection-title"},
//	)
//	// 结果: <h3 style="color:#555;font-size:24px;" class="subsection-title">小节标题</h3>
func H3(content template.HTML, ms ...M) template.HTML {
	return H3El().SetContent(content).SetStyleAndAttr(ms).Get()
}

// H4El 创建并返回一个h4标签的Element实例
//
// h4标签用于表示四级标题，重要性次于h3。
// h4标题通常用于更细分的标题或副标题。
//
// 返回值：
//   - 返回一个tag为"h4"的Element实例
//
// 使用示例：
//
//	el := html.H4El().SetContent("副标题").SetStyle("color", "#666")
func H4El() Element {
	return BaseEl().SetTag("h4")
}

// H4 创建一个完整的h4标签
//
// 该函数创建一个包含指定内容和样式/属性的h4标签。
// h4标签用于表示四级标题，通常用于更细分的标题或副标题。
//
// 参数：
//   - content: h4标签的内容，通常是副标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h4标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H4(
//	    "副标题",
//	    html.M{"color": "#666", "font-size": "20px"},
//	    html.M{"class": "sub-subtitle"},
//	)
//	// 结果: <h4 style="color:#666;font-size:20px;" class="sub-subtitle">副标题</h4>
func H4(content template.HTML, ms ...M) template.HTML {
	return H4El().SetContent(content).SetStyleAndAttr(ms).Get()
}

// H5El 创建并返回一个h5标签的Element实例
//
// h5标签用于表示五级标题，重要性次于h4。
// h5标题通常用于较小的标题或说明性标题。
//
// 返回值：
//   - 返回一个tag为"h5"的Element实例
//
// 使用示例：
//
//	el := html.H5El().SetContent("说明标题").SetStyle("color", "#777")
func H5El() Element {
	return BaseEl().SetTag("h5")
}

// H5 创建一个完整的h5标签
//
// 该函数创建一个包含指定内容和样式/属性的h5标签。
// h5标签用于表示五级标题，通常用于较小的标题或说明性标题。
//
// 参数：
//   - content: h5标签的内容，通常是说明性标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h5标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H5(
//	    "说明标题",
//	    html.M{"color": "#777", "font-size": "18px"},
//	    html.M{"class": "description-title"},
//	)
//	// 结果: <h5 style="color:#777;font-size:18px;" class="description-title">说明标题</h5>
func H5(content template.HTML, ms ...M) template.HTML {
	return H5El().SetContent(content).SetStyleAndAttr(ms).Get()
}

// H6El 创建并返回一个h6标签的Element实例
//
// h6标签用于表示六级标题，是HTML中最小的标题级别。
// h6标题通常用于最小的标题或注释性标题。
//
// 返回值：
//   - 返回一个tag为"h6"的Element实例
//
// 使用示例：
//
//	el := html.H6El().SetContent("注释标题").SetStyle("color", "#888")
func H6El() Element {
	return BaseEl().SetTag("h6")
}

// H6 创建一个完整的h6标签
//
// 该函数创建一个包含指定内容和样式/属性的h6标签。
// h6标签用于表示六级标题，是HTML中最小的标题级别。
//
// 参数：
//   - content: h6标签的内容，通常是最小的标题或注释性标题
//   - ms: 可选参数，用于设置样式和属性
//   - ms[0]: 样式映射（M类型）
//   - ms[1]: 属性映射（M类型）
//
// 返回值：
//   - 返回完整的h6标签HTML字符串
//
// 使用示例：
//
//	htmlStr := html.H6(
//	    "注释标题",
//	    html.M{"color": "#888", "font-size": "16px"},
//	    html.M{"class": "note-title"},
//	)
//	// 结果: <h6 style="color:#888;font-size:16px;" class="note-title">注释标题</h6>
func H6(content template.HTML, ms ...M) template.HTML {
	return H6El().SetContent(content).SetStyleAndAttr(ms).Get()
}
