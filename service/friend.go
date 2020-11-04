package service

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type IFriendService interface {
	CreateFriend(*model.FriendsServiceInput) error
	IsBlockedEachOther(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
}

func (_self FriendService) CreateFriend(friendsServiceInput *model.FriendsServiceInput) error {
	//convert to repo input model
	friendsRepoInput := &model.FriendsRepoInput{
		FirstID:  friendsServiceInput.FirstID,
		SecondID: friendsServiceInput.SecondID,
	}

	//Call repo
	err := _self.IFriendRepo.CreateFriend(friendsRepoInput)
	return err
}

func (_self FriendService) IsBlockedEachOther(firstUserID int, secondUserID int) (bool, error) {
	isBlocked, err := _self.IFriendRepo.IsBlockedEachOther(firstUserID, secondUserID)
	return isBlocked, err
}

func (_self FriendService) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	existed, err := _self.IFriendRepo.IsExistedFriend(firstUserID, secondUserID)
	return existed, err
}
