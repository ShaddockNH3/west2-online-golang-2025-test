package main

import (
	"context"
	user "github.com/ShaddockNH3/west2-online-golang-2025-test/task5/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterUserResponse, err error) {
	// TODO: Your code here...
	return
}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.LoginUserRequest) (resp *user.LoginUserResponse, err error) {
	// TODO: Your code here...
	return
}

// InfoUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) InfoUser(ctx context.Context, req *user.InfoUserRequest) (resp *user.InfoUserResponse, err error) {
	// TODO: Your code here...
	return
}

// AvatarUploadUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) AvatarUploadUser(ctx context.Context, req *user.AvatarUploadUserRequest) (resp *user.AvatarUploadUserResponse, err error) {
	// TODO: Your code here...
	return
}

// QrcodeMFAAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) QrcodeMFAAuth(ctx context.Context, req *user.QrcodeMFAAuthRequest) (resp *user.QrcodeMFAAuthResponse, err error) {
	// TODO: Your code here...
	return
}

// BindMFAAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) BindMFAAuth(ctx context.Context, req *user.BindMFAAuthRequest) (resp *user.BindMFAAuthResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchImage implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchImage(ctx context.Context, req *user.SearchImageRequest) (resp *user.SearchImageResponse, err error) {
	// TODO: Your code here...
	return
}
