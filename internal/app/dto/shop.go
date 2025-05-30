package dto

type GetShopResp struct {
	Takeaway []TakeawayEntity `json:"takeaway"`
}

type TakeawayEntity struct {
	Tag   string        `json:"tag"`
	Name  string        `json:"name"`
	Foods []FoodsEntity `json:"foods"`
}

type FoodsEntity struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	LikeRatioDesc string   `json:"like_ratio_desc"`
	MonthSaled    int64    `json:"month_saled"`
	Unit          string   `json:"unit"`
	FoodTagList   []string `json:"food_tag_list"`
	Price         int64    `json:"price"`
	Picture       string   `json:"picture"`
	Description   string   `json:"description"`
	Tag           string   `json:"tag"`
}
