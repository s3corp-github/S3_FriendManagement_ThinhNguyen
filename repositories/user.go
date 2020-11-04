package repository

import "S3_FriendManagement_ThinhNguyen/model"

type IUserRepo interface {
	IsExistedUser(string) (bool, error)
	Create(*model.UserRepoInput) error
}

type UserRepo struct {

}

func (_self UserRepo) IsExistedUser(string) (bool, error) {
	return false, nil
}

func (_self UserRepo) Create(userRepoInput *model.UserRepoInput) error {
	return nil
}