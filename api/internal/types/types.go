package types

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Ok             = 0
	LoginError     = 1001
	AuthError      = 1002
	RpcError       = 1003
	NotFoundError  = 1004
	ExceptionError = 1005
	LicenseError   = 1006
)

var (
	ErrorCodeMessage = map[int]string{
		0:    "SUCCESS",
		1001: "登录失败",
		1002: "权限验证失败",
		1003: "微服务调用异常",
		1004: "资源未找到",
		1005: "异常错误",
		1006: "产品许可证错误或过期",
	}
)
