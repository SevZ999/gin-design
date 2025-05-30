package service

import (
	"loan-admin/internal/app/dto"
)

type ShopRepo interface {
}

type ShopService struct {
	repo ShopRepo
}

func NewShopService(repo ShopRepo) *ShopService {
	return &ShopService{repo: repo}
}

func (srv *ShopService) GetShop() (dto.GetShopResp, error) {
	return dto.GetShopResp{
		Takeaway: []dto.TakeawayEntity{
			{
				Tag:  "318569657",
				Name: "一人套餐",
				Foods: []dto.FoodsEntity{
					{
						Id:            8078956697,
						Name:          "烤羊肉串(10串)",
						LikeRatioDesc: "好评度100%",
						MonthSaled:    100,
						Unit:          "10串",
						FoodTagList:   []string{"点评网友推荐"},
						Price:         90,
						Picture:       "https://ts1.tc.mm.bing.net/th/id/R-C.d3fe3a1eed0a3ad1bab953eb33ab20e8?rik=135d3iisVsIl8w&riu=http%3a%2f%2fcp1.douguo.net%2fupload%2fcaiku%2f6%2f2%2fc%2fyuan_623aff4797c3a740738ebab21d97665c.jpg&ehk=ZGSvsFGB7DY%2b%2flYntA5MTGf4Jve9FtpKEHZQKPkY1cc%3d&risl=&pid=ImgRaw&r=0",
						Description:   "描",
						Tag:           "318569657",
					},
				},
			},
			{
				Tag:  "318569658",
				Name: "特色烧烤",
				Foods: []dto.FoodsEntity{
					{
						Id:            8078956698,
						Name:          "烤羊肉串(10串)",
						LikeRatioDesc: "好评度100%",
						MonthSaled:    100,
						Unit:          "10串",
						FoodTagList:   []string{"点评网友推荐"},
						Price:         90,
						Picture:       "https://ts1.tc.mm.bing.net/th/id/R-C.d3fe3a1eed0a3ad1bab953eb33ab20e8?rik=135d3iisVsIl8w&riu=http%3a%2f%2fcp1.douguo.net%2fupload%2fcaiku%2f6%2f2%2fc%2fyuan_623aff4797c3a740738ebab21d97665c.jpg&ehk=ZGSvsFGB7DY%2b%2flYntA5MTGf4Jve9FtpKEHZQKPkY1cc%3d&risl=&pid=ImgRaw&r=0",
						Description:   "描",
						Tag:           "318569658",
					},
				},
			},
			{
				Tag:  "318569659",
				Name: "杂粮主食",
				Foods: []dto.FoodsEntity{
					{
						Id:            8078956699,
						Name:          "烤饼(2串)",
						LikeRatioDesc: "好评度100%",
						MonthSaled:    100,
						Unit:          "10串",
						FoodTagList:   []string{"点评网友推荐"},
						Price:         10,
						Picture:       "https://ts1.tc.mm.bing.net/th/id/R-C.d3fe3a1eed0a3ad1bab953eb33ab20e8?rik=135d3iisVsIl8w&riu=http%3a%2f%2fcp1.douguo.net%2fupload%2fcaiku%2f6%2f2%2fc%2fyuan_623aff4797c3a740738ebab21d97665c.jpg&ehk=ZGSvsFGB7DY%2b%2flYntA5MTGf4Jve9FtpKEHZQKPkY1cc%3d&risl=&pid=ImgRaw&r=0",
						Description:   "描",
						Tag:           "3185696579",
					},
				},
			},
		},
	}, nil
}
