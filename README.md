# BDMSystem_kratos


###基于kratos微服务框架数据分析管理系统

#### 1、安装go
下载go语言安装包；
环境变量配置：
GOPATH
D:\Dev-C\go-workspace
GOROOT
D:\Dev-C\go\

#### 2、安装protocbuf
下载protoc-3.11.2-win64.zip解压；
将文件放入D:\Dev-C\protobuf-3.11.2即可；

#### 3、安装kratos
下载源码；
进入文件夹D:\Dev-C\kratos\kratos\tool\kratos；
在cmd中执行：go build 命令；
生成 kratos.exe 文件；
放入文件夹D:\Dev-C\go-workspace\bin
完成

#### 4、new demo
看教程即可
一般在D:\Dev-C\go-workspace\src中创建项目等等

###环境变量配置
Path
D:\Dev-C\go\bin
D:\Dev-C\go-workspace\bin
D:\Dev-C\protobuf-3.11.2\bin

###介绍
api.bm.go 为http的对外接口, BM server即blademaster为热度http框架gin的裁剪.去除了gin中不需要的部分逻辑
api.bm.go 、apb.pb.go 可以通过kartos tool生成
(kratos tool可以基于proto生成http&grpc代码，生成缓存回源代码，生成memcache执行代码，生成swagger文档等工具集) bm、 pb 分别为http和grpc的接口。


###Discovery注册中心
####步骤
#####1、安装discovery，先git clone下来
#####2、进入文件夹discovery/cmd/discovery
#####3、用下面命令进行build，然后生成exe文件
#####4、把exe文件放入go_workspace/bin中
#####5、在有discovery-example.toml文件的目录下
#####6、执行运行命令

####build
#####cd $GOPATH/src
#####git clone https://github.com/bilibili/discovery.git
#####cd discovery/cmd/discovery
#####go build

运行discovery run
discovery -conf discovery-example.toml -alsologtostderr

discovery -conf discovery-example.toml -log.dir="/tmp"   //两种命令，是输入日志文件的方式不同（第一种需要插件，第二种直接输出文件）


set DISCOVERY_NODES=127.0.0.1:7171 //设置节点地址命令(环境变量也可以)

服务器地址:47.107.116.160:80



###phoclus 配置

因为phoclus编译时，依赖gcc编译器

使用了tdm64-gcc，打包了各类gcc编译器
    下载安装tdm64-gcc-5.1.0-2.exe
	配置其环境变量 Path:D:\Dev-C\gcd\bin
	安装结束

从github上clone下phoclus源码

进入D:\Dev-C\go_workspace\src\github.com\henrylee2cn\pholcus文件夹

执行go build -ldflags="-linkmode internal"

生成phoclus.exe文件

可执行文件

###kratos tool swagger使用
####//生成 Swagger文档
####//go:generate kratos tool protoc --swagger api.proto
####//读取Swagger文档
####//go:generate kratos tool swagger serve api.swagger.json --flavor=swagger
执行命令后，浏览器会自动打开swagger文档地址。
同时也可以查看更多的 go-swagger官方（https://github.com/go-swagger/go-swagger）参数进行使用。

###服务器端口可用分配记录
####discovery
##### http:5501
####token
##### grpc:5001
####bdms
##### http:5002~5100
##### grpc:5502~5600
####worknet
##### http:5101~5200
##### grpc:5601~5700
