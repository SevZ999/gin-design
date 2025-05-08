package dto

type GetUserReq struct {
	Id int `json:"id"`
}
type GetUserResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
