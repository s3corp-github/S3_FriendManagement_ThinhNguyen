package service

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type IFriendService interface {
	CreateFriend(*model.FriendsServiceInput) error
	GetFriendListByID(int) ([]string, error)
	IsBlockedEachOther(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
	IUserRepo   repositories.IUserRepo
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

func (_self FriendService) GetFriendListByID(userID int) ([]string, error) {
	blockList := make(map[int]bool)
	friendListID, err := _self.IFriendRepo.GetFriendListByID(userID)
	if err != nil {
		return nil, err
	}

	//Get blocked UserIDs
	blockedListID, err := _self.IFriendRepo.GetBlockedListByID(userID)
	if err != nil {
		return nil, err
	}
	for _, id := range blockedListID {
		blockList[id] = true
	}

	//Get blocking UserIDs
	blockingListID, err := _self.IFriendRepo.GetBlockingListByID(userID)
	if err != nil {
		return nil, err
	}
	for _, id := range blockingListID {
		blockList[id] = true
	}

	//Get UserID list with no blocked
	friendListIDNoBlock := make([]int, 0)
	for _, id := range friendListID {
		if _, isBlock := blockList[id]; !isBlock {
			friendListIDNoBlock = append(friendListIDNoBlock, id)
		}
	}

	friendListEmail, err := _self.IUserRepo.GetEmailListByIDs(friendListIDNoBlock)
	if err != nil {
		return nil, err
	}
	return friendListEmail, err
}

func (_self FriendService) IsBlockedEachOther(firstUserID int, secondUserID int) (bool, error) {
	isBlocked, err := _self.IFriendRepo.IsBlockedEachOther(firstUserID, secondUserID)
	return isBlocked, err
}

func (_self FriendService) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	existed, err := _self.IFriendRepo.IsExistedFriend(firstUserID, secondUserID)
	return existed, err
}
