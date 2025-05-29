package dto

type GetUserReq struct {
	Id int `json:"id"`
}
type GetUserResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResp struct {
	Token string `json:"token"`
}
