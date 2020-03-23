# BDMSystem_kratos
基于kratos微服务框架数据分析管理系统

Discovery注册中心

安装discovery，先git clone下来，进入文件夹discovery/cmd/discovery
用下面命令进行build，然后生成exe文件，把exe文件放入go_workspace/bin中
在有discovery-example.toml文件的目录下，执行运行命令

build
cd $GOPATH/src
git clone https://github.com/bilibili/discovery.git
cd discovery/cmd/discovery
go build

运行discovery run
discovery -conf discovery-example.toml -alsologtostderr

discovery -conf discovery-example.toml -log.dir="/tmp"   //两种命令，是输入日志文件的方式不同（第一种需要插件，第二种直接输出文件）


set DISCOVERY_NODES=127.0.0.1:7171 //设置节点地址命令(环境变量也可以)

服务器地址:47.107.116.160:80


#discovery和token挂载在服务器上操作记录

discovery服务发现
修改本地环境变量 DISCOVERY_NODES 
由127.0.0.1:7171
改为
47.107.116.160:80

