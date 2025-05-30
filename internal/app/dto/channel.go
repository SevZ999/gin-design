package dto

type GetChannelResp struct {
	Channels []Channel `json:"channels"`
}
type Channel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
