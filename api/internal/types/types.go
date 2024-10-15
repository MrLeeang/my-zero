package types

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Ok             = 0
	RpcError       = 1001
	NotFoundError  = 1002
	ExceptionError = 1003
	LicenseError   = 1004
)

var (
	ErrorCodeMessage = map[int]string{
		0:    "SUCCESS",
		1001: "微服务调用异常",
		1002: "资源未找到",
		1003: "异常错误",
		1004: "产品许可证错误或过期",
	}
)
