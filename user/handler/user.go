package handler

import (
	"context"
	userpb "shop/user/api-gen/v1"
	"shop/user/internal/pkg/model"
	"shop/user/repo"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUserList(ctx context.Context, req *userpb.PageInfo,
) (*userpb.UserListResponse, error) {
	// 获取用户列表
	users, err := repo.GetUserList(int(req.GetPn()), int(req.GetPSize()))
	if err != nil {
		return nil, err
	}
	resp := &userpb.UserListResponse{}
	for _, user := range users {
		userInfo := ModelToResponse(user)
		resp.UserInfo = append(resp.GetUserInfo(), userInfo)
	}
	return resp, nil
}

func ModelToResponse(user *model.UserM) *userpb.UserInfo {
	// 在grpc的message中字段有默认值，不能随便赋值nil，容易出错
	// 这里需要清楚，那些字段存在默认值
	UserListRsp := &userpb.UserInfo{
		UserId:   user.UserId,
		Password: user.Password,
		NikeName: user.NikeName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}
	return UserListRsp
}
