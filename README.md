## 新建项目
```
goctl quickstart -t micro
```

## 新建一个logginsvc的rpc服务
```
goctl rpc new logginsvc
```


## 链路追踪
### go-zero默认是开启的，但是需要配置etcd或者consul才可以生效

# api

### 提供对外的api服务，http

# loginsvc
### 登录服务，rpc

# usersvc
### 用户管理服务，rpc

# rpc结构
```
etc -- 配置文件
internal -- 内部文件
    config  -- 配置文件映射
    logic  -- 业务逻辑代码
    server  -- rpc服务接口定义，自己新增的接口要在这里定义
    svc  -- 上下文
loginsvc  -- proto生成的文件
loginsvcclient  -- rpc客户端文件，自己新增的接口要在这里实现调用
```

# http结构
```
etc -- 配置文件，调用的rpc服务，需要在这里配置
internal -- 内部文件
    config  -- 配置文件映射，调用的rpc服务，需要在这里配置
    handler  -- 路由注册，解析参数，调用logic层
    logic  -- 业务逻辑代码，调用rpc服务
    svc -- 上下文，调用的rpc服务，需要在这里定义好
    types  -- http接口接收、返回参数定义
```