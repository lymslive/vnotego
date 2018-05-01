# vnote go 开发笔记

## 使用的第三方库

* https://github.com/gomarkdown/markdown
* https://github.com/pelletier/go-toml

## 收集备用的库

* https://github.com/BurntSushi/toml
* https://github.com/russross/blackfriday (也是 markdown 转化的)

markdown css 主题推荐

* https://www.jianshu.com/p/18876655b452 (vscode)
* http://jasonm23.github.io/markdown-css-themes/
* https://github.com/markdowncss/splendor
* http://markdowncss.github.io/

感觉下载的那个 vscode 可将就，对中文显示较好。

## 参考学习资料

Go Template 使用简介
http://www.cnblogs.com/52php/p/6412554.html

## 编码陷阱

### "notepost" is an incomplete or empty template

template.ParseFiles 生成的模板会以文件名命名，这会用个使用陷阱，
不能先 New 或在执行时明确使用 ExecuteTemplate 方法。参考：
https://stackoverflow.com/questions/10199219/go-template-function

### http.FileServer 404

不是挂在根目录下的静态文件服务，须再包一层 StripPrefix ，否则 404
http.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(path))))

但除非是网络子目录名与服务器子目录名不同，才需要 http.StripPrefix 辅助。
可以直接挂在根目录下，但只有访问某个子目录才开放文件系统。
因为这时网络路径（从 / 开始算）仍与服务器路径相同。

## BUG

有一篇旧笔记是软连接，下载到另一电脑时不可用，但笔记文件名还在。
扫描日志时异常，无法获得标题标签信息，更无法读取内容。
