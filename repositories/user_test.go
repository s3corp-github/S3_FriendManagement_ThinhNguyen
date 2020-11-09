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
		{
			name:           "User not exist",
			input:          "abcd@xyz.com",
			expectedResult: false,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/user",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			if err := testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath); err != nil {
				//t.Error(err)
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

func TestUserRepo_GetUserIDByEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get UserID failed with error",
			input:          "abc@xyz.com",
			expectedResult: 0,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "The user does not exist",
			input:          "mlk@xyz.com",
			expectedResult: 0,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/user",
		},
		{
			name:           "Get UserID by email success",
			input:          "abc@xyz.com",
			expectedResult: 1,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "./testdata/user/user.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.GetUserIDByEmail(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestUserRepo_GetEmailListByIDs(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult []string
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "No data input",
			input:          []int{},
			expectedResult: []string{},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "",
		},
		{
			name:           "Failed with error",
			input:          []int{1},
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "Get email list from UserID list success",
			input:          []int{1},
			expectedResult: []string{"abc@xyz.com"},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/user",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.GetEmailListByIDs(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}
