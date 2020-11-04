package handler

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	mock.Mock
}

func (_self mockUserService) IsExistedUser(email string) (bool, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserService) Create(model *model.UserServiceInput) error {
	args := _self.Called(model)
	var r error
	if args.Get(0) != nil{
		r = args.Get(0).(error)
	}
	return r
}