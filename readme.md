# golang 实现的 vnote 笔记本网页版
`+` `go` `vnote` `web编程`

练习实践 go 语言的 web 编程，将之前用 vnote 插件在 vim 中书写与管理的
markdown 笔记本用网页的形式阅读浏览。

## 基本思路

计划包含以下几个部分（package）：

* `main` 主框架，利用标准库 `net/http` 构建的基础 web 服务。
* `readcfg` 配置读取模块，从启动命令行及 toml 配置文件中读取配置，定制某些服务。
* `markdown` 转化为 html 的模块，预处理格式，缓存。
* `page` 页面模块，规划管理需要提供哪些页面。
* `notebook` 笔记本模块，有关笔记对象的逻辑。

## 页面设计

* `/` `/index` `/home` 平凡的封面主页
* `/yyyymmdd_n` 笔记内容页面，按模板转为 html
* `/tag/tag-name` 列出某个标签页下的笔记链接
* `/date/yyyy/mm/dd` 按日期列出笔记链接
* `/raw/yyyymmdd_n.md` 提供原始的 markdown 文本

## 后期扩展功能

* 统计
* 管理
* 搜索
