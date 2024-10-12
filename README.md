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