package repositories

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/mock"
)

type mockUserDB struct {
	mock.Mock
}

func (_self mockUserDB) CreateUser(userRepoInput *model.UserRepoInput) error {
	args := _self.Called(userRepoInput)
	var r error
	if args.Get(0) != nil {
		r = args.Get(0).(error)
	}
	return r
}

func (_self mockUserDB) GetUserIDByEmail(email string) (int, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(int)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserDB) IsExistedUser(email string) (bool, error) {
	args := _self.Called(email)
	r0 := args.Get(0).(bool)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}

func (_self mockUserDB) GetEmailListByIDs(userIDs []int) ([]string, error) {
	args := _self.Called(userIDs)
	r0 := args.Get(0).([]string)
	var r1 error
	if args.Get(1) != nil {
		r1 = args.Get(1).(error)
	}
	return r0, r1
}
