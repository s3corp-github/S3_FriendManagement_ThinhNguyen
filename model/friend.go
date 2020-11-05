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
		return errors.New("needs exactly two email addresses")
	}
	if _self.Friends[0] == _self.Friends[1] {
		return errors.New("two email addresses must be different")
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

type FriendGetFriendListRequest struct {
	Email string `json:"email"`
}

func (_self FriendGetFriendListRequest) Validate() error {
	if _self.Email == "" {
		return errors.New("\"Email\" is required")
	}
	isValidFirstEmail, firstErr := utils.IsValidEmail(_self.Email)
	if firstErr != nil {
		return errors.New("validate \"email\" format failed")
	}
	if !isValidFirstEmail {
		return errors.New("\"email\" format is not valid. (ex: \"andy@abc.xyz\")")
	}

	return nil
}

type FriendGetCommonFriendsRequest struct {
	Friends []string `json:"friends"`
}

func (_self FriendGetCommonFriendsRequest) Validate() error {
	if _self.Friends == nil {
		return errors.New("\"friends\" is required")
	}
	if len(_self.Friends) != 2 {
		return errors.New("needs exactly two email addresses")
	}
	if _self.Friends[0] == _self.Friends[1] {
		return errors.New("two email addresses must be different")
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

type FriendsResponse struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
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
