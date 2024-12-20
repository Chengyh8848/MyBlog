package core

type CoreService struct {
	*UserService
	*AboutService
	*BlogService
}

func NewCoreService() *CoreService {
	return &CoreService{
		UserService:  NewUserService(),
		AboutService: NewAboutService(),
		BlogService:  NewBlogService(),
	}
}
