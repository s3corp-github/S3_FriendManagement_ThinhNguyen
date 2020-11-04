package model

import (
	"S3_FriendManagement_ThinhNguyen/utils"
	"errors"
)

type FriendConnectionRequest struct {
	Friends []string `json:"friends"`
}

func (_self FriendConnectionRequest) Validate() error {
	if _self.Friends == nil {
		return errors.New("\"friends\" is required")
	}
	if len(_self.Friends) != 2 {
		return errors.New("must has 2 email addresses")
	}

	isValidFirstEmail, firstErr := utils.IsValidEmail(_self.Friends[0])
	if firstErr != nil {
		return errors.New("validate first \"email\" format failed")
	}
	if !isValidFirstEmail {
		return errors.New("first \"email\" is not valid. (ex: \"andy@abc.xyz\")")
	}

	isValidSecondEmail, secondErr := utils.IsValidEmail(_self.Friends[1])
	if secondErr != nil {
		return errors.New("validate second \"email\" format failed")
	}
	if !isValidSecondEmail {
		return errors.New("second \"email\" is not valid. (ex: \"andy@abc.xyz\")")
	}

	return nil
}

//Service model
type FriendsServiceInput struct {
	FirstID  int `json:"first_id"`
	SecondID int `json:"second_id"`
}

//Repo model
type FriendsRepoInput struct {
	FirstID  int `json:"first_id"`
	SecondID int `json:"second_id"`
}
