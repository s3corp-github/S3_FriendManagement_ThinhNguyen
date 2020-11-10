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
	friendIDs, err := _self.IFriendRepo.GetFriendListByID(userID)
	if err != nil {
		return nil, err
	}

	//Get blocked UserIDs
	blockedIDs, err := _self.IFriendRepo.GetBlockedListByID(userID)
	if err != nil {
		return nil, err
	}
	for _, id := range blockedIDs {
		blockList[id] = true
	}

	//Get blocking UserIDs
	blockingIDs, err := _self.IFriendRepo.GetBlockingListByID(userID)
	if err != nil {
		return nil, err
	}
	for _, id := range blockingIDs {
		blockList[id] = true
	}

	//Get UserID list with no blocked
	friendIDsNoBlock := make([]int, 0)
	for _, id := range friendIDs {
		if _, isBlock := blockList[id]; !isBlock {
			friendIDsNoBlock = append(friendIDsNoBlock, id)
		}
	}

	friendEmails, err := _self.IUserRepo.GetEmailListByIDs(friendIDsNoBlock)
	if err != nil {
		return nil, err
	}
	return friendEmails, err
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
	firstFriends, err := _self.GetFriendListByID(userIDList[0])
	if err != nil {
		return nil, err
	}
	secondFriends, err := _self.GetFriendListByID(userIDList[1])
	if err != nil {
		return nil, err
	}

	//Get common friends
	commonFriends := make([]string, 0)
	commonMap := make(map[string]bool)
	for _, firstEmail := range firstFriends {
		commonMap[firstEmail] = true
	}

	for _, secondEmail := range secondFriends {
		if _, ok := commonMap[secondEmail]; ok {
			commonFriends = append(commonFriends, secondEmail)
		}
	}

	return commonFriends, nil
}

func (_self FriendService) GetEmailsReceiveUpdate(senderID int, mentionedEmails []string) ([]string, error) {
	result := make([]string, 0)
	resultIDs := make([]int, 0)
	existedResultIDsMap := make(map[int]bool)
	blockedIDsMap := make(map[int]bool)

	//Get blocked list IDs by sender
	blockedUserIDs, err := _self.IFriendRepo.GetBlockedListByID(senderID)
	if err != nil {
		return nil, err
	}
	for _, IDBlocked := range blockedUserIDs {
		blockedIDsMap[IDBlocked] = true
	}

	//Get friend connection by senderID
	friendIDs, err := _self.IFriendRepo.GetFriendListByID(senderID)
	if err != nil {
		return nil, err
	}
	for _, ID := range friendIDs {
		if _, ok := blockedIDsMap[ID]; !ok {
			//Insert to result and existed list
			resultIDs = append(resultIDs, ID)
			existedResultIDsMap[ID] = true
		}
	}

	//Get subscribers by senderID
	subscriberIDs, err := _self.IFriendRepo.GetSubscriberList(senderID)
	if err != nil {
		return nil, err
	}
	for _, ID := range subscriberIDs {
		if _, ok := blockedIDsMap[ID]; !ok {
			//If not in emailMap then append to result and add to map
			if _, ok := existedResultIDsMap[ID]; !ok {
				resultIDs = append(resultIDs, ID)
				existedResultIDsMap[ID] = true
			}
		}
	}

	//Get mentioned emails
	mentionedEmailIDs, err := _self.IUserRepo.GetUserIDsByEmails(mentionedEmails)
	if err != nil {
		return nil, err
	}
	for _, ID := range mentionedEmailIDs {
		if _, ok := blockedIDsMap[ID]; !ok {
			//If not in emailMap then append to result and add to map
			if _, ok := existedResultIDsMap[ID]; !ok {
				resultIDs = append(resultIDs, ID)
				existedResultIDsMap[ID] = true
			}
		}
	}

	//Get emails to return
	emails, err := _self.IUserRepo.GetEmailListByIDs(resultIDs)
	if err != nil {
		return nil, err
	}
	for _, email := range emails {
		result = append(result, email)
	}

	return result, nil
}
