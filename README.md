# Scaffold
Gin best practices, gin development scaffolding.

使用gin构建的开发脚手架，辅助进行高效web开发。
主要功能有：

1. 请求链路日志打印
2. 支持中英文错误信息提示及自定义错误提示。
3. 支持了多配置环境

项目地址：https://github.com/m2tB/Scaffold
### 现在开始
- 安装软件依赖
go mod使用请查阅：

https://blog.csdn.net/e421083458/article/details/89762113
```
git clone git@github.com:m2tB/Scaffold.git
cd Scaffold
go mod tidy
```
- 确保正确配置了 conf/ 目录下的config文件


### 文件分层
```
├── README.md
├── conf            		配置文件夹
├── cmd             		项目启动入口
├── internal        		业务代码文件夹
│   └── icontroller 		控制器
│   └── imapper     		数据库操作
│   └── imiddleware 		中间件
│   └── imodel      		结构体
│   └── initialize  		初始化
│   └── iresponse   		统一响应
│   └── iroutes     		路由
│   └── startup     		启动信息
├── logging         		日志文件存储位置
├── .gitignore      		工具文件夹
├── go.mod           		工具文件夹
├── go.sum           		工具文件夹
└── utils           		工具文件夹
```
