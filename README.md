# GoBuildDocsifySidebar
> 使用golang 生成 docsify 的sidebar

# 仅供学习
日常使用请选择👉 [C# 版本源码](https://github.com/xxxxue/Docsify-Build-Sidebar)

## ~~使用方法~~
> ~~使用前先将 `/config/config.json` 中的 `HomePath` 
> 改为本机docsify项目的根目录~~

~~1. 自己用go源码编译~~

~~2. 在 `Releases` 中 下载相应系统的文件~~


## 编译
> 在项目根目录
>
> 一次执行一行命令 
>
> cmd和powershell  命令有一点区别

CMD

```shell
set GOARCH=amd64
set GOOS=linux
go build
```

PowerShell

```shell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build

$env:GOOS="windows"
$env:GOARCH="amd64"
go build

$env:GOOS="darwin"
$env:GOARCH="amd64"
go build
```


