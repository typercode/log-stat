# java 程序员来学习go语言


##比较鲜明的几点：

1. 函数可以很方便的返回多个值，而java如果需要返回多个值，还需要封装一个类出来包装一下
1. 函数定义时，参数名称在前，类型在后，跟java相反
1. 异常、错误处理都是通过nil判断，没有try catch finally throw throws这些，所以go的关键字也很少，学习成本低
1. 没有java那么明显的继承关系
1. go支持交叉编译，直接编译成可支持文件，所以不需要目标机器安装类似jvm的环境，也不用处理复杂的依赖关系（java 通过maven来处理）
1. go有自己的go fmt，统一的代码格式
1. 定义、声明未使用的变量，会直接编译报错，包括没有使用到的包，这个在java上深有体会，经常有好多不实用的包、变量，还在代码里，这就是坏味道
1. := 可以把变量类型、声明和赋值都搞定了（java里没有，提一下）
1. var m = make(map[string][]string) ,这种定义方式表示声明、初始化一个map类型的m，key是string，value是[]string，表示slice，大小是动态的。

## 未体验到的

1. go的并发编程



# logStat
logStat for text file

## Install

`go get github.com/qiaodaimadelaowang/log-stat/...`

## Usage

```bash
➜  Downloads ./logStat -h
Usage of ./logStat:
  -f string
    	set the log file name(eg: Log.log.2019-11-03)
  -fp string
    	set the log file prefix,Ignore if f is set. (eg: MSSM-Auth.log.) (default "MSSM-Auth.log.")
  -ma string
    	set the mail address (eg: xxx@xxx.com)
  -p string
    	set the log file path(eg: /path/to/file/logs/app-name)
    	
    	
➜  ./logStat -p "/Users/tinyhuiwang/temp/a" -fp MSSM-Auth.log.

2019-11-03
filePath:/Users/tinyhuiwang/temp/a/MSSM-Auth.log.2019-11-03
mailContent:
appId: 1
api-version-id: 2,3

appId: 2
api-version-id: ,

appId: 3
api-version-id: 2,

appId: 1
api-version-id: 2,

appId: 2
api-version-id: 3,

appId: 1
api-version-id: 17,

appId: 2
api-version-id: ,

appId: 2
api-version-id: 1,

appId: 3
api-version-id: 17,47,

appId:
api-version-id: 1,

appId: id1
api-version-id: ,

appId: id2
api-version-id: 1,2,
```