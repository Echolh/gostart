# gostart

## 介绍

Go学习的开始， 请在此目录下创建你的第一个Go项目，并持续记录学习过程中的代码示例。



## 流程示例

### 1. 创建自己的文件夹

在根目录下创建自己命名的文件夹

![image-20240227101726543](https://sso.image-lh.love/image-20240227101726543.png)

### 2. 初始化module

第一次创建项目时需要初始化mod,在终端中执行 `go mod init xxxx`命令。xxx请用自己第一步创建的文件夹替换。如：

```shell
go mod init lihang
```

初始化成功后，会在根目录生成一个go.mod文件



### 3.创建Hello world 程序

在用户目录下创建 hello-world 文件夹，在 hello-world 文件夹下创建main.go文件

```go
package main	

import "fmt"

func main() {
	fmt.Print("Hello World!")
}

```

### 3. 执行程序

进入用户目录,在终端中执行`go run xxx`命令,输出Hello World!

```shell
$ cd lihang

$ go run hello-world/main.go  
  Hello World!
```



### 4. 注意事项

每一个练习项目都需要单独创建文件夹进行区分，如切片示例程序，创建lihang/slice目录

![image-20240227102908032](https://sso.image-lh.love/image-20240227102908032.png)
