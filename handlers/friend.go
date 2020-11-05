package handlers

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/service"
	"encoding/json"
	"errors"
	"net/http"
)

type FriendHandler struct {
	IUserService    service.IUserService
	IFriendServices service.IFriendService
}

func (_self FriendHandler) CreateFriend(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	friendRequest := model.FriendConnectionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validation
	if err := friendRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate before creating friend
	IDs, statusCode, err := _self.CreateFriendValidation(friendRequest)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	//Model UserIDs service input
	friendsInputModel := &model.FriendsServiceInput{
		FirstID:  IDs[0],
		SecondID: IDs[1],
	}

	//Call service to create friend connection
	if err := _self.IFriendServices.CreateFriend(friendsInputModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Response
	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
	})
	return
}

func (_self FriendHandler) GetFriendListByEmail(w http.ResponseWriter, r *http.Request) {
	//Decode request body
	friendRequest := model.FriendGetFriendListRequest{}
	if err := json.NewDecoder(r.Body).Decode(&friendRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validation
	if err := friendRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Check existed email and get ID by email
	userID, statusCode, err := _self.GetFriendListValidation(friendRequest.Email)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	//Call service
	friendList, err := _self.IFriendServices.GetFriendListByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Response

	json.NewEncoder(w).Encode(model.FriendsResponse{
		Success: true,
		Friends: friendList,
		Count:   len(friendList),
	})
}

func (_self FriendHandler) GetCommonFriendListByEmails(w http.ResponseWriter, r *http.Request) {
}

func (_self FriendHandler) CreateFriendValidation(friendConnectionRequest model.FriendConnectionRequest) ([]int, int, error) {
	//Check first email valid
	firstUserID, err := _self.IUserService.GetUserIDByEmail(friendConnectionRequest.Friends[0])

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if firstUserID == 0 {
		return nil, http.StatusBadRequest, errors.New("the first email does not exist")
	}

	//Check first email valid
	secondUserID, err := _self.IUserService.GetUserIDByEmail(friendConnectionRequest.Friends[1])

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if secondUserID == 0 {
		return nil, http.StatusBadRequest, errors.New("the second email does not exist")
	}

	// Check friend connection exists
	existed, err := _self.IFriendServices.IsExistedFriend(firstUserID, secondUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if existed {
		return nil, http.StatusAlreadyReported, errors.New("friend connection existed")
	}

	//check blocking between 2 emails
	blocked, err := _self.IFriendServices.IsBlockedEachOther(firstUserID, secondUserID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if blocked {
		return nil, http.StatusPreconditionFailed, errors.New("emails blocked each other")
	}

	return []int{firstUserID, secondUserID}, 0, nil
}

func (_self FriendHandler) GetFriendListValidation(email string) (int, int, error) {
	//Check first email valid
	userID, err := _self.IUserService.GetUserIDByEmail(email)

	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	if userID == 0 {
		return 0, http.StatusBadRequest, errors.New("email does not exist")
	}

	return userID, 0, nil
}
