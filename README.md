# HTML构造器

一个简洁、高效的HTML元素构造器库，用于通过Go代码动态生成HTML标签。

## 项目概述

HTML构造器是一个轻量级的Go语言库，提供了一组简洁的API，用于构建HTML元素，支持链式调用和灵活的样式/属性设置。该库设计用于简化Go模板中的HTML生成，特别适合需要动态创建HTML内容的Web应用程序。

### 主要特性

- ✅ 支持常见的HTML标签（div、span、button、form、a、ul、li等）
- ✅ 支持内联样式和HTML属性设置
- ✅ 支持元素的嵌套和组合
- ✅ 使用template.HTML类型确保HTML内容的安全输出
- ✅ 提供简洁的API，支持链式调用
- ✅ 零依赖，易于集成
- ✅ 完全兼容Go的html/template包

## 安装指南

### 前提条件

- Go 1.13或更高版本

### 安装方式

使用`go get`命令安装：

```bash
go get github.com/GoAdminGroup/html
```

## 快速开始

### 基本使用

```go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GoAdminGroup/html"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 创建一个简单的div元素
	simpleDiv := html.Div("Hello World")
	fmt.Println(simpleDiv)

	// 创建带有样式和属性的div元素
	styledDiv := html.Div("内容",
		html.M{"color": "red", "font-size": "14px"},  // 样式
		html.M{"class": "container", "id": "main"})   // 属性
	fmt.Println(styledDiv)

	// 创建嵌套的HTML结构
	nestedHtml := html.Div(
		html.A("点击这里",
			html.M{"color": "blue"},
			html.M{"href": "https://example.com"}),
		html.M{"text-align": "center"},
		html.M{"class": "wrapper"})
	fmt.Println(nestedHtml)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

### 在Go模板中使用

```go
package main

import (
	"html/template"
	"net/http"

	"github.com/GoAdminGroup/html"
)

var tpl = template.Must(template.New("example").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>HTML构造器示例</title>
</head>
<body>
    {{.}}
</body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	// 使用HTML构造器创建页面内容
	pageContent := html.Div(
		html.H1("欢迎使用HTML构造器",
			html.M{"color": "#333", "text-align": "center"}),
		html.P("这是一个使用HTML构造器创建的段落",
			html.M{"line-height": "1.6"}),
		html.Button("点击按钮",
			html.M{"padding": "10px 20px", "background-color": "#007bff"},
			html.M{"class": "btn btn-primary", "onclick": "alert('按钮被点击了！')"}),
		html.M{"max-width": "800px", "margin": "0 auto", "padding": "20px"},
	)

	// 渲染模板
	tpl.Execute(w, pageContent)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

## 核心概念

### Element结构体

Element是HTML元素的核心表示，包含以下字段：

| 字段名 | 类型 | 描述 |
| ------ | ---- | ---- |
| Tag | template.HTML | HTML标签名，如"div"、"span"等 |
| Content | template.HTML | 元素的内容，可以是文本或其他HTML元素 |
| Style | Style | 元素的内联样式 |
| Attribute | Attribute | 元素的HTML属性 |

### M类型

M是一个字符串映射类型的别名，用于简化样式和属性的定义：

```go
type M map[string]string
```

使用示例：

```go
style := html.M{"color": "red", "font-size": "14px"}
attr := html.M{"class": "btn", "id": "submit"}
```

### Style和Attribute类型

- `Style`：表示HTML元素的样式
- `Attribute`：表示HTML元素的属性

这两个类型都提供了`String()`方法，可以将映射转换为HTML字符串。

## API文档

### 基础API

#### BaseEl()

创建并返回一个基础的Element实例：

```go
func BaseEl() Element
```

#### SetTag()

设置元素的HTML标签名：

```go
func (b Element) SetTag(tag template.HTML) Element
```

#### SetContent()

设置元素的内容：

```go
func (b Element) SetContent(content template.HTML) Element
```

#### SetStyle()

设置元素的单个样式属性：

```go
func (b Element) SetStyle(key, value string) Element
```

#### SetAttr()

设置元素的单个属性：

```go
func (b Element) SetAttr(key, value string) Element
```

#### SetStyleAndAttr()

同时设置元素的样式和属性：

```go
func (b Element) SetStyleAndAttr(ms []M) Element
```

#### Get()

将Element转换为完整的HTML标签字符串：

```go
func (b Element) Get() template.HTML
```

### 便捷函数

以下是一些常用HTML标签的便捷函数，它们直接返回完整的HTML字符串：

#### Body()

创建一个完整的body标签：

```go
func Body(content template.HTML, ms ...M) template.HTML
```

#### Div()

创建一个完整的div标签：

```go
func Div(content template.HTML, ms ...M) template.HTML
```

#### Span()

创建一个完整的span标签：

```go
func Span(content template.HTML, ms ...M) template.HTML
```

#### Button()

创建一个完整的button标签：

```go
func Button(content template.HTML, ms ...M) template.HTML
```

#### Form()

创建一个完整的form标签：

```go
func Form(content template.HTML, ms ...M) template.HTML
```

#### A()

创建一个完整的a标签：

```go
func A(content template.HTML, ms ...M) template.HTML
```

#### Ul()

创建一个完整的ul标签：

```go
func Ul(content template.HTML, ms ...M) template.HTML
```

#### Li()

创建一个完整的li标签：

```go
func Li(content template.HTML, ms ...M) template.HTML
```

#### Br()

创建一个或多个br标签：

```go
func Br(num ...int) template.HTML
```

#### 标题标签

```go
func H1(content template.HTML, ms ...M) template.HTML
func H2(content template.HTML, ms ...M) template.HTML
func H3(content template.HTML, ms ...M) template.HTML
func H4(content template.HTML, ms ...M) template.HTML
func H5(content template.HTML, ms ...M) template.HTML
func H6(content template.HTML, ms ...M) template.HTML
```

### 所有支持的标签

该库支持以下HTML标签，每个标签都有对应的便捷函数：

- body
- div
- span
- i (图标)
- p (段落)
- button
- form
- a (超链接)
- ul (无序列表)
- li (列表项)
- b (粗体)
- br (换行)
- h1, h2, h3, h4, h5, h6 (标题)

## 使用示例

### 创建复杂的HTML结构

```go
// 创建一个导航菜单
navMenu := html.Ul(
	html.Li(
		html.A("首页",
			html.M{"color": "#333"},
			html.M{"href": "/", "class": "nav-link"})
	) +
	html.Li(
		html.A("关于我们",
			html.M{"color": "#333"},
			html.M{"href": "/about", "class": "nav-link"})
	) +
	html.Li(
		html.A("服务",
			html.M{"color": "#333"},
			html.M{"href": "/services", "class": "nav-link"})
	) +
	html.Li(
		html.A("联系我们",
			html.M{"color": "#333"},
			html.M{"href": "/contact", "class": "nav-link"})
	),
	html.M{"list-style-type": "none", "display": "flex", "gap": "20px"},
	html.M{"class": "navbar-nav", "id": "main-nav"}
)
```

### 创建表单

```go
// 创建一个登录表单
loginForm := html.Form(
	html.Div(
		html.Label("用户名",
			html.M{"display": "block", "margin-bottom": "5px"},
			html.M{"for": "username"})
		+ html.Input(
			html.M{"width": "100%", "padding": "10px", "margin-bottom": "15px"},
			html.M{"type": "text", "id": "username", "name": "username", "placeholder": "请输入用户名"}),
		html.M{"margin-bottom": "15px"}
	) +
	html.Div(
		html.Label("密码",
			html.M{"display": "block", "margin-bottom": "5px"},
			html.M{"for": "password"})
		+ html.Input(
			html.M{"width": "100%", "padding": "10px", "margin-bottom": "15px"},
			html.M{"type": "password", "id": "password", "name": "password", "placeholder": "请输入密码"}),
		html.M{"margin-bottom": "15px"}
	) +
	html.Button("登录",
		html.M{"width": "100%", "padding": "10px", "background-color": "#007bff", "color": "white", "border": "none", "border-radius": "4px"},
		html.M{"type": "submit", "class": "btn btn-primary"})
	+
	html.A("忘记密码?",
		html.M{"color": "#007bff", "text-decoration": "none", "display": "block", "margin-top": "10px", "text-align": "center"},
		html.M{"href": "/forgot-password"}),
	html.M{"max-width": "400px", "margin": "0 auto", "padding": "20px", "border": "1px solid #ddd", "border-radius": "8px"},
	html.M{"action": "/login", "method": "post", "class": "login-form"}
)
```

## 常见问题解答

### Q: 如何创建自定义HTML标签？

A: 可以使用`BaseEl()`函数创建基础元素，然后通过`SetTag()`方法设置自定义标签名：

```go
customEl := html.BaseEl().SetTag("custom-tag").SetContent("自定义标签内容").Get()
```

### Q: 如何处理HTML转义？

A: 该库使用`template.HTML`类型，确保HTML内容的安全输出。如果需要直接输出HTML内容而不转义，可以使用`template.HTML()`函数将字符串转换为`template.HTML`类型：

```go
rawHtml := template.HTML("<script>alert('hello');</script>")
div := html.Div(rawHtml)
```

### Q: 如何添加多个class？

A: 可以使用`SetClass()`方法，支持添加多个class名称：

```go
// 使用SetClass方法
el := html.BaseEl().SetTag("div").SetClass("class1", "class2", "class3")

// 或使用M类型
el := html.Div("内容",
	html.M{},
	html.M{"class": "class1 class2 class3"})
```

## 贡献指南

### 贡献流程

1. Fork 仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

### 开发环境设置

1. 克隆仓库：
   ```bash
   git clone https://github.com/GoAdminGroup/html.git
   cd html
   ```

2. 运行测试：
   ```bash
   go test -v
   ```

3. 格式化代码：
   ```bash
   gofmt -w .
   ```

### 测试指南

- 所有新功能都应该添加对应的测试用例
- 测试文件应该放在与源代码相同的目录中，命名为`*_test.go`
- 运行所有测试：`go test -v`

## 许可证信息

该项目使用MIT许可证，详情请查看[LICENSE](LICENSE)文件。

## 联系方式

- 项目地址：https://github.com/GoAdminGroup/html
- 问题反馈：https://github.com/GoAdminGroup/html/issues
- 贡献代码：https://github.com/GoAdminGroup/html/pulls

## 致谢

感谢所有为该项目做出贡献的开发者！

---

**更新日期**：2026-01-05
**版本**：v1.0.0
