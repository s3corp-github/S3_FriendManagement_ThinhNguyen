package services

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/repositories"
)

type IFriendService interface {
	CreateFriend(*model.FriendsServiceInput) error
	GetCommonFriendListByID([]int) ([]string, error)
	GetFriendListByID(int) ([]string, error)
	IsBlockedByOtherEmail(int, int) (bool, error)
	IsExistedFriend(int, int) (bool, error)
	GetEmailsReceiveUpdate(int, []string) ([]string, error)
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
	//Get all friend connection
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

func (_self FriendService) IsBlockedByOtherEmail(firstUserID int, secondUserID int) (bool, error) {
	isBlocked, err := _self.IFriendRepo.IsBlockedByOtherEmail(firstUserID, secondUserID)
	return isBlocked, err
}

func (_self FriendService) IsExistedFriend(firstUserID int, secondUserID int) (bool, error) {
	existed, err := _self.IFriendRepo.IsExistedFriend(firstUserID, secondUserID)
	return existed, err
}

func (_self FriendService) GetCommonFriendListByID(userIDList []int) ([]string, error) {
	firstFriendList, err := _self.GetFriendListByID(userIDList[0])
	if err != nil {
		return nil, err
	}
	secondFriendList, err := _self.GetFriendListByID(userIDList[1])
	if err != nil {
		return nil, err
	}

	//Get common friends
	commonFriendList := make([]string, 0)
	commonMap := make(map[string]bool)
	for _, firstEmail := range firstFriendList {
		commonMap[firstEmail] = true
	}

	for _, secondEmail := range secondFriendList {
		if _, ok := commonMap[secondEmail]; ok {
			commonFriendList = append(commonFriendList, secondEmail)
		}
	}

	return commonFriendList, nil
}

func (_self FriendService) GetEmailsReceiveUpdate(senderID int, mentionedEmails []string) ([]string, error) {
	//Friend email list which can receive update from sender, for check existed
	friendConnectionListMap := make(map[string]bool)

	//Blocked email list which can NOT receive update from sender
	blockListMap := make(map[string]bool)

	//Friend email list which can receive update from sender, for return
	result := make([]string, 0)

	//Get email list which blocked by sender
	blockedUserIDList, err := _self.IFriendRepo.GetBlockedListByID(senderID)
	if err != nil {
		return nil, err
	}
	blockedEmailList, err := _self.IUserRepo.GetEmailListByIDs(blockedUserIDList)
	if err != nil {
		return nil, err
	}

	for _, emailBlocked := range blockedEmailList {
		blockListMap[emailBlocked] = true
	}

	//Get all email friend connection
	friendIDList, err := _self.IFriendRepo.GetFriendListByID(senderID)
	if err != nil {
		return nil, err
	}
	friendEmailList, err := _self.IUserRepo.GetEmailListByIDs(friendIDList)
	if err != nil {
		return nil, err
	}

	//Set data for result
	for _, email := range friendEmailList {
		if _, ok := blockListMap[email]; !ok {
			//Insert to result and Friend list
			result = append(result, email)
			friendConnectionListMap[email] = true
		}
	}

	//Get subscriber list
	subscriberIDList, err := _self.IFriendRepo.GetSubscriberList(senderID)
	if err != nil {
		return nil, err
	}
	subscriberEmails, err := _self.IUserRepo.GetEmailListByIDs(subscriberIDList)
	if err != nil {
		return nil, err
	}
	for _, email := range subscriberEmails {
		// If not blocked
		if _, ok := blockListMap[email]; !ok {
			// If not in emailMap then append to result and add to map
			if _, ok := friendConnectionListMap[email]; !ok {
				result = append(result, email)
				friendConnectionListMap[email] = true
			}
		}
	}

	//Email in @mentioned
	for _, email := range mentionedEmails {
		// If not blocked
		if _, ok := blockListMap[email]; !ok {
			// If not exist in result
			if _, ok := friendConnectionListMap[email]; !ok {
				result = append(result, email)
				friendConnectionListMap[email] = true
			}
		}
	}

	return result, nil
}
