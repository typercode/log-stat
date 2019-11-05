# java 程序员来学习go语言


## 比较鲜明的几点

1. 函数可以很方便的返回多个值，而java如果需要返回多个值，还需要封装一个类出来包装一下
1. go是强类型的语言，函数定义时，参数名称在前，类型在后，返回值声明也是放在方法参数后面，跟java相反
1. 异常、错误处理都是通过nil判断，没有try catch finally throw throws这些，所以go的关键字也很少，学习成本低
1. 没有java那么明显的继承关系 
1. go支持交叉编译，直接编译成可支持文件，所以不需要目标机器安装类似jvm的环境，也不用处理复杂的依赖关系（java 通过maven来处理）
1. go有自己的go fmt，统一的代码格式，这一点太爽了，大家统一风格，没有好坏之分，都是走官方的，也不用争了，例如方法、if语句、for语句的花括号必须跟语句在同一行，另起一行直接编译报错，不管你喜不喜欢。
1. 定义、声明未使用的变量，会直接编译报错，包括没有使用到的包，这个在java上深有体会，经常有好多不实用的包、变量，还在代码里，这就是坏味道
1. := 可以把变量类型、声明和赋值都搞定了（java里没有，提一下）
1. var m = make(map[string][]string) ,这种定义方式表示声明、初始化一个map类型的m，key是string，value是[]string，表示slice，大小是动态的。
1. 交叉编译：CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o logStat main.go
1. go的权限访问控制是通过首字母是否大小写来声明的
1. go对于命令行接收参数非常友好，内置了flag包，还提供了默认的-h选项来查看命令的使用方式。
1. 日期格式化不是类似java的yyyy-MM-dd HH:mm:ss 而是用的2006-01-02 15:04:05 这个固定的日期来格式化
1. java调用其他包中类的方法是通过类/对象.方法，go是通过包名.方法，所以一个包下面定义的方法名称都不能有重复的，
   不然编译报错： xxxFunc redeclared in this package，但是main函数除外。main函数只有在main 包下才能运行，
   package packageName
   这里的packageName可以不跟目录一一对应，但是建议保持一致。
1. 结构体定义、声明、赋值
    ```
    //定义
    type Car struct {
        weight int
        name   string
    }
    //声明
    var car Car
    //car := new(Car) 这样也可以
    
    //赋值
    car.name = "特斯拉"
    car.weight = 490
    ```
    可以看到，声明之后就可以直接赋值了，不是类似java必须通过new来创建，不然car是null，.的时候报NPE
    
1. go的继承
    ```
    java是通过extends、implements 关键字来表示继承关系，但是go没有
    
        //给Car结构体增加方法
        func (i *Car) testMethod(){
            fmt.Println("testMethod is called!!!")
        }
    
        type Bike struct {
             //只写结构体名称，不写变量名称
              Car //继承Car
              name string
        }
     
    这样就继承了Car 的属性和方法
     
         var bike Bike
         //方法就继承过来了
         bike.testMethod()
         
    多重继承
        type Bus struct {
            people int
        }
        
        type Bike struct {
              Car //继承Car
              Bus //继承Bus
              name string
        }
        
        这个时候，Bus结构体的方法、属性就不能跟Car相同，不然就在Bike调用的时候出现模凌两可的问题，ambiguous reference 'funcNam'，编译器报错。
        
         
    
    ```


欢迎pr，哈哈

## 未体验到的

1. go的并发编程
1. web开发
1. 包管理
1. 单元测试



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


## 我的学习路线

1. 4、5月份的时候公司搞了一个go培训 讲师的github: https://github.com/bingoohuang 
1. https://learnxinyminutes.com/docs/zh-cn/lua-cn/
1. google
这里面的代码，基本是面向google编程来的。