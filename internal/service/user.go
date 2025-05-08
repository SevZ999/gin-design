package service

type UserRepo interface {
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}
