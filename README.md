## 新建项目
```
goctl quickstart -t micro
```

## 新建一个logginsvc的rpc服务
```
goctl rpc new logginsvc
```

## 生成proto
```
protoc *.proto --go_out=. --go-grpc_out=.
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

# 权限认证,headers中添加Authorization
```
server.AddRoutes(
		routers,
		rest.WithPrefix("/api"),
		rest.WithJwt(ctx.Config.JwtAuth.AccessSecret),
	)
```

# 生成token
```
// 获取token，token中携带用户的uuid
token, err := GenJwtToken(in.AccessExpire, in.AccessSecret, map[string]interface{}{"uid": userUuid})

func GenJwtToken(accessExpire int64, accessSecret string, payloads map[string]interface{}) (string, error) {

	now := time.Now().Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(accessSecret))
}

```

# 认证成功后token中的uid怎么获取？？？？
```
l.ctx.Value("uid")
```

# rcp怎么统一认证,TLS
```
// api，配置文件
UserSvc:
  Etcd:
    Hosts:
    - 192.168.2.235:2379
    Key: usersvc.rpc

// 加载密钥
func loadTLSCredentials(cert *string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(*cert + "/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

// 使用密钥
func NewServiceContext(c config.Config, cert *string) *ServiceContext {

	tlsCredentials, err := loadTLSCredentials(cert)
	if err != nil {
		panic(err)
	}

    // WithTransportCredentials使用密钥
	return &ServiceContext{
		Config:   c,
		LoginSvc: loginsvcclient.NewLoginsvc(zrpc.MustNewClient(c.LoginSvc, zrpc.WithTransportCredentials(tlsCredentials))),
		UserSvc:  usersvcclient.NewUsersvc(zrpc.MustNewClient(c.UserSvc, zrpc.WithTransportCredentials(tlsCredentials))),
	}
}
```

```
// rpc，不需要配置文件,直接加载密钥文件
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(*cert+"/server-cert.pem", *cert+"/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	c := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(c), nil
}

// 使用密钥文件启动
s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		loginsvc.RegisterLoginsvcServer(grpcServer, server.NewLoginsvcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	tlsCfg, err := loadTLSCredentials()
	if err != nil {
		panic(err)
	}

	s.AddOptions(grpc.Creds(tlsCfg))



```
