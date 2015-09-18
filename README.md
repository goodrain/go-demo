# Go Web 示例程序

## 项目介绍

本项目是一个简单的 [Golang](http://golang.org) Web 应用示例，目录结构：

```
.
├── Godeps
│   ├── Godeps.json
│   ├── Readme
│   └── _workspace
├── hello
│   └── main.go
├── Procfile
├── README.md
└── static
    └── index.html
```

## 项目要求

项目的根目录下需要有 `Godeps/Godeps.json` 来指定 `Golang` 版本以及相关依赖，推荐使用 [Godep](https://github.com/tools/godep) 来管理项目依赖可以加快项目部署速度。同时还需要一个 `Procfile` 文件来指定应用的启动命令，对于 go 来说，这个命令就是可执行文件名。

