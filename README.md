# week-04
go course week-04 code

### 此为一个简单例子，暂未做配置信息等（参考了kratos v2）

#### /api  放置对外接口的定义信息
```
通过proto文件生成对应的关联信息，
对应目录下，根据项目和版本号放置，如：
address_book项目的v1版本，对应的文件路径为api/address_book/v1
```

#### /cmd 放置服务启动脚本，主要放置启动逻辑
```
根据项目放置，如：address_book项目,
启动一个其中服务，就会在cmd目录下方创建cmd/address_book/
里面会有一个main.go用于生成对应的可执行文件address_book
用于启动整个项目

这里用了wire来将内置的服务进行依赖注入
```

#### /internal 放置不希望被外部的包引用的内容
```
/data     放置数据读取逻辑，从数据库读取或者从缓存读取
/biz      放置业务逻辑
/server   放置grpc服务注册逻辑
/service  将biz业务逻辑的合并和部分处理
```

#### /pkg 放置可以被外部引用的包的内容

#### /third_party 放置依赖的第三方proto内容