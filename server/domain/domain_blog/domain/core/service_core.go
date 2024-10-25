package core

type CoreService struct {
	*UserService
	*AboutService
}

func NewCoreService() *CoreService {
	return &CoreService{
		UserService:  NewUserService(),
		AboutService: NewAboutService(),
	}
}
