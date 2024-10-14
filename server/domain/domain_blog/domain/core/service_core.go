package core

type CoreService struct {
	*UserService
}

func NewCoreService() *CoreService {
	return &CoreService{
		UserService: NewUserService(),
	}
}
