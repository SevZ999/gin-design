package service

import (
	"gin-design/internal/app/dto"
	"gin-design/internal/app/mock"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := mock.NewMockUserRepo(ctrl)

	t.Run("empty", func(t *testing.T) {
		user.EXPECT().GetUser("").Return(dto.GetUserResp{}, nil)
	})
}
