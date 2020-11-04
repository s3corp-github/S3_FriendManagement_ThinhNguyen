package handler

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/service"
	"encoding/json"
	"errors"
	"net/http"
)

type UserHandler struct {
	IUserService service.IUserService
}

func (_self *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	//Decode request body
	userRequest := model.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Validation
	if err := userRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if statusCode, err := _self.IsExistedUser(userRequest.Email); err != nil{
		http.Error(w, err.Error(), statusCode)
		return
	}

	//Convert to service input model
	userServiceInp := &model.UserServiceInput{
		Email: userRequest.Email,
	}

	//Call service
	if err := _self.IUserService.Create(userServiceInp); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Response
	json.NewEncoder(w).Encode(&model.SuccessResponse{
		Success: true,
	})
}

func (_self *UserHandler) IsExistedUser(email string) (int, error) {
	//Call service
	existed, err := _self.IUserService.IsExistedUser(email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if existed{
		return http.StatusAlreadyReported, errors.New("this email address existed")
	}
	return 0, nil
}