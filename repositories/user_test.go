package repositories

import (
	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/testhelpers"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepo_CreateUser(t *testing.T) {
	testCases := []struct {
		name        string
		input       *model.UserRepoInput
		expectedErr error
		mockDB      *sql.DB
	}{
		{
			name: "Create new user failed with error",
			input: &model.UserRepoInput{
				Email: "abc@gmail.com",
			},
			expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDB:      testhelpers.ConnectDBFailed(),
		},
		{
			name: "Create user success",
			input: &model.UserRepoInput{
				Email: "xyz@abc.com",
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			dbMock := testCase.mockDB

			UserRepo := UserRepo{
				Db: dbMock,
			}

			// When
			err := UserRepo.CreateUser(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserRepo_IsExistedUser(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check existed failed with error",
			input:          "abc@xyz.com",
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "User existed",
			input:          "abc@xyz.com",
			expectedResult: true,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/user",
		},
		//{
		//	name:           "User not exist",
		//	input:          "abcd@xyz.com",
		//	expectedResult: false,
		//	expectedErr:    nil,
		//	mockDb:         testhelpers.ConnectDB(),
		//	preparePath:    "../testhelpers/user",
		//},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			if err := testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath); err != nil {
				t.Error(err)
			}

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.IsExistedUser(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}
