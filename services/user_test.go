package services

import (
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/require"
)

func TestUserService_CreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		input         *model.UserServiceInput
		expectedErr   error
		mockRepoInput *model.UserRepoInput
		mockRepoErr   error
	}{
		{
			name: "Create user failed with error",
			input: &model.UserServiceInput{
				Email: "xyz@gmail.com",
			},
			expectedErr: errors.New("create user failed with error"),
			mockRepoInput: &model.UserRepoInput{
				Email: "xyz@gmail.com",
			},
			mockRepoErr: errors.New("create user failed with error"),
		},
		{
			name: "Create user success",
			input: &model.UserServiceInput{
				Email: "xyz@gmail.com",
			},
			expectedErr: nil,
			mockRepoInput: &model.UserRepoInput{
				Email: "xyz@gmail.com",
			},
			mockRepoErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserRepo := new(mockUserRepo)
			mockUserRepo.On("CreateUser", testCase.mockRepoInput).
				Return(testCase.mockRepoErr)

			service := UserService{
				IUserRepo: mockUserRepo,
			}

			//When
			err := service.CreateUser(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
