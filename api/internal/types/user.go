package types

type CreateUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResp struct {
	Resp
	Data struct {
		UserUuid string `json:"user_uuid"`
		Username string `json:"username"`
	} `json:"data"`
}
