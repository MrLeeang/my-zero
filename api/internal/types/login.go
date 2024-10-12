package types

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Resp
	Data struct {
		Token    string `json:"token"`
		UserUuid string `json:"user_uuid"`
	} `json:"data"`
}
