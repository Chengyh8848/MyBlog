package core

import (
	"context"
	"domain_blog/common"
	"domain_blog/infrastructure/database/entity"
	"domain_blog/infrastructure/database/persistence"
	pb "domain_blog/infrastructure/grpc/pb"
	"domain_blog/internal/cacheUtil"
	"domain_blog/internal/crypto"
	"domain_blog/internal/helper"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type UserDomainI interface {
	LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserReply, error)
	LogoutUser(ctx context.Context, request *pb.LogoutUserRequest) (*pb.LogoutUserReply, error)
	ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error)
	AuthToken(ctx context.Context, request *pb.AuthTokenRequest) (*pb.AuthTokenReply, error)
	FindByUsername(ctx context.Context, request *pb.FindByUsernameRequest) (*pb.FindByUsernameReply, error)
	FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdReply, error)
	UpdateUserByUsername(ctx context.Context, request *pb.UpdateUserByUsernameRequest) (*pb.UpdateUserByUsernameReply, error)
}

type UserService struct {
	UserPersistence persistence.UserPersistenceI
}

func NewUserService() *UserService {
	return &UserService{
		UserPersistence: persistence.NewUserPersistence(),
	}
}

type UserCtx struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *UserService) FindByUsername(ctx context.Context, request *pb.FindByUsernameRequest) (*pb.FindByUsernameReply, error) {
	return nil, nil
}
func (s *UserService) FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdReply, error) {
	return nil, nil
}
func (s *UserService) UpdateUserByUsername(ctx context.Context, request *pb.UpdateUserByUsernameRequest) (*pb.UpdateUserByUsernameReply, error) {
	return nil, nil
}

func (s *UserService) AuthToken(ctx context.Context, request *pb.AuthTokenRequest) (*pb.AuthTokenReply, error) {
	key := fmt.Sprintf("token:%s", request.Token)
	userByte, exists := cacheUtil.Get(key)
	if exists {
		user := UserCtx{}
		str := userByte.(string)
		err := json.Unmarshal([]byte(str), &user)
		if err != nil {
			common.Log.ErrorMsg("解析用户token失败:", err.Error())
			return nil, errors.New("请先登录")
		}
		return &pb.AuthTokenReply{}, nil
	} else {
		common.Log.ErrorMsg("获取用户token失败")
		return nil, errors.New("请先登录")
	}
}

// 检查用户名和密码是否匹配
func (s *UserService) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserReply, error) {
	u, err := s.UserPersistence.GetUserByName(ctx, request.Username)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("不存在该用户")
	}
	if crypto.EncryptData(request.Password) == u.Password {
		userCtx := UserCtx{
			Id:   u.ID,
			Name: u.Username,
		}
		token := helper.RandStr(32)
		data, _ := json.Marshal(userCtx)
		cacheUtil.Set("token:"+string(token), string(data), time.Hour*24*7)
		return &pb.LoginUserReply{Token: token}, nil
	} else {
		return nil, errors.New("用户名或者密码错误")
	}
}

func (s *UserService) LogoutUser(ctx context.Context, request *pb.LogoutUserRequest) (*pb.LogoutUserReply, error) {
	token := request.Token
	key := fmt.Sprintf("token:%s", token)
	cacheUtil.Delete(key)
	return &pb.LogoutUserReply{}, nil
}

// 修改密码
func (s *UserService) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	u, err := s.UserPersistence.GetUserByName(ctx, request.Username)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("不存在该用户")
	}
	if crypto.EncryptData(request.OldPassword) != u.Password {
		return nil, errors.New("旧密码错误")
	}
	if request.OldPassword == request.NewPassword {
		return nil, errors.New("新旧密码不能一样")
	}
	u.Password = crypto.EncryptData(request.NewPassword)
	err = s.UserPersistence.Update(nil, u)
	if err != nil {
		return nil, err
	}
	// 删除token
	token := request.Token
	key := fmt.Sprintf("token:%s", token)
	cacheUtil.Delete(key)
	return &pb.ChangePasswordReply{}, nil
}

// 获取用户信息
func (s *UserService) GetUserInfo(name string) (u *entity.User, err error) {
	u, err = s.UserPersistence.GetUserByName(nil, name)
	return
}

// 初始化用户
func (s *UserService) ForceCreateAdmin() error {
	if err := s.UserPersistence.Delete(nil, "admin"); err != nil {
		return err
	}
	salt := helper.RandStr(8)
	u := &entity.User{
		Username: "admin",
		Password: crypto.EncryptData("test-123"),
		Salt:     salt,
	}
	if err := s.UserPersistence.Create(nil, u); err != nil {
		return err
	}
	return nil
}

// 是否已创建admin用户
func (s *UserService) IsExistAdmin() bool {
	u, _ := s.UserPersistence.GetUserByName(nil, "admin")
	return u != nil
}
